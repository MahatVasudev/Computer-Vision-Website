import torch
import torch.nn as nn
import torch.nn.functional as F
import math


class ModelOneDiff(nn.Module):
    def __init__(self, base_channels=64):
        super().__init__()
        
        # Encoder
        self.enc1 = ResBlock(4, base_channels)
        self.enc2 = ResBlock(base_channels, base_channels*2, downsample=True)
        self.enc3 = ResBlock(base_channels*2, base_channels*4, downsample=True, use_attention=True)
        self.enc4 = ResBlock(base_channels*4, base_channels*8, downsample=True, use_attention=True)
        
        # Bottleneck
        self.bottleneck = ResBlock(base_channels*8, base_channels*8, use_attention=True)
        self.time_emb = TimeEmbedding(base_channels*8)
        
        # Decoder
        self.dec4 = ResBlock(base_channels*8*2, base_channels*4, upsample=True, use_attention=True)
        self.dec3 = ResBlock(base_channels*4*2, base_channels*2, upsample=True, use_attention=True)
        self.dec2 = ResBlock(base_channels*2*2, base_channels, upsample=True)
        self.dec1 = ResBlock(base_channels*2, 32)

        self.final = nn.Conv2d(32, 3, kernel_size=1)
        
    def forward(self, x, t):
        t_emb = self.time_emb(t)
        
        # Encoder
        e1 = self.enc1(x)
        e2 = self.enc2(e1)
        e3 = self.enc3(e2)
        e4 = self.enc4(e3)
        
        # Bottleneck
        b = self.bottleneck(e4)
        b = b + t_emb
        
        # Decoder with skip connections
        d4 = self.dec4(torch.cat([b, e4], dim=1))
        d3 = self.dec3(torch.cat([d4, e3], dim=1))
        d2 = self.dec2(torch.cat([d3, e2], dim=1))
        d1 = self.dec1(torch.cat([d2, e1], dim=1))
        d0 = self.final(d1)
        return d0


class ResBlock(nn.Module):
    def __init__(self, in_channels, out_channels, downsample=False, upsample=False, use_attention=False, dropout_prob=0.5):
        super().__init__()
        self.downsample = downsample
        self.upsample = upsample
        
        # Main layers
        self.conv1 = nn.Conv2d(in_channels, out_channels, kernel_size=3, padding=1)
        self.conv2 = nn.Conv2d(out_channels, out_channels, kernel_size=3, padding=1)
        self.norm1 = nn.GroupNorm(32, out_channels)
        self.norm2 = nn.GroupNorm(32, out_channels)
        self.dropout = nn.Dropout(dropout_prob)  # Add dropout
        
        # Down/Upsample layers
        if downsample:
            self.downsample_conv = nn.Conv2d(out_channels, out_channels, kernel_size=3, stride=2, padding=1)
        if upsample:
            self.upsample_conv = nn.ConvTranspose2d(out_channels, out_channels, kernel_size=3, stride=2, padding=1, output_padding=1)
        
        # Shortcut
        self.shortcut = nn.Identity()
        if in_channels != out_channels or downsample or upsample:
            layers = []
            if downsample:
                layers.append(nn.AvgPool2d(2))
            elif upsample:
                layers.append(nn.Upsample(scale_factor=2, mode='nearest'))
            layers.append(nn.Conv2d(in_channels, out_channels, kernel_size=1))
            self.shortcut = nn.Sequential(*layers)
        
        # Attention
        self.use_attention = use_attention
        if use_attention:
            self.attention = SelfAttention(out_channels)

    def forward(self, x):
        residual = self.shortcut(x)
        
        x = self.conv1(x)
        x = self.norm1(x)
        x = F.relu(x)
        x = self.dropout(x)  # Apply dropout after the first activation
        
        if self.downsample:
            x = self.downsample_conv(x)
        elif self.upsample:
            x = self.upsample_conv(x)
        
        x = self.conv2(x)
        x = self.norm2(x)
        
        if self.use_attention:
            x = self.attention(x)
        
        x = F.relu(x + residual)
        return x


class TimeEmbedding(nn.Module):
    def __init__(self, dim):
        super().__init__()
        self.dim = dim
        half_dim = dim // 2
        emb = math.log(10000) / (half_dim - 1)
        emb = torch.exp(torch.arange(half_dim, dtype=torch.float32) * -emb)
        self.register_buffer('emb', emb)

    def forward(self, t):
        t = t.float()[:, None] * self.emb[None, :]
        t = torch.cat([torch.sin(t), torch.cos(t)], dim=1)
        return t.view(-1, self.dim, 1, 1)


class SelfAttention(nn.Module):
    def __init__(self, channels):
        super().__init__()
        self.query = nn.Conv2d(channels, channels, kernel_size=1)
        self.key = nn.Conv2d(channels, channels, kernel_size=1)
        self.value = nn.Conv2d(channels, channels, kernel_size=1)
        self.softmax = nn.Softmax(dim=-1)

    def forward(self, x):
        b, c, h, w = x.shape
        q = self.query(x).view(b, c, -1).permute(0, 2, 1)
        k = self.key(x).view(b, c, -1)
        v = self.value(x).view(b, c, -1).permute(0, 2, 1)

        attn = self.softmax(torch.bmm(q, k))
        out = torch.bmm(attn, v).permute(0, 2, 1).view(b, c, h, w)

        return out + x



def AutoColorizationInstance(weights_location: str, key: str = "", device: str = "cpu"):
    
    model_params = torch.load(weights_location, weights_only=True)
    
    model = ModelOneDiff().to(device)

    model.eval()
    if key == "" :
        weights = model_params 
    else:
        weights = model_params[key]

    model.load_state_dict(weights)

    return model

import torch 
import torch.nn as nn




class UNet_Noise(nn.Module):
    
    def __init__(self, time_emb_dim=128):
        super().__init__()
    
        self.downs = (3, 64, 128, 256, 512)

        self.down_sample = nn.ModuleList([
            UNetBlock(self.downs[i], self.downs[i+1], time_emb_dim) 
                for i in range(len(self.downs)-1)
            ])

        self.pool = nn.MaxPool2d(2)

        self.bottle_neck = UNetBlock(self.downs[-1], self.downs[-1], time_emb_dim)

        self.up_sample = nn.ModuleList([
            UNetBlock(self.downs[j]*2, self.downs[j-1], time_emb_dim) 
                for j in range(len(self.downs)-1, 0, -1)
            ])

        self.upConv = nn.ModuleList([
            nn.ConvTranspose2d(self.downs[k], self.downs[k-1], kernel_size=2, stride=2)
                for k in range(len(self.downs)-1, 0,-1)
            ])

    
    def forward(self, x, t):
        t_emb = self.time_embedding(t)

        skip_conn = []
        
        for down in self.down_sample:
            x = down(x, t_emb)
            skip_conn.append(x)

        x = self.bottle_neck(x)

        for up, upconv, skip in zip(self.up_sample, self.upConv, reversed(skip_conn)):
            catted = torch.cat([upconv(x), skip])
            x = up(catted, t_emb)

        return x

    def time_embedding(self, time_emb_dim=128):
        return nn.Sequential(
                nn.Linear(1, time_emb_dim),
                nn.SiLU(),
                nn.Linear(time_emb_dim, time_emb_dim)
                )

class UNetBlock(nn.Module):
    def __init__(self, in_channels, out_channels, time_emb_dim):
        super().__init__()

        self.conv1 = nn.Conv2d(in_channels,out_channels,kernel_size=3, padding=1)

        self.norm1 = nn.GroupNorm(8, out_channels)

        self.conv2 = nn.Conv2d(out_channels, out_channels, kernel_size=3, padding=1)

        self.norm2 = nn.GroupNorm(8, out_channels)

        self.time_emb_layer = nn.Linear(time_emb_dim, out_channels)

        self.act = nn.SiLU()


    def forward(self,x,t):
        time_emb = self.time_emb_layer(t).unsqueeze(-1).unsqueeze(-1)

        x = self.conv1(x) + time_emb
        x = self.norm1(x)
        x = self.act(x)

        x = self.conv2(x)
        x = self.norm2(x)
        x = self.act(x)

        return x


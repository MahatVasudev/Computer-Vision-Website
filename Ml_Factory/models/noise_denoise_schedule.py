import torch
import torch
import torch.nn as nn
import torch.optim as optim
import torch.nn.functional as F
from torchvision.utils import save_image



class Diffusion:
    def __init__(self, timesteps: int = 1000, beta: tuple= (1e-4, 0.02), device="cpu"):
        
        assert len(beta) == 2
        
        self.device = device
        self.timesteps = timesteps

        self.betas = torch.linspace(beta[0], beta[1], timesteps, device=device)

        self.alphas = 1.0 - self.betas

        self.alphas_bars = torch.cumprod(self.alphas, dim=0).to(device)

    def add_noise(self, x0, t):

        noise = torch.randn_like(x0).to(self.device)

        sqrt_alpha_bar_t = torch.sqrt(self.alphas_bars[t]).view(-1,1,1,1).to(self.device)

        sqrt_one_minus_alpha_bar_t = torch.sqrt(1-self.alphas_bars[t]).view(-1,1,1,1).to(self.device)

        xt = sqrt_alpha_bar_t * x0 + sqrt_one_minus_alpha_bar_t * noise

        return xt, noise

    def sample_timesteps(self, batch_size):

        return torch.randint(0, self.timesteps, (batch_size,), dtype=torch.long).to(self.device)



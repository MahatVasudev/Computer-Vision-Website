
import torch
import torch.nn.functional as F

def linear_beta_scheduling(timesteps, beta_start=1e-4, beta_end=0.02):
    return torch.linspace(beta_start, beta_end, timesteps)


def add_noise(x0, t, alpha_hats):
    noise = torch.randn_like(x0)

    sqrt_alpha_hat = torch.sqrt(alpha_hats[t]).view(-1,1,1,1)

    one_minus_sqrt = torch.sqrt(1-alpha_hats[t]).view(-1,1,1,1)

    xt = sqrt_alpha_hat * x0 + one_minus_sqrt * noise

    return xt, noise


def diffusion_loss(model, x0, t, alpha_hats):
    xt, noise = add_noise(x0, t, alpha_hats=alpha_hats)
    predicted_noise = model(xt, t.float())

    return F.mse_loss(predicted_noise, noise)


import torch


def wasserstein_loss(real_output, fake_output):
    return torch.mean(fake_output) - torch.mean(real_output)

def compute_gradient_penalty(D, real_samples, fake_samples, device):
    alpha = torch.rand(real_samples.size(0),1,1,1).to(device)
    interpolates = (alpha *real_samples + (1-alpha)*fake_samples).requires_grad_(True)

    d_interpolates = D(interpolates)
    grad_outputs = torch.ones_like(d_interpolates, requires_grad=False).to(device)

    gradients = torch.autograd.grad(
            outputs=d_interpolates,
            inputs=interpolates,
            grad_outputs=grad_outputs,
            create_graph=True,
            retain_graph=True,
            only_inputs=True
            )[0]

    gradients = gradients.view(gradients.size(0), -1)

    gradients_penalty = ((gradients.norm(2, dim=1)-1)**2).mean()

    return gradients_penalty

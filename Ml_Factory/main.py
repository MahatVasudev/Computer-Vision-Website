import torch
from torch.amp.grad_scaler import GradScaler
from torch.amp.autocast_mode import autocast
import traceback
import torch.nn as nn
import torch.optim as optim
import matplotlib.pyplot as plt
from torch.optim.lr_scheduler import ReduceLROnPlateau
from torch.utils.data import DataLoader, random_split
from torchvision.transforms import transforms
from Datasets.custom_image_dataset import CustomImageDataset
from Datasets.datasets import AnimeAllImages, LandScapeColoredDataset
from models.diffusion_model import ModelOneDiff
from models.noise_denoise_schedule import Diffusion

def denoise(model, gray, noisy_rgb, t):
    """
    Denoiser function: Predicts noise from noisy RGB using the model.
    Args:
        model: The diffusion model.
        gray: Grayscale input image (batch_size, 1, H, W).
        noisy_rgb: Noisy RGB image (batch_size, 3, H, W).
        t: Timestep tensor (batch_size,).
    Returns:
        Predicted noise (batch_size, 3, H, W).
    """
    # Concatenate grayscale and noisy RGB as input
    model_input = torch.cat([gray, noisy_rgb], dim=1).to(device)
    return model(model_input, t)

def train_diffusion(model, train_loader, val_loader, device, epochs=100, timesteps=1000, start_epoch=-1):
    """
    Training process for the diffusion model.
    Args:
        model: The diffusion model.
        train_loader: DataLoader for training data.
        val_loader: DataLoader for validation data.
        device: Device to run training on (e.g., "cuda" or "cpu").
        epochs: Number of epochs to train.
        timesteps: Number of diffusion timesteps.
    """
    # Initialize diffusion scheduler
    # Training loop
    optimizer.zero_grad()
    for epoch in range(start_epoch+1,epochs):
        model.train()
        total_loss = 0
        for step, (gray, rgb) in enumerate(train_loader):
            gray = gray.to(device)
            rgb = rgb.to(device)
            
            with autocast(device_type=device):  # Mixed precision
                t = diffusion.sample_timesteps(gray.size(0))
                noisy_rgb, noise = diffusion.add_noise(rgb, t)
                pred_noise = denoise(model, gray, noisy_rgb, t)
                loss = loss_fn(pred_noise, noise) / accumulation_steps
            
            scaler.scale(loss).backward()
            # Predict noise using the denoiser function
            if (step + 1) % accumulation_steps == 0:
                scaler.step(optimizer)
                scaler.update()
            
                optimizer.zero_grad()



            scheduler.step(loss)
            # Logging every 50 steps
            if step % 50 == 0:
                print(f"Epoch [{epoch+1}/{epochs}] | Step [{step}/{len(train_loader)}] | Loss: {loss.item()}")
            
            # Save images every 250 steps
            if step % 250 == 0:
                save_images(model, diffusion, (gray, rgb), device, epoch, step)

            total_loss += loss.item() if not torch.isnan(loss) else 0


        print(f"Train Loss = {total_loss/len(train_loader)}")
        validation_set(val_loader)


        torch.save({"epoch" : epoch, 
                    "model_param": model.state_dict(),
                    "optimizer_params": optimizer.state_dict()},
                   f"./products/diffusion_model_saved_2_epoch_{epoch}.pth")


def validation_set(validation_dataset):
    model.eval()
    total_val_loss = 0
    with torch.no_grad():
        for i, (gray, rgb) in enumerate(val_loader):
            gray, rgb = gray.to(device), rgb.to(device)
            t = diffusion.sample_timesteps(gray.size(0))
            noisy_rgb, noise = diffusion.add_noise(rgb, t)
            pred_noise = denoise(model, gray, noisy_rgb, t)
            loss = val_loss(pred_noise, noise)
            if i % 250 == 0:
                print(f"Step {i}/{len(validation_dataset)} | Loss {loss.item()}")
            
            total_val_loss += loss.item() if not torch.isnan(loss) else 0

    avg_val_loss = total_val_loss / len(validation_dataset)
    print("avg validation loss: ", avg_val_loss)




def save_images(model, diffusion, images, device, epoch, step):
    model.eval()
    with torch.no_grad():
        val_gray, val_rgb = images
        val_gray, val_rgb = val_gray.to(device), val_rgb.to(device)

        # Start with pure noise
        noisy_rgb = torch.randn_like(val_gray).expand(-1, 3, -1, -1)

        for t in reversed(range(900)):  # Loop from 999 → 0
            t_tensor = torch.full((val_gray.size(0),), t, device=device)

            # Predict noise
            predicted_noise = denoise(model, val_gray, noisy_rgb, t_tensor) 

            # Compute denoised image using diffusion process
            sqrt_alpha_bar = torch.sqrt(diffusion.alphas_bars[t])
            sqrt_one_minus = torch.sqrt(1 - diffusion.alphas_bars[t])
            noisy_rgb = (noisy_rgb - sqrt_one_minus * predicted_noise) / sqrt_alpha_bar

            # Clamp values to valid range
            noisy_rgb = torch.clamp(noisy_rgb, -1, 1)

        denoised = noisy_rgb  # Final denoised image

        # Convert tensors to numpy
        denoised_np = denoised.cpu().permute(0, 2, 3, 1).numpy()
        noisy_np = predicted_noise.cpu().permute(0, 2, 3, 1).numpy()
        gt_np = val_rgb.cpu().permute(0, 2, 3, 1).numpy()

        # Save images
        fig, axes = plt.subplots(1, 3, figsize=(15, 5))
        axes[0].imshow((denoised_np[0] + 1) / 2)
        axes[0].set_title("Denoised")
        axes[0].axis("off")

        axes[1].imshow((noisy_np[0] + 1) / 2)
        axes[1].set_title("Predicted Noise")
        axes[1].axis("off")

        axes[2].imshow((gt_np[0] + 1) / 2)
        axes[2].set_title("Ground Truth")
        axes[2].axis("off")

        plt.suptitle(f"Epoch {epoch+1}, Step {step}")
        plt.savefig(f"./public/images3_epoch_{epoch+1}_step_{step}.png")
        plt.close()


if __name__ == "__main__":
    try:
        # Empty Cache First...
        torch.clear_autocast_cache()
        torch.cuda.empty_cache()


        device = "cuda" if torch.cuda.is_available() else "cpu"
        transform_list = [
                transforms.Resize((64,64)),
                transforms.Normalize(mean=[0.5], std=[0.5])
            ]

        transform_extra = [
            transforms.RandomHorizontalFlip(p=0.5),
            transforms.RandomRotation(30),
            transforms.RandomVerticalFlip(p=0.5),
            transforms.RandomAffine(degrees=15, translate=(0.1, 0.1)),
        ]

        transform_normal = transforms.Compose(transform_list)
        transform_with_rotations = transforms.Compose(transform_list + transform_extra)
        AnimeDataset = AnimeAllImages(transform_normal)
        Landscape = LandScapeColoredDataset(transform_with_rotations)
        Cusdataset = CustomImageDataset("D:/Image Dataset/data/train_color", transform=transform_with_rotations)

        dataset = torch.utils.data.ConcatDataset([Landscape, Cusdataset, AnimeDataset])
        total_size = len(dataset)
        val_size = int(0.2 * total_size)
        train_size = total_size - val_size
        train_dataset, val_dataset = random_split(dataset, [train_size, val_size])
        train_loader = DataLoader(train_dataset, batch_size=4, num_workers=3, shuffle=True)
        val_loader = DataLoader(val_dataset, batch_size=4, num_workers=3, shuffle=False)

        model = ModelOneDiff(base_channels=64*2).to(device)
        optimizer = torch.optim.Adam(model.parameters(), lr=1e-3, weight_decay=1e-5)
        # data = torch.load("./products/diffusion_model_saved_1_epoch_21.pth",weights_only=True)
        
        # model.load_state_dict(data["model_param"])
        # optimizer.load_state_dict(data["optimizer_params"])
        timesteps = 1000
        diffusion = Diffusion(timesteps=timesteps, beta=(1e-4, 0.02), device=device)
    
        # Optimizer and loss function
        scheduler = ReduceLROnPlateau(optimizer, 
                                  mode="min", 
                                  factor=0.01, 
                                  patience=3
                                  )
        loss_fn = nn.MSELoss()
        val_loss = nn.MSELoss()
        scaler = GradScaler(device=device)
        accumulation_steps = 5
        train_diffusion(model, train_loader=train_loader, val_loader = val_loader,device=device, epochs=100)

    except KeyboardInterrupt as KE:
        print("Stopping Training Process\nSaving Model")
        torch.save({"model_params": model.state_dict()}, "InterruptedModel_Diffusion_Trial.pth")
        print("Saved Interrupted Model...")
    except Exception as E:
        print("Error Occured", E)

        traceback.print_exc()
    finally:
        print("Exiting Training Process...")

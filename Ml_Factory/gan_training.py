
import torch
from torch.utils.data import DataLoader
from torchvision.transforms import transforms
from Datasets.custom_image_dataset import CustomImageDataset
from Datasets.datasets import AnimeAllImages, LandScapeColoredDataset
from models.torch_model import ModelFighter, ModelFighter2, ModelOneGen, ModelOneGen2, ModelOneGenX
import torch.nn as nn
import matplotlib.pyplot as plt
import random

from utils.convert_color import convertScalesto255
from utils.wasserstein import compute_gradient_penalty

global device 

device = 'cuda' if torch.cuda.is_available() else 'cpu'
 
modelonegen = ModelOneGenX().to(device)
modelfighter = ModelFighter2().to(device)

optimizer_gen = torch.optim.Adam(modelonegen.parameters(), lr=0.005, betas=(0.5, 0.999))
optimizer_fighter = torch.optim.Adam(modelfighter.parameters(), lr=0.00007, betas=(0.5, 0.999))

scheduler_gen = torch.optim.lr_scheduler.ReduceLROnPlateau(optimizer_gen, mode='min', factor=0.5, patience=5)
scheduler_fighter = torch.optim.lr_scheduler.ReduceLROnPlateau(optimizer_fighter, mode='min', factor=0.5, patience=5)

criterion_g = nn.L1Loss().to(device)
criterion_gmse = nn.MSELoss().to(device)

def Plot(real_image: torch.Tensor, fake_image: torch.Tensor):

    n_rows = 2

    fig, ax = plt.subplots(n_rows,2)
    
    ax[0,0].set_title("Real Image")
    ax[0,1].set_title("AI Image")
    print("Printed Sample...")
    for i in range(n_rows):
        real_images = real_image[i].squeeze(0).permute(1,2,0).cpu().detach().numpy()
        
        real_images = convertScalesto255(real_images)
        fake_images = fake_image[i].squeeze(0).permute(1,2,0).cpu().detach().numpy()
        fake_images = convertScalesto255(fake_images)
        
        ax[i,0].imshow(real_images)
        
        ax[i,1].imshow(fake_images)
    
    fig.tight_layout()
    num = random.randint(1,1_000_000_0)
    plt.savefig(f"./public/Image_{num}.png")
    plt.close()   


def TrainProcess(root_dir: str):
    preprocess = transforms.Compose([
        transforms.Resize((256,256)),
        transforms.Normalize(mean=[0.5], std=[0.5])
        ])
    landscape_preprocess = transforms.Compose([
        transforms.RandomHorizontalFlip(p=0.5),  # 50% chance to flip horizontally
        transforms.RandomRotation(30),           # Rotate randomly within -30 to 30 degrees
        transforms.RandomResizedCrop(256),       # Crop and resize the image
        transforms.ColorJitter(brightness=0.2, contrast=0.2, saturation=0.2, hue=0.2),  # Change color properties
        transforms.RandomVerticalFlip(p=0.5),    # 50% chance to flip vertically
        transforms.RandomAffine(degrees=15, translate=(0.1, 0.1)),  # Apply affine transformation
        transforms.Normalize(mean=[0.5], std=[0.5])
        ])
    
    image_dataset = CustomImageDataset(root_dir, transform=preprocess)
    anime_dataset = AnimeAllImages(transform=preprocess)
    landscape_dataset = LandScapeColoredDataset(transform=landscape_preprocess)
    final_image = torch.utils.data.ConcatDataset([image_dataset,landscape_dataset, anime_dataset])
    dataloader = DataLoader(final_image, batch_size=12, num_workers=2, shuffle=True)

    auto_color_training(dataloader,epochs=10,critic_iterations=3)

def auto_color_training(dataloader, epochs=1, critic_iterations=5):
    for epoch in range(epochs):
        for i, images in enumerate(dataloader):
            
            # Get images
            gray_image, real_image = images
            gray_image, real_image = gray_image.to(device), real_image.to(device)
            
            # print(torch.min(gray_image), torch.max(gray_image))
            # Wasserstein hyperparameters
            lambda_gp = 20  # Gradient Penalty weight
            lambda_color = 50 # Color loss weight
            
            lambdaL1 = 10
            lambda_mse = 1

            d_loss_total = 0

            for _ in range(critic_iterations):
                # --- Discriminator Step ---
                fake_outputs = modelonegen(gray_image).detach() # Detach to avoid tracking gradients
                # Get discriminator results for real images
                real_outputs = modelfighter(real_image)
                real_outputs = real_outputs.view(real_outputs.shape[0], -1)
                real_outputs = real_outputs.mean(dim=1, keepdim=True)

                # Get discriminator results for fake images
                fake_results = modelfighter(fake_outputs)
                fake_results = fake_results.view(fake_results.shape[0], -1)
                fake_results = fake_results.mean(dim=1, keepdim=True)

                # Compute gradient penalty
                gradient_penalty = compute_gradient_penalty(modelfighter, real_image, fake_outputs, device)
            
                # Discriminator loss
                d_loss = torch.mean(fake_results) - torch.mean(real_outputs) + lambda_gp * gradient_penalty
                d_loss_total += d_loss.item()
                optimizer_fighter.zero_grad()
            
                d_loss.backward()
                optimizer_fighter.step()

            d_loss_avg = d_loss_total/critic_iterations
            scheduler_fighter.step(d_loss_avg)

            fake_outputs = modelonegen(gray_image)  # Recompute fake images with gradients
            # Get discriminator results for the new fake images
            fake_results = modelfighter(fake_outputs)
            fake_results = fake_results.view(fake_results.shape[0], -1)
            fake_results = fake_results.mean(dim=1, keepdim=True)

            # Adversarial loss for generator
            loss_adv = -torch.mean(fake_results)

            # Color loss (reconstruction loss)
            recon_l1 = criterion_g(fake_outputs, real_image)
            recon_mse = criterion_gmse(fake_outputs, real_image)

            recon = recon_l1 * lambdaL1 + recon_mse * lambda_mse
            # Total generator loss
            g_loss = loss_adv  + recon

            optimizer_gen.zero_grad()
            g_loss.backward()
            optimizer_gen.step()
            scheduler_gen.step(g_loss)

            if i % 50 == 0:
                print(f"Epoch [{epoch}/{epochs}] ({i}/{len(dataloader)}) | D Loss: {d_loss_avg} | G Loss: {g_loss.item()} | Gradient Penalty: {gradient_penalty}")

            if i % 250 == 0:
                Plot(real_image, fake_outputs)

        # Save model checkpoint
        torch.save({
            "epoch": epoch, 
            "model_params": modelonegen.state_dict(),
            "optimizer_params": optimizer_gen.state_dict()
        }, f"./products/CheckPoint_Epoch_{epoch}_ModelOneGenXo.pth")

if __name__ == "__main__":
    try:
        TrainProcess("D:/Image Dataset/data/train_color")
    except Exception as e:

        raise Exception(e)
    finally:
        torch.save(modelonegen.state_dict(),"./products/ModelOneGenXoPanic.pth")

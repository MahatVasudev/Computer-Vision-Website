from PIL import Image
from io import BytesIO
from fastapi import FastAPI, File, Response, UploadFile
import uvicorn
import numpy as np
from torchvision import transforms
from models.auto_color_model import AutoColorizationInstance
from models.utils import Diffusion
from utils.weights import WEIGHTS
import cv2
import torch
# import torch.nn as nn
import os

global ROOT_DIR
ROOT_DIR = os.getcwd()

app = FastAPI()
device = "cuda" if torch.cuda.is_available() else "cpu"

print("Current Device...")

def weights(location: str) -> str:
    return f"{ROOT_DIR}/{WEIGHTS}/{location}"

AutoColor_ModelInstance = AutoColorizationInstance(weights("diffusion_model_saved_1_epoch_15.pth"), "model_param", device)

timesteps = 1000

diffusion = Diffusion(timesteps=1000, beta=(1e-4, 0.02), device=device)

transform_data = transforms.Compose([
    transforms.ToTensor(),
    transforms.Resize((128,128)),
    transforms.Normalize(mean=[0.5], std=[0.5])
    ])



@app.get("/checkHealth")
def read_root():
    return {"status": True, "message": "still alive..."}




@app.post("/autoColor/color")
async def autocolor_image(file: UploadFile = File(...)):
    
   # Open the uploaded image
    image = Image.open(file.file).convert("L")

    width, height = image.size

    processed_image = transform_data(image)  # Assuming this is defined elsewhere

    AutoColor_ModelInstance.eval()

    with torch.no_grad():
        processed_image = processed_image.unsqueeze(0).to(device)
        noisy_rgb = torch.randn_like(processed_image).expand(-1, 3, -1, -1)
        
        for t in reversed(range(1000)):
            t_tensor = torch.full((processed_image.shape[0],), t, device=device)

            data = torch.cat([processed_image, noisy_rgb], dim=1).to(device)
            predicted_noise = AutoColor_ModelInstance(data, t_tensor)

            sqrt_alpha_bar = torch.sqrt(diffusion.alphas_bars[t])
            sqrt_one_minus = torch.sqrt(1 - diffusion.alphas_bars[t])
            noisy_rgb = (noisy_rgb - sqrt_one_minus * predicted_noise) / sqrt_alpha_bar

            # Clamp values to valid range
            noisy_rgb = torch.clamp(noisy_rgb, -1, 1)
        
        denoised = noisy_rgb
        denoised_np = denoised.cpu().permute(0, 2, 3, 1).squeeze(0).numpy()

        # Normalize from [-1,1] to [0,255]
        denoised_np = ((denoised_np + 1) / 2 * 255).astype("uint8")

        # Convert to PIL Image
        final_image = Image.fromarray(denoised_np)
        
        # Resize to original dimensions
        final_image = final_image.resize((width, height), Image.BILINEAR)

        # Save to byte stream
        bytes_io = BytesIO()
        final_image.save(bytes_io, format="JPEG")
        bytes_io.seek(0)

        return Response(content=bytes_io.getvalue(), media_type="image/jpeg")


if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=8000)

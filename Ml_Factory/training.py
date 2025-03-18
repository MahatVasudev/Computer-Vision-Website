import torch 

import torch.nn as nn
import glob
import matplotlib.pyplot as plt
from torchvision.transforms import transforms
from Datasets.datasets import AnimeAllImages, LandScapeColoredDataset
from models.torch_model import ModelOneGen, ModelOneGenX
import cv2
from PIL import Image


if __name__ == "__main__":

    model = ModelOneGenX().to('cuda')

    weights = torch.load("./products/CheckPoint_Epoch_2_ModelOneGenX4.pth", weights_only=True)
    model.load_state_dict(weights["model_params"])
    
    images = glob.glob("D:/Image Dataset/data/train_black/*.jpg")
        
    transform = transforms.Compose([
        transforms.Resize((256,256)),
        transforms.ToTensor()
        ])

    image = Image.open(images[10]).convert("L")
    
    image = transform(image).unsqueeze(0).to('cuda')
    model.eval()

    predicted_image = model(image)
    
    predicted_image = predicted_image.squeeze(0).detach().cpu().numpy().transpose(1,2,0)
    plt.imshow(predicted_image)
    plt.show()


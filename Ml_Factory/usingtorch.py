from typing import Any
import numpy as np
import cv2
import glob
import torch 
import torch.nn as nn
from torch.nn.modules import ReLU, Upsample
import torchvision
import torchvision.models as torch_models
from torch.utils.data import DataLoader, Dataset


class ModelOne(nn.Module):

    def __init__(self):
        super(ModelOne, self).__init__()

        self.fc2_feature = torch_models.vgg16(pretrained=True).classifier[5]

        self.image_dense = nn.Sequential(
                nn.Dropout(0.5),
                nn.Linear(4096, 1024),
                nn.ReLU()
                )
        self.encoder = nn.Sequential(
                nn.Conv2d(3, 64, kernel_size=3, padding=1, stride=2),
                nn.ReLU(),
                nn.Conv2d(64, 128, kernel_size=3, padding=1),
                nn.ReLU(),
                nn.Conv2d(128,128, kernel_size=3, padding=1, stride=2),
                nn.ReLU(),
                nn.Conv2d(128,256, kernel_size=3, padding=1, stride=2),
                nn.ReLU(),
                nn.Conv2d(256,256, kernel_size=3, padding=1, stride=2),
                nn.ReLU(),
                nn.Conv2d(256,512, kernel_size=3, padding=1, stride=2),
                nn.ReLU(),
                nn.Dropout(0.3),
                nn.Conv2d(512,512, kernel_size=3, padding=1, stride=2),
                nn.ReLU(),
                nn.Conv2d(512,256, kernel_size=3, padding=1, stride=2),
                nn.ReLU()
                )

        self.fusion_conv = nn.Conv2d(256 + 1024, 128,kernel_size=3, padding=1)

        self.decoder = nn.Sequential(
                nn.Upsample(scale_factor=2),
                nn.Conv2d(128, 64, kernel_size=3, padding=1),
                nn.ReLU(),
                nn.Upsample(scale_factor=2),
                nn.Conv2d(64, 32, kernel_size=3, padding=1),
                nn.ReLU(),
                nn.Conv2d(32, 16, kernel_size=3, padding=1),
                nn.ReLU(),
                nn.Conv2d(16, 2, kernel_size=3, padding=1),
                nn.Tanh(),
                nn.Upsample(scale_factor=2)
            )

    def forward(self, xx: torch.Tensor):

        x1: torch.Tensor = self.fc2_feature(xx)
        x1  = self.image_dense(x1)

        x2: torch.Tensor = self.encoder(xx)

        x1 = x1.squeeze(-1).unsqueeze(-1).repeat(1, 1, x2.size(2), x2.size(3))

        x: torch.Tensor = torch.cat((x2,x1), dim=1)

        x = self.fusion_conv(x)

        x = self.decoder(x)

        return x


model = ModelOne()

import os

os.chdir("D:/Image Dataset/")

def data_generator(image_dir: str) -> tuple[torch.Tensor, torch.Tensor]:
    
    data: list = glob.glob("./data/train_color/*.jpg")

    for d in data:
        image = cv2.imread(d)

        lab = cv2.cvtColor(image, cv2.COLOR_BGR2LAB)
        
        l, ab = lab[:,:,0], lab[:,:,1:]

        l_3x = np.stack(l,axis=-1)

        yield (torch.from_numpy(l, dtype=torch.float64), 
               torch.from_numpy(ab, dtype=torch.float64))


from PIL import Image
from torch.utils.data import DataLoader, Dataset
import torchvision.transforms as transforms
import glob
import os


class CustomImageDataset(Dataset):
    def __init__(self, 
                 root_dir: str, 
                 transform = None, 
                 extension : list[str] = ["jpg"], all_images: bool = False):
        
        self.images = []
        for extensions in extension:
            if all_images == True:
                self.images.extend(glob.glob(os.path.join(root_dir,f"**/*.{extensions}")))
            else:
                self.images.extend(glob.glob(os.path.join(root_dir,f"*.{extensions}")))
        self.extension = extension
        self.transform = transform
    
    def __len__(self):
        return len(self.images)

    def __getitem__(self, idx: int):
        image_path = self.images[idx]
        image = Image.open(image_path).convert("RGB")
        gray_scale = image.convert("L")

        gray_tensor = transforms.ToTensor()(gray_scale)
        # gray_tensor = (gray_tensor / 127.5)-1
        image_tensor = transforms.ToTensor()(image)
        # image_tensor = (image_tensor / 127.5) - 1
        if self.transform != None:
            image_tensor = self.transform(image_tensor)
            gray_tensor = self.transform(gray_tensor)
        return gray_tensor, image_tensor

if __name__ == "__main__":

    images = CustomImageDataset("D:/Image Dataset/data/train_color")
    print("Length of Image Dataset", len(images))
    loader = DataLoader(images,batch_size=8,num_workers=4)
    print(next(iter(loader)))
    


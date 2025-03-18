from .custom_image_dataset import CustomImageDataset


LandScapeColoredDataset = lambda transform = None: CustomImageDataset("D:/Image Dataset/landscape Images/color", extension=["jpg"], transform=transform)

AnimeAllImages = lambda transform = None: CustomImageDataset("D:/Image Dataset/Anime Images/", extension=["jpg","png"], all_images=True, transform=transform)

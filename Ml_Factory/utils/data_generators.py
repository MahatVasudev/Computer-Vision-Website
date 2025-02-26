import glob
import numpy as np
from pickle import load
import os
import cv2

from .convert_color import convertRGBtoGrayYCrCb, convertRGBtoYCrCb

def data_generator_merge(image_dir, num_train_samples, batch_size, start = 0):
    # loop for ever over images
    image_list = glob.glob(image_dir + "*.jpg")
    current_batch_size=0
    while 1:
         #'coco_images\processed'
        for file_idx in range(start, num_train_samples):
            # retrieve the photo feature
            if current_batch_size == 0:
                X1, X2, Y = list(), list(), list()

            
            text = image_list[file_idx].replace(image_dir[:-1] + "\\", "").split(".")[0]
            img = cv2.imread(image_list[file_idx])
            img_arr_rgb = cv2.cvtColor(img, cv2.COLOR_BGR2RGB) #change to np array
            img_arr_ycc = convertRGBtoYCrCb(img_arr_rgb)
            img_arr_ycc_gray = convertRGBtoGrayYCrCb(img_arr_rgb) 
            
            file_feature = os.getcwd() + '\\vgg16\\preprocessed\\' + text + '.pk'
            fid = open(file_feature, 'rb')
            fc2_features = load(fid)['fc2_features']
            fid.close()
            
            img_arr_ycc_crcb = img_arr_ycc[:,:,1:] #for loss
            fc2_features = fc2_features/(np.max(fc2_features)-np.min(fc2_features)+1e-8)
            # img_arr_ycc_gray_expandD = np.expand_dims(img_arr_ycc_gray, axis=-1)
            # img_arr_ycc_gray_expandD = np.repeat(img_arr_ycc_gray_expandD, 3, axis=-1)
            # if img_arr_rgb
            img_arr_ycc_gray_expandD = np.stack([img_arr_ycc_gray]*3, axis=-1)

            
            
            X1.append(fc2_features)
            X2.append(img_arr_ycc_gray_expandD/255)
            Y.append(img_arr_ycc_crcb/255)
            
            current_batch_size += 1
            if file_idx == 5738:
                remaining = batch_size - current_batch_size
                for i in range(remaining):
                    X1.append(fc2_features)
                    X2.append(img_arr_ycc_gray_expandD/255)
                    Y.append(img_arr_ycc_crcb/255)
                current_batch_size = batch_size
            if current_batch_size == batch_size:
                current_batch_size = 0
                # print([[np.squeeze(np.array(X1)).shape, np.array(X2).shape], np.array(Y).shape])
                yield [[np.squeeze(np.array(X1, dtype=np.float32)), np.array(X2, dtype=np.float32)], np.array(Y, dtype=np.float32)]

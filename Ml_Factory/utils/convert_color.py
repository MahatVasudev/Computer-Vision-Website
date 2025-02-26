# converting color formats
import numpy as np


def convertRGBtoYCrCb(arr):
    xform = np.array([[0.299, 0.587, 0.114], [-.1687, -.3313, .5], [.5, -.4187, -.0813]])
    ycc = arr.dot(xform.T)
    ycc[:,:,[1,2]] += 128
    return np.uint8(ycc)

#converting YCrCb back to RGB
def convertingYCrCbtoRGB(arr):
    xform = np.array([[1,0,1.402], [1,-0.34414, -0.71414], [1, 1.772, 0]])
    rgb = arr.astype(np.float32)
    rgb[:,:,[1,2]] -=128
    rgb = rgb.dot(xform.T)
    np.putmask(rgb, rgb > 255, 255)
    np.putmask(rgb, rgb < 0, 0)
    return np.uint8(rgb)

def convertRGBtoGrayYCrCb(arr):
    ycc = convertRGBtoYCrCb(arr)
    ret = np.zeros_like(ycc[:,:,0])
    ret[:,:] = ycc[:,:,0]
    return ret

def convertYCCtoGrayRGB(arr):
    ret = np.zeros_like(arr)
    ret[:, :, 0] = arr[:, :, 0]
    ret[:, :, 1] = arr[:, :, 0]
    ret[:, :, 2] = arr[:, :, 0]
    return ret

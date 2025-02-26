from keras.models import Input, Conv2D, Dense, Dropout, UpSampling2D, concatenate
from keras.models import Model
import keras.backend as K
from keras.optimizers import SGD
import tensorflow as tf
from keras.models import Model
from keras.layers import (
    Conv2D, UpSampling2D, Input, Dropout, Dense, RepeatVector, Reshape, concatenate, GlobalAveragePooling2D
)
from utils.repeat_and_reshape import repeat_and_reshape
from keras.utils import plot_model

def ModelOne(feature_size):
    """
    Model One Uses Features used from VGG16 which makes it a merged model
    it has two inputs 
        1: Feauters of image extracted by VGG16
        2: Gray Scale Image stacked in three dimentions
    Outputs: AB in case if LAB model or CrCb if YCrCb
    """
    inputs1 = Input(shape=(feature_size,))
    image_feature = Dropout(0.5)(inputs1)
    image_feature = Dense(1024, activation='relu')(image_feature)
    
    inputs2 = Input(shape=(None, None, 3,))  # <--- Change to flexible input size
    encoder_output = Conv2D(64, (3,3), activation='relu', padding='same', strides=2)(inputs2)
    encoder_output = Conv2D(128, (3,3), activation='relu', padding='same')(encoder_output)
    encoder_output = Dropout(0.3)(encoder_output)
    encoder_output = Conv2D(128, (3,3), activation='relu', padding='same', strides=2)(encoder_output)
    encoder_output = Conv2D(256, (3,3), activation='relu', padding='same')(encoder_output)
    encoder_output = Dropout(0.3)(encoder_output)
    
    encoder_output = Conv2D(256, (3,3), activation='relu', padding='same', strides=2)(encoder_output)
    encoder_output = Conv2D(512, (3,3), activation='relu', padding='same')(encoder_output)
    encoder_output = Dropout(0.3)(encoder_output)
    encoder_output = Conv2D(512, (3,3), activation='relu', padding='same')(encoder_output)
    encoder_output = Conv2D(256, (3,3), activation='relu', padding='same')(encoder_output)
    
    image_feature = tf.keras.layers.Lambda(repeat_and_reshape)([inputs1,encoder_output])

    fusion_output = concatenate([encoder_output, image_feature], axis=-1)
    
    decoder_output = Conv2D(128, (3,3), activation='relu', padding='same')(fusion_output)
    decoder_output = UpSampling2D((2, 2))(decoder_output)
    decoder_output = Conv2D(64, (3,3), activation='relu', padding='same')(decoder_output)
    decoder_output = UpSampling2D((2, 2))(decoder_output)
    decoder_output = Conv2D(32, (3,3), activation='relu', padding='same')(decoder_output)
    decoder_output = Conv2D(16, (3,3), activation='relu', padding='same')(decoder_output)
    decoder_output = Conv2D(2, (3,3), activation='tanh', padding='same')(decoder_output)
    decoder_output = UpSampling2D((2, 2))(decoder_output)

    model = Model(inputs=[inputs1, inputs2], outputs=decoder_output)
    model.compile(optimizer=SGD(learning_rate=0.01,momentum=0.9), loss='mean_squared_error', metrics=['acc'])

    print(model.summary())
    # plot_model(model, to_file='autoencoder_colorization_merged.png', show_shapes=True)
    
    return model

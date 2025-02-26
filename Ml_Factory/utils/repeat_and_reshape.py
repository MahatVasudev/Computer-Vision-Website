import tensorflow as tf

def repeat_and_reshape(args: tf.Tensor):
    inputs, encoder_output = args
    encoder_shape = tf.shape(encoder_output)
    batch_size = encoder_shape[0]
    height = encoder_shape[1]
    width = encoder_shape[2]

    repeated = tf.repeat(inputs, height * width, axis=1)  # Repeat features
    return tf.reshape(repeated, (batch_size, height, width, tf.shape(inputs)[1]))

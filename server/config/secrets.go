package config

type SecretConfig struct {
	UserEncrypt []byte
	OTPEncrpt   []byte
}

func DefaultConfig() *SecretConfig {
	return &SecretConfig{
		UserEncrypt: getEnvByte("USER_ENCRYPT_TOKEN", "secretsofsixteen"),
		OTPEncrpt:   getEnvByte("OTP_ENCRYPT_TOKEN", "secretsofsixtean"),
	}
}

var SecretEnvs = DefaultConfig()

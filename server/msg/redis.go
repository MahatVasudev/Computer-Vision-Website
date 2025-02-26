package msg

import "fmt"

func UserSessionKeys(key string) string {
	return fmt.Sprintf("U:<%s>", key)
}

func OTPSessionKeys(key string) string {
	return fmt.Sprintf("OTP:<%s>", key)
}

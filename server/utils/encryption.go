package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(data interface{}, secret_text []byte) (string, error) {

	plain_text, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secret_text)

	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)

	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	cipherText := make([]byte, len(plain_text))

	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText, plain_text)

	final_data := append(iv, cipherText...)
	return base64.StdEncoding.EncodeToString(final_data), nil
}

func Decrypt(encrpted_text string, secret_text []byte) (map[string]interface{}, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encrpted_text)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(secret_text)

	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, fmt.Errorf("Text is too small")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	decrypted_text := make([]byte, len(cipherText))

	cipher.NewCFBDecrypter(block, iv).XORKeyStream(decrypted_text, cipherText)

	decrypted_text = bytes.Trim(decrypted_text, "\x00")

	var data map[string]interface{}

	if err := json.Unmarshal(decrypted_text, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), err
}

func ComparePassword(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)

	return err == nil
}

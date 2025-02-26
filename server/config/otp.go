package config

import (
	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
)

type EmailSenderConfig struct {
	Identity string
	EMAIL    string
	PASSWORD string
}

func CreateNewEmailSender(
	identity string,
	email_env string,
	password_env string,
) (*EmailSenderConfig, error) {
	email := getEnv(email_env, "")
	password := getEnv(password_env, "")
	if email == "" || password == "" {
		return nil, msg.ErrorNotFound
	}

	return &EmailSenderConfig{Identity: identity, EMAIL: email, PASSWORD: password}, nil

}

func NewFoodEmailSender() *EmailSenderConfig {
	return &EmailSenderConfig{
		Identity: "OnSight",
		EMAIL:    getEnv("FOOD_EMAIL", "unknown"),
		PASSWORD: getEnv("FOOD_PASSWORD", "unknown"),
	}
}

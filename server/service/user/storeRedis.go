package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/config"
	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
	"github.com/MahatVasudev/Computer-Vision-Website/server/utils"
)

// Store User Info
// Format key = U:<Session ID>
// Format Value = {
//  User ID
//  Username
//  Email
//  IP
//  Logged In
//}

// read user key if exist
func (s *Store) ReadUserKey(ctx context.Context, key string) (*typestore.Redis_UserSession, error) {

	var payload typestore.Redis_UserSession

	err := utils.ReadEncryptedRedisDataToAStruct(ctx,
		msg.UserSessionKeys(key),
		&payload,
		config.SecretEnvs.UserEncrypt, s.RedDb)

	if err != nil {

		return nil, err
	}

	log.Println(payload)

	return &payload, nil

}

func (s *Store) DeleteOTPKey(ctx context.Context, key string) error {
	err := s.RedDb.Del(ctx, msg.OTPSessionKeys(key))

	if err.Err() != nil {
		return err.Err()
	}

	return nil
}

func (s *Store) ReadOTPKey(ctx context.Context, key string) (*typestore.OTP_Redis, error) {
	encryptedText := s.RedDb.Get(ctx, msg.OTPSessionKeys(key))

	if encryptedText.Val() == "" {
		return nil, fmt.Errorf("data not found")
	}

	decryptedText, err := utils.Decrypt(encryptedText.Val(), config.SecretEnvs.OTPEncrpt)

	if err != nil {
		return nil, err
	}

	var payload typestore.OTP_Redis
	if err := utils.MapToStruct(decryptedText, &payload); err != nil {

		return nil, err
	}

	return &payload, nil
}

func (s *Store) CreateOTPSession(
	ctx context.Context,
	key string,
	otp int,
	duration time.Duration,
) error {
	encrypted, err := utils.Encrypt(typestore.OTP_Redis{OTP: otp}, config.SecretEnvs.OTPEncrpt)

	if err != nil {
		return err
	}

	if err := s.RedDb.Set(ctx, msg.OTPSessionKeys(key), encrypted, duration); err != nil {
		return err.Err()
	}

	return nil
}

// Create 'key':'value'
func (s *Store) CreateUserSession(
	ctx context.Context,
	key string,
	user typestore.Redis_UserSession,
	duration time.Duration,
) error {

	encrypted, err := utils.Encrypt(user, config.SecretEnvs.UserEncrypt)

	if err != nil {
		return err
	}

	if err := s.RedDb.Set(ctx, msg.UserSessionKeys(key), encrypted, duration); err != nil {
		return err.Err()
	}

	return nil
}

// Update 'value'

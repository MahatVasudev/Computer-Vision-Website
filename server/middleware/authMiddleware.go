package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/config"
	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/utils"
	"github.com/MahatVasudev/Computer-Vision-Website/server/writer"
)

func AuthMiddleWareUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		defer cancel()

		token, err := r.Cookie(msg.SSID)

		if err != nil {
			writer.WriteNotAuthorized(w, msg.ErrorUnAuthorized)
			return
		}

		ssid, err := utils.GetRedisSSID(ctx, token.Value)

		if err != nil {
			writer.WriteNotAuthorized(w, msg.ErrorUnAuthorized)
			return
		}

		log.Println(utils.Decrypt(ssid, config.SecretEnvs.UserEncrypt))

		next.ServeHTTP(w, r)

	})
}

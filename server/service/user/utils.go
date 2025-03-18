package user

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
	"github.com/MahatVasudev/Computer-Vision-Website/server/utils"
	"github.com/MahatVasudev/Computer-Vision-Website/server/writer"
)

func (h *Handler) CreateSessionOfUser(
	w http.ResponseWriter,
	ctx context.Context,
	u typestore.User,
) {

	session_id := createSessionId()

	if err := h.store.CreateUserSession(
		ctx,
		session_id,
		typestore.Redis_UserSession{
			Id:       u.ID,
			Username: u.Username,
			IP:       "Some",
			LoggedIn: time.Now(),
		},
		time.Hour*9,
	); err != nil {
		writer.WriteServerError(w, err)
		return
	}

	writer.WriteCookieSSID(w, session_id)
}

func (h *Handler) GetSessionData(
	ctx context.Context,
	r *http.Request,

) (*typestore.Redis_UserSession, error) {

	ssid, err := r.Cookie(msg.SSID)

	if err != nil {
		return nil, fmt.Errorf("Not Authorized")
	}

	session_data, err := h.store.ReadUserKey(ctx, ssid.Value)
	log.Println(session_data)
	if err != nil {
		return nil, fmt.Errorf("Not Authorized")
	}

	return session_data, nil
}

func createSessionId() string {
	return utils.RandStringBytesMaskImprSrcSB(50)
}

package follows

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/payloads"
	"github.com/MahatVasudev/Computer-Vision-Website/server/utils"
	"github.com/MahatVasudev/Computer-Vision-Website/server/writer"
)

func (h *Handler) HandleFollow(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)

	defer cancel()

	ssid_key, _ := r.Cookie(msg.SSID)

	sess_id, _ := h.userStore.ReadUserKey(ctx, ssid_key.Value)

	var payload payloads.FollowOrUnFollow

	if err := utils.ParseJson(r, &payload); err != nil {
		writer.WriteBadRequest(w, msg.ErrorBadRequest)
		return
	}

	if sess_id.Id == payload.Following_id {
		writer.WriteConflict(w, fmt.Errorf("User Cannot Follow Itself"))
		return
	}

	err := h.store.FollowSomeOne(ctx, sess_id.Id, payload.Following_id)

	if err != nil {
		if err == msg.ErrorNotFound {
			writer.WriteNotFound(w, err)
			return
		} else if err == msg.ErrorConflict {
			writer.WriteConflict(w, err)
			return
		}

		writer.WriteServerError(w, err)
		return
	}

	writer.WriteOk(w, "Followed!")
}

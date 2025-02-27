package user

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/MahatVasudev/Computer-Vision-Website/server/auth"
	"github.com/MahatVasudev/Computer-Vision-Website/server/db"
	"github.com/MahatVasudev/Computer-Vision-Website/server/msg"
	"github.com/MahatVasudev/Computer-Vision-Website/server/payloads"
	response_success "github.com/MahatVasudev/Computer-Vision-Website/server/responses/success"
	"github.com/MahatVasudev/Computer-Vision-Website/server/typestore"
	"github.com/MahatVasudev/Computer-Vision-Website/server/utils"
	"github.com/MahatVasudev/Computer-Vision-Website/server/writer"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) handleUserNameDetails(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	session_data, err := h.GetSessionData(ctx, r)

	if err != nil {
		writer.WriteNotAuthorized(w, err)
		return
	}

	sess_username := session_data.Username

	target_username := chi.URLParam(r, "username")

	data, err := h.store.GetUserByUserName(ctx, target_username)

	if err != nil {
		if err == msg.ErrorNotFound {
			writer.WriteNotFound(w, err)
			return
		}

		writer.WriteServerError(w, msg.ErrorServerSide)
		return
	}

	permissions := []string{}

	if sess_username == target_username {
		permissions = []string{
			"permission:edit",
			"permission:analytics",
			"permission:delete",
			"permission:add",
		}
	}

	writer.WriteOk(w, map[string]interface{}{
		"user":        data,
		"permissions": permissions,
	})

}

func (h *Handler) handleSetup(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) SeeUsernameExists(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*6)

	defer cancel()

	var payload temp_username

	if err := utils.ParseJson(r, &payload); err != nil {
		writer.WriteBadRequest(w, err)
		return
	}

	if _, err := h.store.GetUserByUserName(ctx, payload.Username); err != nil {
		if err == msg.ErrorNotFound {
			writer.WriteNotFound(w, err)
			return
		}

		writer.WriteServerError(w, err)
		return
	}

	writer.WriteOk(w, "User Found")
}

func (h *Handler) SeeEmailExists(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*6)

	defer cancel()

	var payload temp_email

	if err := utils.ParseJson(r, &payload); err != nil {
		writer.WriteBadRequest(w, err)
		return
	}

	if _, err := h.store.GetUserByEmail(ctx, payload.Email); err != nil {
		if err == msg.ErrorNotFound {
			writer.WriteNotFound(w, err)
			return
		}

		writer.WriteServerError(w, err)
		return
	}

	writer.WriteOk(w, "User Found")
}

func (h *Handler) handleVerifyOTP(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var payload temp_email_otp

	if err := utils.ParseJson(r, &payload); err != nil {
		writer.WriteBadRequest(w, err)
		return
	}

	otpToken, err := h.store.ReadOTPKey(ctx, payload.Email)

	if err != nil {
		writer.WriteServerError(w, err)
		return
	}

	if otpToken.OTP != payload.Otp_code {
		writer.WriteNotAuthorized(w, msg.ErrorUnAuthorized)
		h.store.DeleteOTPKey(ctx, payload.Email)
		return
	}

	val := db.RedisInstance.Get(ctx, msg.OTPSessionKeys(payload.Email))

	if val.Err() != nil {
		writer.WriteServerError(w, msg.ErrorServerSide)
		return
	}

	writer.WriteOk(w, map[string]string{
		"handshake": val.Val(),
	})
}

func (h *Handler) handleVerifyEmail(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	emailsender, err := auth.TestEmailSender()

	if err != nil {
		writer.WriteServerError(w, msg.ErrorServerSide)
		return
	}

	var payload temp_email

	if err := utils.ParseJson(r, &payload); err != nil {
		writer.WriteBadRequest(w, err)
		return
	}

	otp_value := 8500*rand.Float32() + 1000
	otp := int(math.Floor(float64(otp_value)))

	emailsender.SendOTP(payload.Email, otp)

	h.store.CreateOTPSession(ctx, payload.Email, otp, time.Minute*3)

	writer.WriteOk(w, "Email Sent Successfully!")
}

func (h *Handler) handleGetAllData(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	users, err := h.store.GetAllUsers(ctx)

	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(res, http.StatusOK, response_success.DataMessage(users))
	return
}

func (h *Handler) handleCreateUser(res http.ResponseWriter, req *http.Request) {

	var payload payloads.CreateUser
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()
	if err := utils.ParseJson(req, &payload); err != nil {
		writer.WriteServerError(res, err)
		return
	}

	defer h.store.DeleteOTPKey(ctx, payload.Email)

	val := db.RedisInstance.Get(ctx, msg.OTPSessionKeys(payload.Email))

	if val.Err() != nil {
		writer.WriteNotAuthorized(res, msg.ErrorUnAuthorized)
		return
	}

	if val.Val() != payload.Handshake {
		writer.WriteNotAuthorized(res, msg.ErrorUnAuthorized)
		return
	}

	hashedpassword, _ := utils.HashPassword(payload.Password)
	u := typestore.User{
		ID:        uuid.NewString(),
		Username:  "JaneDoe_" + strconv.Itoa(int(rand.Int31n(1000)+500)),
		FirstName: "Anonymous",
		Email:     payload.Email,
		Password:  hashedpassword,
	}

	err := h.store.CreateUser(ctx, u)

	if err != nil {
		if err == msg.ErrorAlreadyExists {
			writer.WriteConflictError(res, err)
			return
		}

		return
	}

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
		writer.WriteServerError(res, err)
		return
	}

	writer.WriteCookieSSID(res, session_id)

	writer.WriteOk(res, u)

}

func (h *Handler) handleLoginUser(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var payload payloads.LoginUser

	if err := utils.ParseJson(r, &payload); err != nil {
		writer.WriteBadRequest(w, err)
		return
	}

	u, err := h.store.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		if err == msg.ErrorNotFound {
			writer.WriteNotFound(w, err)
			return
		}

		writer.WriteServerError(w, err)
		return
	}

	if ok := utils.ComparePassword(u.Password, []byte(payload.Password)); !ok {
		writer.WriteNotFound(w, fmt.Errorf("username or password not found"))
		return
	}

	h.CreateSessionOfUser(w, ctx, *u)

	writer.WriteOk(w,
		map[string]string{
			"status":   "Successful",
			"username": u.Username,
		})
}

func (h *Handler) handleVerifyLogin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	session_data, err := h.GetSessionData(ctx, r)

	if err != nil {
		writer.WriteNotAuthorized(w, err)
		return
	}

	writer.WriteOk(w, map[string]string{
		"username": session_data.Username,
		"id":       session_data.Id,
	})

}

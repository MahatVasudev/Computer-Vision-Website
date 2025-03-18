package user

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

func (h *Handler) handleUsersPosts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	target_username := chi.URLParam(r, "username")

	limit := 10

	posts, err := h.postStore.Get_All_Posts_From_User(ctx, target_username, limit)

	if err != nil {
		if err == msg.ErrorNotFound {
			writer.WriteOk(w, &posts)
			return
		}
		writer.WriteServerError(w, err)
		return
	}

	writer.WriteOk(w, &posts)

}

func (h *Handler) handleUserNameDetails(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	session_data, err := h.GetSessionData(ctx, r)
	user_name := "*"
	if err == nil {
		user_name = session_data.Username
	}

	target_username := chi.URLParam(r, "username")
	println(target_username)
	if target_username == "" || len(target_username) < 5 || len(target_username) > 20 {
		writer.WriteBadRequest(w, msg.ErrorBadRequest)
		return
	}

	// Channels
	userChan := make(chan *typestore.User, 1)
	followChan := make(chan *typestore.FollowingAggDetails, 1)
	isFollowingChan := make(chan bool, 1)
	isFollowedByChan := make(chan bool, 1)
	errChan := make(chan error, 3)
	postCountChan := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(4)

	// ✅ Get user details
	go func() {
		defer wg.Done()
		user, err := h.store.GetUserByUserName(ctx, target_username)
		if err != nil {
			log.Println("Error fetching user:", err)
			errChan <- err
			userChan <- nil // ✅ Ensure channel isn't blocked
			return
		}
		log.Println("User Data: ", user)
		userChan <- user
	}()

	// ✅ Get follower/following count
	go func() {
		defer wg.Done()
		data, err := h.followStore.GetAggregate(ctx, target_username)
		if err != nil {
			log.Println("Error fetching follow data:", err)
			errChan <- err
			followChan <- nil
			return
		}
		followChan <- data
	}()

	// ✅ Check if user is following/followed
	go func() {
		defer wg.Done()

		isFollowing, isFollowedBy, err := h.followStore.IsFollowingOrFollowed(
			ctx,
			user_name,
			target_username,
		)
		if err != nil {
			log.Println("Error checking follow status:", err)
			errChan <- err
		}

		isFollowingChan <- isFollowing
		isFollowedByChan <- isFollowedBy
	}()

	go func() {
		defer wg.Done()

		posts_count, _ := h.postStore.CountOfPostOfEachUser(ctx, target_username)

		// if err != nil {
		// 	log.Println("Error checking post count: ", err)
		// 	errChan <- err
		// }

		postCountChan <- posts_count

	}()
	// ✅ Wait for goroutines & close channels safely
	go func() {
		wg.Wait()
		close(userChan)
		close(followChan)
		close(isFollowingChan)
		close(isFollowedByChan)
		close(postCountChan)
		close(errChan)
	}()

	// ✅ Handle errors properly AFTER all goroutines
	var userData *typestore.User
	var followerCount, followingCount int
	var isFollowing, isFollowedBy bool
	var postCount int
	// Handle errors properly after all goroutines finish
	select {
	case <-ctx.Done():
		writer.WriteReqTimeOut(w, ctx.Err())
		return

	case e := <-errChan:
		if e != nil {
			if e == msg.ErrorNotFound {
				writer.WriteNotFound(w, e)
				return
			}
			writer.WriteServerError(w, e)
			return
		}
	}

	// Read from channels safely
	if u := <-userChan; u != nil {
		userData = u
	} else {
		log.Println("Error Occured here")
		writer.WriteNotFound(w, msg.ErrorNotFound)
		return
	}

	if f, ok := <-followChan; ok && f != nil {
		followerCount = f.FollowerCount
		followingCount = f.FollowingCount
	}

	if isf, ok := <-isFollowingChan; ok {
		isFollowing = isf
	}

	if isfb, ok := <-isFollowedByChan; ok {
		isFollowedBy = isfb
	}

	if p_count, ok := <-postCountChan; ok {
		postCount = p_count
	}

	// ✅ Set permissions if user is requesting their own profile
	permissions := []string{}
	if strings.EqualFold(user_name, target_username) {
		permissions = []string{"profile:edit", "profile:analytics", "profile:delete"}
	}

	// ✅ Send response
	writer.WriteOk(w, map[string]interface{}{
		"user": userData,
		"follows": map[string]int{
			"followers": followerCount,
			"following": followingCount,
		},
		"post_count":   postCount,
		"isfollowing":  isFollowing,
		"isfollowedby": isFollowedBy,
		"permissions":  permissions,
	})
}

func (h *Handler) handleSetup(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)

	defer cancel()

	var payload payloads.SetupUser

	if err := utils.ParseJson(r, &payload); err != nil {
		writer.WriteBadRequest(w, msg.ErrorBadRequest)
		return
	}

	session_data, err := h.GetSessionData(ctx, r)

	if err != nil {
		writer.WriteNotAuthorized(w, err)
		return
	}

	sess_id := session_data.Id

	lastname := "*"

	if payload.LastName != nil {

		lastname = *payload.LastName
	}

	users := typestore.User{
		ID:        sess_id,
		FirstName: payload.FirstName,
		LastName:  &lastname,
		Username:  payload.Username,
		Email:     "",
	}

	users_details := typestore.UserDetails{
		CoverPhoto:    "/no_cover.jpeg",
		Avatar:        "/no_profile.jpeg",
		PreferedColor: payload.PreferedColor,
		Gender:        payload.Gender,
		DarkMode:      *payload.DarkMode,
		BirthYear:     payload.BirthYear,
	}

	errChan := make(chan error, 2)

	go func() {

		err := h.store.UpdateUserByID(ctx, users)

		errChan <- err

	}()

	go func() {

		err := h.store.UpdateORCreateUserDetailsByID(ctx, sess_id, users_details)

		errChan <- err

	}()

	var finalErr error
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			finalErr = ctx.Err()
		case err := <-errChan:
			finalErr = err
		}
	}

	// Handle errors
	if finalErr != nil {
		writer.WriteServerError(w, finalErr)
	} else {
		writer.WriteOk(w, "Setup completed successfully")

		if payload.Username != "*" {
			new_session_data := typestore.Redis_UserSession{
				Username: payload.Username,
				Id:       sess_id,
				IP:       session_data.IP,
				LoggedIn: session_data.LoggedIn,
			}

			key, _ := r.Cookie(msg.SSID)
			h.store.CreateUserSession(ctx, key.Value, new_session_data, time.Hour*8)
		}
	}

	close(errChan)
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

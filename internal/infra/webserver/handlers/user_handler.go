package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LucasLCabral/go-api/internal/dto"
	"github.com/LucasLCabral/go-api/internal/entity"
	"github.com/LucasLCabral/go-api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB       database.UserInterface
	JWT          *jwtauth.JWTAuth
	JWTExpiresIn int64
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, JwtExperiesIn int64) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		JWT:          jwt,
		JWTExpiresIn: JwtExperiesIn,
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
    http.Error(w, "invalid credentials", http.StatusUnauthorized)
    return
	}
	if !u.ValidatePassword(user.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	_, tokenString, err := h.JWT.Encode(map[string]interface{}{
		"sub":   u.ID.String(),
		"exp":   time.Now().Add(time.Duration(h.JWTExpiresIn) * time.Second).Unix(),
		"email": u.Email,
		"name":  u.Name,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	acessToken := struct {
		AcessToken string `json:"acess_token"`
	}{
		AcessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.UserDB.FindByEmail(user.Email)
	if err == nil {
		http.Error(w, "error: user already exist", http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "missing email parameter", http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

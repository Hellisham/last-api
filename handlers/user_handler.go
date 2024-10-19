package handlers

import (
	"encoding/json"
	"github.com/Hellisham/last-api/db"
	"github.com/Hellisham/last-api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing Password", http.StatusInternalServerError)
			return
		}

		var user = models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashPassword),
		}

		result := db.DB.Create(&user)
		if result.Error != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		userResponse := UserResponse{
			Name:  user.Name,
			Email: user.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userResponse)
	}
}

type LoginRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req_login LoginRequest

		err := json.NewDecoder(r.Body).Decode(&req_login)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		result := db.DB.Where("email = ?", req_login.Email).First(&models.User{})
		if result.Error != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		token, tokenerr := auth.GenerateToken(req_login.Name, req_login.Email)
	}

}

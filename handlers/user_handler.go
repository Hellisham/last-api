package handlers

import (
	"encoding/json"
	"github.com/Hellisham/last-api/auth"
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
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req_login LoginRequest
		var user models.User

		// Decode JSON request body into the LoginRequest struct
		err := json.NewDecoder(r.Body).Decode(&req_login)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Query the user from the database by name
		result := db.DB.Where("name = ?", req_login.Name).First(&user)
		if result.Error != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Compare the hashed password with the provided password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req_login.Password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Generate JWT token
		token, tokenErr := auth.JwtGnarator(user.Name, user.Email)
		if tokenErr != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to "application/json"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Send the token in the response
		json.NewEncoder(w).Encode(LoginResponse{Token: token})
	}
}

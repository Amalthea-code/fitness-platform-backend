package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dmitrijkrasikov/fitness-platform-backend/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB *sql.DB
}

// Register создает нового пользователя
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// проверяем уникальность username/email
	var exists bool
	err := h.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 OR username=$2)",
		input.Email, input.Username,
	).Scan(&exists)
	if err != nil {
		log.Println("check user exists failed:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "username or email already taken", http.StatusConflict)
		return
	}

	// хэшируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed to hash password:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// генерируем случайный session_token
	sessionToken := uuid.NewString()

	_, err = h.DB.Exec(
		"INSERT INTO users (username, email, password_hash, session_token) VALUES ($1, $2, $3, $4)",
		input.Username, input.Email, string(hash), sessionToken,
	)
	if err != nil {
		log.Println("failed to insert user:", err)
		http.Error(w, "could not create user", http.StatusInternalServerError)
		return
	}

	// ставим cookie для сессии
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		HttpOnly: true,
		Path:     "/",
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user created",
	})
}

// Login — авторизация пользователя
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	var user models.User
	err := h.DB.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE email=$1",
		input.Email,
	).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// создаем новый session_token
	sessionToken := uuid.NewString()
	_, err = h.DB.Exec(
		"UPDATE users SET session_token=$1 WHERE id=$2",
		sessionToken, user.ID,
	)
	if err != nil {
		log.Println("failed to update session_token:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// ставим cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		HttpOnly: true,
		Path:     "/",
	})

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	})
}

// Me — получить текущего пользователя по cookie
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var user models.User
	err = h.DB.QueryRow(
		"SELECT id, username, email FROM users WHERE session_token=$1",
		cookie.Value,
	).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Logout — удаление сессии
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// удаляем cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "logged out",
	})
}

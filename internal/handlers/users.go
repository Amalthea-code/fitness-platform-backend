package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dmitrijkrasikov/fitness-platform-backend/internal/models"
	"github.com/go-chi/chi/v5"
)

type UsersHandler struct {
	DB *sql.DB
}

// CreateUser — ручка для создания нового пользователя (только для админа/dev)
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	_, err := h.DB.Exec(
		"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)",
		input.Username, input.Email, input.Password, // пока без bcrypt для примера
	)
	if err != nil {
		log.Println("failed to insert user:", err)
		http.Error(w, "could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user created",
	})
}

// GetUser — получить одного пользователя по id
func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var user models.User
	err = h.DB.QueryRow(
		"SELECT id, username, email FROM users WHERE id=$1",
		id,
	).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			log.Println("failed to query user:", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

// ListUsers — получить всех пользователей
func (h *UsersHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, username, email FROM users ORDER BY id ASC")
	if err != nil {
		log.Println("failed to query users:", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
			log.Println("failed to scan user:", err)
			continue
		}
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

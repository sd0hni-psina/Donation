package handlers

import (
	"database/sql"
	"donation-app/database"
	"donation-app/models"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secretkey")

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var streamer models.Streamer
		err := json.NewDecoder(r.Body).Decode(&streamer)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		id, err := database.CreateStreamer(db, streamer)
		if err != nil {
			http.Error(w, "Failed to create streamer", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int64{"id": id})
	}
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var streamer models.Streamer
		err := json.NewDecoder(r.Body).Decode(&streamer)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		pass, err := database.GetStreamerByName(db, streamer.Name)
		if err != nil {
			http.Error(w, "Streamer not found", http.StatusNotFound)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(streamer.Password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"streamer_id": pass.ID,
			"exp":         time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	sr := http.NewServeMux()

	sr.HandleFunc("POST /register", Register(db))
	http.ListenAndServe(":8080", sr)
}

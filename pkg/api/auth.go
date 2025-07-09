package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type authRequest struct {
	Password string `json:"password"`
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Error reading request body: %v", err))
		return
	}

	var authReq authRequest

	if err = json.Unmarshal(buf.Bytes(), &authReq); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	realPassword := os.Getenv("TODO_PASSWORD")

	if authReq.Password != realPassword {
		writeError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	hash := sha256.Sum256([]byte(authReq.Password))
	hashHex := hex.EncodeToString(hash[:])
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(hashHex))

	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error signing token: %v", err))
		return
	}

	writeJson(w, http.StatusOK, authSuccessResponse{Token: tokenString})
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")

		if len(pass) > 0 {
			var jwtString string

			cookie, err := r.Cookie("token")

			if err == nil {
				jwtString = cookie.Value
			} else {
				http.Error(w, "Cookie not found", http.StatusUnauthorized)
				return
			}

			hash := sha256.Sum256([]byte(pass))
			jwtSecret := hex.EncodeToString(hash[:])

			token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("wrong signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil {
				http.Error(w, "Failed to parse token", http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

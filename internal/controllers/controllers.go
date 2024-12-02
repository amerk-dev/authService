package controllers

import (
	"authService/internal/store"
	"authService/pkg/generator"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type AccessResponse struct {
	GuId string `json:"gu_id"`
}

func AccessMethod(w http.ResponseWriter, r *http.Request) {
	var body AccessResponse
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	accessToken, err := generateAccessToken(body.GuId, r.RemoteAddr, time.Minute*5)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, hashedRefreshToken, err := generateRefreshToken()
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	err = store.StoreRefreshToken(body.GuId, hashedRefreshToken, r.RemoteAddr, "")
	if err != nil {
		http.Error(w, "Failed to store refresh token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func RefreshMethod(w http.ResponseWriter, r *http.Request) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	user, err := store.FindUserByRefreshToken(body.RefreshToken)
	if err != nil {
		http.Error(w, "Failed to find user", http.StatusNotFound)
		return
	}

	if user.Ip != r.RemoteAddr {
		// Отправляем на почту уведомление, что вйпишник другой
	}

	newAccessToken, err := generateAccessToken(user.Guid, r.RemoteAddr, time.Minute*5)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	newRefreshToken, newHashedRefreshToken, err := generateRefreshToken()
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	err = store.UpdateRefreshToken(user.Guid, newHashedRefreshToken, r.RemoteAddr)
	if err != nil {
		http.Error(w, "Failed to update refresh token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func generateAccessToken(guid, ip string, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"guid": guid,
		"ip":   ip,
		"exp":  time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateRefreshToken() (string, string, error) {
	refreshToken := generator.GenerateSecureToken(32)
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	//
	//encodedToken := base64.StdEncoding.EncodeToString([]byte(refreshToken))
	log.Println(refreshToken, hashedToken)
	return refreshToken, string(hashedToken), nil
}

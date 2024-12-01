package controllers

import (
	"authService/internal/store"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
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

	accessToken, err := generateAccessToken(body.GuId, r.RemoteAddr, time.Minute*5) // Ну вроде по заданию надо guid и ip
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, hashedRefreshToken, err := "", "", nil // ToDo потом сделаю
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	err = store.StoreRefreshToken(body.GuId, hashedRefreshToken)
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

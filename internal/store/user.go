package store

import (
	"authService/internal/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func StoreRefreshToken(guid, hashedToken, ip, email string) error {
	var user models.User

	res := Db.First(&user, "guid = ?", guid)
	if res.Error != nil {
		user = models.User{
			Guid:         guid,
			RefreshToken: hashedToken,
			Ip:           ip,
			Email:        email,
		}

		res = Db.Create(&user)
		if res.Error != nil {
			return errors.New("Failed to store refresh token")
		}
	} else {
		user.RefreshToken = hashedToken
		if ip != "" {
			user.Ip = ip
		}

		res = Db.Save(&user)
		if res.Error != nil {
			return errors.New("Failed to update refresh token")
		}
	}
	return nil
}

func FindUserByRefreshToken(refreshToken string) (*models.User, error) {
	var users []models.User

	res := Db.Find(&users)
	if res.Error != nil {
		return nil, errors.New("Failed to retrieve users")
	}

	for _, user := range users {
		err := bcrypt.CompareHashAndPassword([]byte(user.RefreshToken), []byte(refreshToken))
		if err == nil {
			return &user, nil
		}
	}

	return nil, errors.New("Failed to find user by refresh token")
}

func UpdateRefreshToken(guid, newHashedToken, ip string) error {
	var user models.User

	res := Db.First(&user, "guid = ?", guid)
	if res.Error != nil {
		return errors.New("Failed to find user by refresh token")
	} else {
		user.RefreshToken = newHashedToken
		if ip != "" {
			user.Ip = ip
		}
		res = Db.Save(&user)
		if res.Error != nil {
			return errors.New("Failed to update refresh token")
		}
		return nil
	}
}

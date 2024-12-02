package store

import (
	"authService/internal/models"
	"errors"
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
		if ip != user.Ip {
			// Отправляем на почту уведомление, что вйпишник другой
		}

		res = Db.Save(&user)
		if res.Error != nil {
			return errors.New("Failed to update refresh token")
		}
	}
	return nil
}

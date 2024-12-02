package store

import (
	"authService/internal/models"
	"errors"
)

func StoreRefreshToken(guid, hashedToken string) error {
	var user models.User

	res := Db.First(&user, "guid = ?", guid)
	if res.Error != nil {
		user = models.User{Guid: guid, RefreshToken: hashedToken}
		res = Db.Create(&user)
		if res.Error != nil {
			return errors.New("Failed to store refresh token")
		}
	} else {
		user.RefreshToken = hashedToken
		res = Db.Save(&user)
		if res.Error != nil {
			return errors.New("Failed to update refresh token")
		}
	}
	return nil
}

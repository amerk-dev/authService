package store

import (
	"authService/internal/models"
	"errors"
)

func StoreRefreshToken(guid, hashedToken string) error {
	user := Db.First(models.User{}, "guid = ?", guid)
	if user == nil {
		err := user.Create(models.User{Guid: guid, RefreshToken: hashedToken})
		if err != nil {
			return errors.New("Failed to store refresh token")
		}
	} else {
		// Update юзера
	}
	Db.Save(&user)
	return nil
}

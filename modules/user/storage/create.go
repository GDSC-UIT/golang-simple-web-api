package userstorage

import (
	usermodel "golang-simple-web-api/modules/user/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, data *usermodel.User) error {
	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

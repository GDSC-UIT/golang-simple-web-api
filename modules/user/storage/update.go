package userstorage

import (
	usermodel "golang-simple-web-api/modules/user/model"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB, id string, data *usermodel.ReqUpdateUser) error {
	if err := db.
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return err
	}

	return nil
}

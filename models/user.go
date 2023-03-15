package models

import (
	"errors"
	"gorm.io/gorm"
	"linhx1999.com/gin-blog/utils"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	IsAdmin  int    `gorm:"type:tinyint unsigned" json:"isAdmin"`
	//Avatar   string
}

func SaveUser(user *User) error {
	user.Password = utils.ScryptPassword(user.Password)
	err := db.Create(user).Error
	return err
}

func CountUserByName(username string) (count int64, err error) {
	err = db.Model(&User{}).Where("username = ?", username).
		Count(&count).Error
	return
}

func GetUserByName(username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).
		First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, err
}

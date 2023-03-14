package models

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
	//Avatar   string
}

func SaveUser(user *User) error {
	user.Password = ScryptPassword(user.Password)
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

func ScryptPassword(password string) string {
	salt := make([]byte, 8, 8)
	salt = []byte{19, 99, 07, 31, 03, 07, 06, 12}

	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}

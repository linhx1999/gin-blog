package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CountByName(name string) (count int64, err error) {
	err = db.Model(&Category{}).Where("name = ?", name).Count(&count).Error
	return
}

func Save(category *Category) error {
	return db.Create(category).Error
}

func Page(pageSize int, pageNum int) ([]Category, error) {
	var categories []Category
	err := db.Limit(pageSize).Offset((pageSize - 1) * pageNum).
		Find(&categories).Error
	return categories, err
}

func UpdateByID(id uint, category *Category) error {
	category.ID = id
	err := db.Model(category).Updates(category).Error
	return err
}

func RemoveByID(id uint) error {
	var category Category
	category.ID = id
	return db.Delete(&category).Error
}

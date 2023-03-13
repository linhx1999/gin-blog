package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CountCategoryByName(name string) (count int64, err error) {
	err = db.Model(&Category{}).Where("name = ?", name).Count(&count).Error
	return
}

func SaveCategory(category *Category) error {
	return db.Create(category).Error
}

func PageCategory(pageSize int, pageNum int) ([]Category, error) {
	var categories []Category
	err := db.Limit(pageSize).Offset((pageSize - 1) * pageNum).
		Find(&categories).Error
	return categories, err
}

func UpdateCategoryByID(id uint, category *Category) error {
	category.ID = id
	err := db.Model(category).Updates(category).Error
	return err
}

func RemoveCategoryByID(id uint) error {
	var category Category
	category.ID = id
	return db.Delete(&category).Error
}

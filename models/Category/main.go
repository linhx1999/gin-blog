package Category

import (
	"gorm.io/gorm"
	"linhx1999.com/gin-blog/models"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CountByName(name string) (count int64, err error) {
	err = models.DB.Model(&Category{}).Where("name = ?", name).
		Count(&count).Error
	return
}

func Save(category *Category) error {
	return models.DB.Create(category).Error
}

func Page(pageSize int, pageNum int) ([]Category, error) {
	var categories []Category
	err := models.DB.Limit(pageSize).Offset((pageSize - 1) * pageNum).
		Find(&categories).Error
	return categories, err
}

func UpdateByID(id uint, category *Category) error {
	category.ID = id
	err := models.DB.Model(category).Updates(category).Error
	return err
}

func RemoveByID(id uint) error {
	var category Category
	category.ID = id
	return models.DB.Delete(&category).Error
}

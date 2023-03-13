package models

import (
	"errors"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:CategoryID"`
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID int    `gorm:"type:int;not null" json:"categoryID"`
	Desc       string `gorm:"type:varchar(200)" json:"desc"`
	Content    string `gorm:"type:longtext" json:"content"`
	Img        string `gorm:"type:varchar(100)" json:"img"`
}

func SaveArticle(article *Article) error {
	return db.Create(article).Error
}

func PageArticle(pageSize int, pageNum int) ([]Article, error) {
	var articles []Article
	err := db.Preload("Category").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&articles).Error
	return articles, err
}

func GetArticleByID(id int) (*Article, error) {
	var article Article
	err := db.Preload("Category").
		Where("id = ?", id).
		First(&article).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return &Article{}, err
	}
	return &article, err
}

func PageArticlesInCategory(categoryID int, perPage int, page int) ([]Article, error) {
	var articles []Article
	err := db.Preload("Category").
		Limit(perPage).Offset((page-1)*perPage).
		Where("category_id = ?", categoryID).
		Find(&articles).Error
	return articles, err
}

func UpdateArticleByID(id uint, articles *Article) error {
	articles.ID = id
	err := db.Model(articles).Updates(articles).Error
	return err
}

func RemoveArticleByID(id uint) error {
	var articles Article
	articles.ID = id
	return db.Delete(&articles).Error
}

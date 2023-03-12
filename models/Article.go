package models

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignKey:CategoryID"`
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID int    `gorm:"type:int;not null" json:"categoryID"`
	Desc       string `gorm:"type:varchar(200)" json:"desc"`
	Content    string `gorm:"type:longtext" json:"content"`
	Img        string `gorm:"type:varchar(100)" json:"img"`
}

package Article

import (
	"gorm.io/gorm"
	"linhx1999.com/gin-blog/models/Category"
)

type Article struct {
	Category Category.Category `gorm:"foreignKey:CategoryID"`
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID int    `gorm:"type:int;not null" json:"categoryID"`
	Desc       string `gorm:"type:varchar(200)" json:"desc"`
	Content    string `gorm:"type:longtext" json:"content"`
	Img        string `gorm:"type:varchar(100)" json:"img"`
}

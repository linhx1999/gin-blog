package v1

import (
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models/Category"
	"linhx1999.com/gin-blog/utils/Result"
	"net/http"
)

func PostCategory(c *gin.Context) {
	var category Category.Category
	var err error

	if err = c.ShouldBindJSON(&category); err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(
				http.StatusBadRequest,
				err.Error(),
				"A0402",
				"",
			),
		)
		return
	}

	count, err := Category.CountByName(category.Name)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(
				http.StatusBadRequest,
				err.Error(),
				"A0402",
				"",
			),
		)
		return
	}
	if count > 0 {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(
				http.StatusBadRequest,
				err.Error(),
				"A0402",
				"分类已存在",
			),
		)
		return
	}

	err = Category.Save(&category)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(
				http.StatusBadRequest,
				err.Error(),
				"C0300",
				"分类保存失败",
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"保存成功",
			nil,
		),
	)
}

func GetCategories(c *gin.Context) {

}

func PutCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}

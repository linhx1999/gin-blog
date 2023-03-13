package v1

import (
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils/Result"
	"net/http"
)

func PostCategory(c *gin.Context) {
	var category models.Category
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

	count, err := models.CountByName(category.Name)
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
				"分类已存在",
				"A0402",
				"分类已存在",
			),
		)
		return
	}

	err = models.Save(&category)
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
			[]any{},
		),
	)
}

func GetCategories(c *gin.Context) {

}

func PutCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}

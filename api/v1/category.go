package v1

import (
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils/Result"
	"net/http"
	"strconv"
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
				"A0402", // 无效的用户输入
				"",
			),
		)
		return
	}

	count, err := models.CountCategoryByName(category.Name)
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

	err = models.SaveCategory(&category)
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
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(
				http.StatusBadRequest,
				err.Error(),
				"A0402", // 无效的用户输入
				"请输入正整数",
			),
		)
		return
	}

	pageNum, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(
				http.StatusBadRequest,
				err.Error(),
				"A0402", // 无效的用户输入
				"请输入正整数",
			),
		)
		return
	}

	data, err := models.PageCategory(pageSize, pageNum)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(
				http.StatusBadRequest,
				err.Error(),
				"",
				"",
			),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"保存成功",
			data,
		),
	)

}

func PutCategory(c *gin.Context) {

}

func DeleteCategory(c *gin.Context) {

}

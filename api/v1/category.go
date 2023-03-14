package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils/result"
	"net/http"
	"strconv"
)

func PostCategory(c *gin.Context) {
	var category models.Category
	var err error

	if err = c.ShouldBindJSON(&category); err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	count, err := models.CountCategoryByName(category.Name)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}
	if count > 0 {
		c.JSON(
			http.StatusBadRequest,
			result.New("分类已存在"),
		)
		return
	}

	err = models.SaveCategory(&category)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		result.New("新增分类成功"),
	)
}

func GetCategories(c *gin.Context) {
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			perPage = -1
		} else {
			c.JSON(
				http.StatusBadRequest,
				result.New(err.Error()),
			)
			return
		}
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			page = 2
		} else {
			c.JSON(
				http.StatusBadRequest,
				result.New(err.Error()),
			)
			return
		}
	}

	categories, err := models.PageCategory(perPage, page)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		result.New(
			"查找分类成功",
			categories,
		),
	)
}

func PutCategory(c *gin.Context) {
	var category models.Category

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	if err = c.ShouldBindJSON(&category); err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	count, err := models.CountCategoryByName(category.Name)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}
	if count > 0 {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	err = models.UpdateCategoryByID(uint(id), &category)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		result.New("修改分类成功"),
	)
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	err = models.RemoveCategoryByID(uint(id))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusNoContent,
		result.New("删除分类成功"),
	)
}

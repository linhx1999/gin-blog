package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils/Result"
	"net/http"
	"strconv"
)

func PostArticle(c *gin.Context) {
	var article models.Article
	var err error

	if err = c.ShouldBindJSON(&article); err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	err = models.SaveArticle(&article)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"新增文章成功",
			[]any{},
		),
	)
}

func GetArticles(c *gin.Context) {
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			perPage = -1
		} else {
			c.JSON(
				http.StatusBadRequest,
				Result.NewFail(http.StatusBadRequest, err.Error()),
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
				Result.NewFail(http.StatusBadRequest, err.Error()),
			)
			return
		}
	}

	articles, err := models.PageArticle(perPage, page)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"查找文章成功",
			articles,
		),
	)
}

func GetArticleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	article, err := models.GetArticleByID(id)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"查找文章成功",
			article,
		),
	)
}

func GetArticlesInCategory(c *gin.Context) {
	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			perPage = -1
		} else {
			c.JSON(
				http.StatusBadRequest,
				Result.NewFail(http.StatusBadRequest, err.Error()),
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
				Result.NewFail(http.StatusBadRequest, err.Error()),
			)
			return
		}
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	articles, err := models.PageArticlesInCategory(id, perPage, page)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"查找文章成功",
			articles,
		),
	)
}

func PutArticle(c *gin.Context) {
	var article models.Article

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	if err = c.ShouldBindJSON(&article); err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	err = models.UpdateArticleByID(uint(id), &article)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"修改文章成功",
			[]any{},
		),
	)
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	err = models.RemoveArticleByID(uint(id))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			Result.NewFail(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		Result.NewSuccess(
			"删除文章成功",
			[]any{},
		),
	)
}

package v1

import (
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils/result"
	"net/http"
)

func PostUser(c *gin.Context) {
	var user models.User
	var err error

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	count, err := models.CountUserByName(user.Username)
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
			result.New("用户已存在"),
		)
		return
	}

	err = models.SaveUser(&user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		result.New("新增用户成功"),
	)
}

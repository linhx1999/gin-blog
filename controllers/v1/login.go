package v1

import (
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils/result"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}

	_, err := models.GetUserByName(user.Username)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}
}

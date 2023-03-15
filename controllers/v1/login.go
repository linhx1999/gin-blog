package v1

import (
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/middlewares"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils"
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

	user2, err := models.GetUserByName(user.Username)
	if err != nil {
		if user2 == nil {
			c.JSON(
				http.StatusUnauthorized,
				result.New(err.Error()),
			)
			return
		} else {
			c.JSON(
				http.StatusBadRequest,
				result.New(err.Error()),
			)
			return
		}
	}
	if utils.ScryptPassword(user.Password) == user2.Password {
		if user2.IsAdmin == 1 {
			tokenString, err := middlewares.CreatToken(user.Username)
			if err != nil {
				c.JSON(
					http.StatusBadRequest,
					result.New(err.Error()),
				)
				return
			}

			c.JSON(
				http.StatusOK,
				result.New("管理员登陆成功", tokenString),
			)
		} else {
			c.JSON(
				http.StatusForbidden,
				result.New(http.StatusText(http.StatusForbidden)),
			)
		}
	} else {
		c.JSON(
			http.StatusUnauthorized,
			result.New(http.StatusText(http.StatusUnauthorized)),
		)
	}
}

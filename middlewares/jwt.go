package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"linhx1999.com/gin-blog/utils"
	"linhx1999.com/gin-blog/utils/Result"
	"net/http"
	"strings"
	"time"
)

var mySigningKey = []byte(utils.JwtKey)

type MyCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func CreatTokenByCustomClaims(username string, password string) (string, error) {
	//Create the Claims
	claims := MyCustomClaims{
		username,
		password,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "gin-blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func JwtParseUser(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
	)

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationValue := c.GetHeader("Authorization")
		//if authorizationValue == "" {
		//
		//}
		authorizationValues := strings.Split(authorizationValue, " ")

		if len(authorizationValues) != 2 && authorizationValues[0] != "Bearer" {
			c.AbortWithStatusJSON(
				http.StatusPreconditionFailed,
				Result.NewFail(
					http.StatusPreconditionFailed,
					"header Authorization has not Bearer token",
				),
			)
			return
		}

		MyCustomClaims, err := JwtParseUser(authorizationValues[1])
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusPreconditionFailed,
				Result.NewFail(
					http.StatusPreconditionFailed,
					err.Error(),
				),
			)
			return
		}

		if time.Now().Unix() > MyCustomClaims.ExpiresAt.Unix() {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				Result.NewFail(
					http.StatusUnauthorized,
					"用户未提供身份验证凭据，或者没有通过身份验证。",
				),
			)
			return
		}

		c.Set("username", MyCustomClaims.Username)
		c.Next()
	}
}

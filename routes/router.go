package routes

import (
	"github.com/gin-gonic/gin"
	v1 "linhx1999.com/gin-blog/api/v1"
	setting "linhx1999.com/gin-blog/utils"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	gin.SetMode(setting.AppMode)
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router := r.Group("api/v1")
	{
		router.POST("users", v1.PostUser)
		//router.GET("users", v1.GetUsers)
		//router.PUT("users/:id", v1.PutUser)
		//router.DELETE("users/:id", v1.DeleteUser)

		router.POST("categories", v1.PostCategory)
		router.GET("categories", v1.GetCategories)
		router.PUT("categories/:id", v1.PutCategory)
		router.DELETE("categories/:id", v1.DeleteCategory)

		router.POST("articles", v1.PostArticle)
		router.DELETE("articles/:id", v1.DeleteArticle)
		router.PUT("articles/:id", v1.PutArticle)
		router.GET("articles", v1.GetArticles)
		router.GET("articles/:id", v1.GetArticleByID)
		router.GET("categories/:id/articles", v1.GetArticlesInCategory)
	}

	// Get user value
	//r.GET("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := db[user]
	//	if ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//	}
	//})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")
		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if c.Bind(&json) == nil {
	//		db[user] = json.Value
	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})

	return r
}

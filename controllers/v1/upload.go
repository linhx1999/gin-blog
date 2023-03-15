package v1

import (
	"github.com/gin-gonic/gin"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/utils/result"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			result.New(err.Error()),
		)
		return
	}
	fileSize := fileHeader.Size
	url, err := models.UploadFile(file, fileSize)
	if err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			result.New(err.Error()),
		)
		return
	}
	c.JSON(
		http.StatusCreated,
		result.New("文件上传成功", url),
	)
}

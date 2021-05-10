package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lairdnote/liberal/serializer"
)

func Index(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Hello world",
	})
}

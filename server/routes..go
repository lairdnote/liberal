package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lairdnote/liberal/api"
)
func NewRouter() *gin.Engine{

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		// 主页的展示
		v1.GET("index", api.Index)
	}
	return r


}

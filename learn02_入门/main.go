package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	//name 必须传入已经加载的模板的名字
	r.GET("/1", func(context *gin.Context) {
		context.HTML(http.StatusOK, "1.html", nil) //渲染模板时只认名字不管路径 所以这会出现冲突
	})
	//遇到重名的情况 需要自己使用define 解决冲突
	r.GET("/2", func(context *gin.Context) {
		context.HTML(http.StatusOK, "2.html", nil)
	})
	r.Run()
}

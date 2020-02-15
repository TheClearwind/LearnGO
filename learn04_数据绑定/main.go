package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userinfo struct {
	Username string `form:"username"` //通过tag指定绑定
	Password string `form:"pwd"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		var u userinfo
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
		fmt.Println(u)
	})

	r.POST("/", func(context *gin.Context) {
		var u userinfo
		err := context.ShouldBind(&u) //post请求也可以
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
		fmt.Println(u)
	})
	r.Run(":8080")
}

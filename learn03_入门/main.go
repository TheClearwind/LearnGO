package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test(a [5]int) {
	a[0] = 5
}
func main() {
	//b := [5]int{0}
	//fmt.Println(b)
	//test(b)
	//fmt.Println(b)

	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		//query := context.Query("query") //获取get参数  post请求对应PostForm
		//query,ok:=context.GetQuery("query")
		//if !ok{
		//	return
		//}
		query := context.DefaultQuery("query", "None")
		context.JSON(http.StatusOK, gin.H{"query": query}) //返回json 可以传入一个map或者一个struct
	})

	获取路径参数
	r.GET("/path/:user/:age", func(context *gin.Context) {
		user := context.Param("user") // 获取到路径中的参数 例如访问/path/lambda/18 则user=lambda age=18
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{"name": user, "age": age})
	})
	r.Run(":8080")
}

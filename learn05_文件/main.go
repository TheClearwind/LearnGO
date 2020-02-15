package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("fl")
		if err != nil {
			c.JSON(500, gin.H{"msg": "error"})
		} else {
			dst := path.Join("./", file.Filename)
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				fmt.Printf("save failed %v", err)
			} else {
				log.Print("保存成功")
			}

		}
	})
	r.Run(":8080")
}

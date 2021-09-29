package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

func main() {
	router := gin.Default()
	router.POST("/upload/file", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "获取文件失败",
			})
		} else {
			log.Println(reflect.TypeOf(file))
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "success",
			})
		}
	})
	_ = router.Run(":8088")
}

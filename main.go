package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zanproject/go-restapi-gin/controllers/productcontroller"
	models "github.com/zanproject/go-restapi-gin/setup"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()
}

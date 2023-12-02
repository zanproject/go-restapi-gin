package productcontroller

import (
	"github.com/gin-gonic/gin"
	models "github.com/zanproject/go-restapi-gin/setup"
)

func Index(*gin.Context) {

	var products []models.Product

	models.DB.Find(&products)

}

func Show(*gin.Context) {

}

func Create(*gin.Context) {

}

func Update(*gin.Context) {

}

func Delete(*gin.Context) {

}

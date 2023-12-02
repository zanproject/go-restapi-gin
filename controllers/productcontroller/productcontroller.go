package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/zanproject/go-restapi-gin/setup"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(*gin.Context) {

}

func Create(*gin.Context) {

}

func Update(*gin.Context) {

}

func Delete(*gin.Context) {

}

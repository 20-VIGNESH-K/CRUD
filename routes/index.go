package routes

import (
	"github.com/20-VIGNESH-K/crud/services"
	"github.com/gin-gonic/gin"
)

func ProfileRoute(router *gin.Engine) {
	router.Static("/createProfile", "./frontend/createProfile/")

	router.POST("/create", services.Create)

}

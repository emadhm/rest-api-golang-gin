package routes

import (
	"emad.com/controllers"
	"github.com/gin-gonic/gin"
)



func SetUpRowters (router *gin.Engine)  {

	router.POST("/api/user/create", controllers.Create)
	
}
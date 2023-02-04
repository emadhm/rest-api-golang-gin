package routes

import (
	"emad.com/controllers"
	"github.com/gin-gonic/gin"
)



func SetUpRowters (router *gin.Engine)  {

	router.POST("/api/user/create", controllers.Create)
	router.GET("/api/user/list", controllers.List)
	router.GET("/api/user/one/:id", controllers.One)
	router.PUT("/api/user/update/:id", controllers.Update)
	router.DELETE("/api/user/delete/:id", controllers.Delete)




	
}



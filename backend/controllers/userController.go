package controllers

import (
	"emad.com/config"
	"emad.com/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Create (c *gin.Context)  {

	var body struct{

		Name  string
		Role  string
		Password  string
	}


	c.Bind(&body)

	if body.Name == "" || body.Role == "" || body.Password == "" {

			 c.JSON(500, gin.H{
				"messege": "all data required",
			} )
		
		}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

		if err != nil {
			
			 c.JSON(400, gin.H{
				"error": err,
			 })
		}
	

	post := models.Users{Name: body.Name, Role: body.Role, Passsword: string(hash)}

    result := config.DB.Create(&post)

	if result.Error !=nil {
		c.JSON(400, gin.H{
			"messege": result.Error,
		})
		 
		 
	}

	c.JSON(200, gin.H{
		"user": result,
	})


}
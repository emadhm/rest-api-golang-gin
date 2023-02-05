package auth

import (
	"emad.com/config"
	"emad.com/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)


var body struct {
		 Email string
		 Password string
	}


func Login (c *gin.Context)  {

	

	var user = new(models.Users)

		c.Bind(&body)

		if body.Email == "" || body.Password == "" {
			c.JSON(400, gin.H{
				"messege": "Invalid data",
			})
			return 
		}

		 result := config.DB.Where("email = ?", body.Email).First(&user)

		if result.Error == gorm.ErrRecordNotFound {
				c.JSON(400, gin.H{
					"messege": "user not found",
				})
				return
		    }	
		

       if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		
		    	c.JSON(400, gin.H{
					"messege": "Incorrect email or password",
				})
				return
      }


	
}



 func generateJWT(id int) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "",err
  	 }
     secretKey := os.Getenv("SECRET_KEY")


	// Create JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp":      time.Now().Add(time.Hour * 720).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
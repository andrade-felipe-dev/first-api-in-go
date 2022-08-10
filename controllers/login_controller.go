package controllers

import (
	"github.com/andrade-felipe-dev/first-api-in-go/database"
	"github.com/andrade-felipe-dev/first-api-in-go/models"
	"github.com/andrade-felipe-dev/first-api-in-go/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var login models.Login

	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível vincular a um JSON " + err.Error(),
		})

		return
	}

	var user models.User

	dbError := db.Where("email = ?", login.Email).First(&user).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível encontrar seu email" + err.Error(),
		})

		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(401, gin.H{
			"error": "Credenciais inválidas",
		})

		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

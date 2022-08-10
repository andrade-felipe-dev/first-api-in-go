package controllers

import (
	"strconv"

	"github.com/andrade-felipe-dev/first-api-in-go/database"
	"github.com/andrade-felipe-dev/first-api-in-go/models"
	"github.com/andrade-felipe-dev/first-api-in-go/services"
	"github.com/gin-gonic/gin"
)

func ShowAllUsers(c *gin.Context) {
	db := database.GetDatabase()

	var users []models.User

	err := db.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar seus usuários" + err.Error(),
		})
	}

	c.JSON(200, users)
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id precisa ser um inteiro",
		})

		return
	}

	db := database.GetDatabase()

	var user models.User

	err = db.First(&user, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Usuário não encontrado " + err.Error(),
		})
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possivel vincular a um JSON " + err.Error(),
		})

		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível cadastrar um usuário no banco de dados " + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível vincular a um JSON " + err.Error(),
		})

		return
	}

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível editar o livro " + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "O id precisa ser um valor inteiro",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.User{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar o usuário",
		})

		return
	}

	c.JSON(200, gin.H{
		"sucesso": "Usuário excluído com sucesso",
	})

}

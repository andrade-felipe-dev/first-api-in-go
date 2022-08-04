package controllers

import (
	"strconv"

	"github.com/andrade-felipe-dev/first-api-in-go/database"
	"github.com/andrade-felipe-dev/first-api-in-go/models"
	"github.com/gin-gonic/gin"
)

func ShowAllBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível listar seus livros " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID precisa ser um inteiro",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Livro não encontrado " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível vincular a um JSON " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível adicionar um livro no banco de dados " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível vincular a um JSON " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível adicionar um livro no banco de dados " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID precisa ser um inteiro",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar o livro",
		})
	}

	c.JSON(200, gin.H{
		"sucesso": "Livro excluído com sucesso",
	})
}

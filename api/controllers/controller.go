package controllers

import (
	"net/http"

	"github.com/danielsdev/microservices-go-gRPC/api/database"
	"github.com/danielsdev/microservices-go-gRPC/api/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
	//c.JSON(200, fixtures.Alunos)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	// vincula todos os dados com a struct
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	err := database.DB.Delete(&aluno, id).Error
	// database.DB.Delete(&aluno, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusNoContent, gin.H{})
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

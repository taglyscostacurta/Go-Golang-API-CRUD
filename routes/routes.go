package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taglyscostacurta/Go-Golang-API-CRUD/controllers"
	"github.com/taglyscostacurta/Go-Golang-API-CRUD/models"
)

func routeHearth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// lista todos os estudantes
func routeGetStudents(c *gin.Context) {
	c.JSON(200, gin.H{"students": models.Students})
}

// cria um estudante
func routePostStudents(c *gin.Context) {
	var student models.Student

	err := c.Bind(&student)
	if err != nil {

		c.JSON(400, controllers.NewResponseMessageError(err.Error()))
		return
	}

	student.ID = models.Students[len(models.Students)-1].ID + 1
	models.Students = append(models.Students, student)

	c.JSON(201, student)
}

// lista um estudante
func routePutStudents(c *gin.Context) {
	var studentPayload models.Student
	var studentLocal models.Student
	var newStudents []models.Student

	err := c.BindJSON(&studentPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messge_error": "Não foi possível obter o payload",
		})
		return
	}

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messge_error": "Não foi possível obter o id",
		})
		return
	}

	for _, studentElement := range models.Students {
		if studentElement.ID == id {
			studentLocal = studentElement

		}
	}
	// se o id for igual a zero talvez o estudande procurado nao exista mais ou nao foi encontrado
	if studentLocal.ID == 0 {
		c.JSON(404, gin.H{
			"messge_error": "Não foi possível encontrar o estudante",
		})
		return

	}

	studentLocal.FullName = studentPayload.FullName
	studentLocal.Age = studentPayload.Age

	for _, studentElement := range models.Students {
		if id == studentElement.ID {
			newStudents = append(newStudents, studentLocal)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	models.Students = newStudents

	c.JSON(http.StatusOK, studentLocal)
}

// delet student
func routeDeleteStudents(c *gin.Context) {
	var newStudents []models.Student

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messge_error": "Não foi possível obter o id",
		})
		return
	}

	for _, studentElement := range models.Students {
		if id != studentElement.ID {
			newStudents = append(newStudents, studentElement)
		}
	}

	models.Students = newStudents

	c.JSON(http.StatusOK, controllers.NewResponseMessage("Estudante deletado com sucesso"))
}

func GetRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHearth)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.POST("/", routePostStudents)
	groupStudents.PUT("/:id", routePutStudents)
	groupStudents.DELETE("/:id", routeDeleteStudents)

	return c
}

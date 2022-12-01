package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taglyscostacurta/Go-Golang-API-CRUD/models"
)

func routeHearth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// lista todos os estudantes
func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Students)
}

// cria um estudante
func routePostStudents(c *gin.Context) {
	var student models.Student

	err := c.Bind(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messge_error": "Não foi possível adicionar um estudante",
		})
		return
	}
	student.ID = len(models.Students) + 1
	models.Students = append(models.Students, student)

	// student.ID = models.Students[len(models.Students)-1].ID
	// models.Students = append(models.Students, student)

	c.JSON(http.StatusCreated, student)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"messge_error": "Não foi possível encontrar o estudante",
		})
		return

	}

	studentLocal.FullName = studentPayload.FullName
	studentLocal.Age = studentLocal.Age

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

func GetRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHearth)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.POST("/", routePostStudents)
	groupStudents.PUT("/:id", routePutStudents)

	return c
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routeHearth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func routePostStudents(c *gin.Context) {
	var student Student

	err := c.Bind(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messge_error": "Não foi possível obter o payload",
		})
		return
	}
	student.ID = len(Students) + 1
	Students = append(Students, student)

	c.JSON(http.StatusCreated, student)
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHearth)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.POST("/", routePostStudents)

	return c
}

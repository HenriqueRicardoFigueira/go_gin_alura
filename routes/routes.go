package routes

import (
	"go_gin_alura/controllers"
	"go_gin_alura/database"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	database.ConnectDb()
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudent)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/search", controllers.SearchStudentByName)
	r.Run()
}

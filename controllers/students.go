package controllers

import (
	"go_gin_alura/database"
	"go_gin_alura/models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(200, student)
}

func SearchStudentByName(c *gin.Context) {
	var student []models.Student
	name := c.Query("name")

	database.DB.Where("name LIKE ?", "%"+name+"%").Find(&student)
	if len(student) == 0 {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(200, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&student).Where("id = ?", id).Updates(&student)
	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.Delete(&student, id)
	c.JSON(200, gin.H{"message": "Student deleted"})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(200, student)
}

func Index(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.HTML(200, "index.html", gin.H{
		"students": students,
	})
}

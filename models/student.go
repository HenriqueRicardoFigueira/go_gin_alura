package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name"`
	CPF  string `json:"CPF"`
	RG   string `json:"RG"`
}

func GetAllStudents() []Student {
	var students []Student

	return students
}

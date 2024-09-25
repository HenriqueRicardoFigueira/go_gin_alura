package main

import (
	"go_gin_alura/controllers"
	"go_gin_alura/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func routesSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestCheckStatusCodeToGetAllStudents(t *testing.T) {
	r := routesSetup()
	database.ConnectDb()
	r.GET("/students", controllers.GetStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

package main

import (
	"webserviceapi/config"
	"webserviceapi/controllers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "webserviceapi/docs"

	"github.com/gin-gonic/gin"
)

// @title Students API
// @version 1.0
// @description This is a basic API Students using Gin and Gorm.

// @contact.name API Support
// @contact.email davidwinalda@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/students", controllers.GetStudents)
		v1.GET("/student/:id", controllers.GetStudentById)
		v1.POST("/student", controllers.CreateStudent)
		v1.PUT("/student/:id", controllers.UpdateStudentById)
		v1.DELETE("/student/:id", controllers.DeleteStudentById)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

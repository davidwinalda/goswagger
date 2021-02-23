package controllers

import (
	"net/http"
	"webserviceapi/config"
	"webserviceapi/models"

	"github.com/gin-gonic/gin"
)

// GetStudents godoc
// @Summary Show a list of students
// @Tags student
// @Accept  json
// @Produce  json
// @success 200 {object} models.Student
// @Router /students [get]
func GetStudents(c *gin.Context) {
	var students []models.Student

	if err := config.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	c.JSON(http.StatusOK, students)
}

// GetStudentById godoc
// @Summary Show student by Id
// @Tags student
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Student ID"
// @success 200 {object} models.Student
// @Router /student/{id} [get]
func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// CreateStudent godoc
// @Summary Create student data
// @Description Add by JSON Student
// @Tags student
// @Accept  json
// @Produce  json
// @Param student body models.Student true "Add Student"
// @success 200 {object} models.Student
// @Router /student [post]
func CreateStudent(c *gin.Context) {
	var student models.Student

	c.ShouldBindJSON(&student)

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// UpdateStudentById godoc
// @Summary Update student data by Id
// @Description Update by JSON Student
// @Tags student
// @Accept  json
// @Produce  json
// @Param id path int true "Student Id"
// @Param student body models.Student true "Update Student"
// @success 200 {object} models.Student
// @Router /student/{id} [put]
func UpdateStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&student)

	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudentById godoc
// @Summary Delete student by Id
// @Tags student
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Student ID"
// @success 204 {object} models.Student
// @Router /student/{id} [delete]
func DeleteStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	if err := config.DB.Where("id = ?", id).Delete(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
}

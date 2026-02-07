package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		// ปรับปรุง error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		// เพิ่ม error message ที่ชัดเจน
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		// ปรับปรุง error message
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		// แยก validation error กับ server error
		if err.Error() == "id must not be empty" ||
			err.Error() == "name must not be empty" ||
			err.Error() == "gpa must be between 0.00 and 4.00" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// Update Student (PUT)
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		// เพิ่ม error message ที่ชัดเจน
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// เรียก Service เพื่ออัปเดต
	updatedStudent, err := h.Service.UpdateStudent(id, student)
	if err != nil {
		// แยก validation error กับ not found error
		if err.Error() == "id must not be empty" ||
			err.Error() == "name must not be empty" ||
			err.Error() == "gpa must be between 0.00 and 4.00" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// ถ้าหาไม่เจอ ส่ง 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

// Delete Student (DELETE)
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteStudent(id)
	if err != nil {
		// เพิ่ม error message ที่ชัดเจน
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// ส่ง 204 No Content เมื่อลบสำเร็จ
	c.Status(http.StatusNoContent)
}

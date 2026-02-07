package services

import (
	"errors"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	// เพิ่ม validation ก่อนสร้างข้อมูล
	if err := s.ValidateStudent(student); err != nil {
		return err
	}
	return s.Repo.Create(student)
}

// ตรวจสอบความถูกต้องของข้อมูลนักเรียน
func (s *StudentService) ValidateStudent(student models.Student) error {
	// ตรวจสอบว่า ID ไม่ว่าง
	if student.Id == "" {
		return errors.New("id must not be empty")
	}

	// ตรวจสอบว่า Name ไม่ว่าง
	if student.Name == "" {
		return errors.New("name must not be empty")
	}

	// ตรวจสอบว่า GPA อยู่ในช่วง 0.00 - 4.00 ไหม
	if student.GPA < 0.00 || student.GPA > 4.00 {
		return errors.New("gpa must be between 0.00 and 4.00")
	}

	return nil
}

// อัปเดตข้อมูลนักเรียน (มีการ validate)
func (s *StudentService) UpdateStudent(id string, student models.Student) (*models.Student, error) {
	// ตรวจสอบข้อมูลก่อน
	if err := s.ValidateStudent(student); err != nil {
		return nil, err
	}

	// เรียก Repository เพื่ออัปเดต
	return s.Repo.Update(id, student)
}

// Delete Student
func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.Delete(id)
}

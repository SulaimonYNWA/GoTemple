// services/school_service.go
package services

import (
	entity "github.com/SulaimonYNWA/GoTemple/entity/schools"
	courses_entity "github.com/SulaimonYNWA/GoTemple/entity/courses"
	instructors_entity "github.com/SulaimonYNWA/GoTemple/entity/instructors"
	repo "github.com/SulaimonYNWA/GoTemple/repo"
)

type SchoolService interface {
	GetAll() ([]entity.School, error)
	GetCoursesBySchool(schoolID int) ([]courses_entity.Course, error)
	GetAllInstructors() ([]instructors_entity.Instructor, error)
}

type schoolService struct {
	SchoolRepo repo.SchoolRepo
	CourseRepo repo.CourseRepo
	InstructorRepo repo.InstructorRepo
}

func NewSchoolService(schoolRepo repo.SchoolRepo, courseRepo repo.CourseRepo, instructorRepo repo.InstructorRepo) SchoolService {
	return &schoolService{SchoolRepo: schoolRepo, CourseRepo: courseRepo, InstructorRepo: instructorRepo}
}

func (s *schoolService) GetAll() ([]entity.School, error) {
	return s.SchoolRepo.GetAll()
}

func (s *schoolService) GetCoursesBySchool(schoolID int) ([]courses_entity.Course, error) {
	return s.CourseRepo.GetBySchoolID(schoolID)
}

func (s *schoolService) GetAllInstructors() ([]instructors_entity.Instructor, error) {
	return s.InstructorRepo.GetAll()
}
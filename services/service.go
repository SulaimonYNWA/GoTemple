// services/school_service.go
package services

import (
	entity "github.com/SulaimonYNWA/GoTemple/entity/schools"
	courses_entity "github.com/SulaimonYNWA/GoTemple/entity/courses"
	repo "github.com/SulaimonYNWA/GoTemple/repo"
)

type SchoolService interface {
	GetAll() ([]entity.School, error)
	GetCoursesBySchool(schoolID int) ([]courses_entity.Course, error)
}

type schoolService struct {
	SchoolRepo repo.SchoolRepo
	CourseRepo repo.CourseRepo
}

func NewSchoolService(schoolRepo repo.SchoolRepo, courseRepo repo.CourseRepo) SchoolService {
	return &schoolService{SchoolRepo: schoolRepo, CourseRepo: courseRepo}
}

func (s *schoolService) GetAll() ([]entity.School, error) {
	return s.SchoolRepo.GetAll()
}

func (s *schoolService) GetCoursesBySchool(schoolID int) ([]courses_entity.Course, error) {
	return s.CourseRepo.GetBySchoolID(schoolID)
}
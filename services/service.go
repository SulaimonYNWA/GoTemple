// services/school_service.go
package services

import (
	entity "github.com/SulaimonYNWA/GoTemple/entity/schools"
	repo "github.com/SulaimonYNWA/GoTemple/repo"
)

type SchoolService interface {
	GetAll() ([]entity.School, error)
}

type schoolService struct {
	Repo repo.SchoolRepo
}

func NewSchoolService(repo repo.SchoolRepo) SchoolService {
	return &schoolService{Repo: repo}
}

func (s *schoolService) GetAll() ([]entity.School, error) {
	return s.Repo.GetAll()
}

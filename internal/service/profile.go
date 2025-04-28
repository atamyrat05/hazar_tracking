package service

import (
	"hazar_tracking/internal/model"
	"hazar_tracking/internal/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *ProfileService) GetById(userId int) (model.UpdateProfileInput, error) {
	return s.repo.GetById(userId)
}

func (s *ProfileService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *ProfileService) Update(userId int, input model.UpdateProfileInput) error {
	return s.repo.Update(userId, input)
}

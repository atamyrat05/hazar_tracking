package service

import (
	"hazar_tracking/internal/model"
	"hazar_tracking/internal/repository"
)

type AnnouncementService struct {
	repo repository.Announcement
}

func NewAnnouncementService(repo repository.Announcement) *AnnouncementService {
	return &AnnouncementService{repo: repo}
}

func (s *AnnouncementService) Create(order model.AnnouncementInput, userId int) (int, error) {
	return s.repo.Create(order, userId)
}

func (s *AnnouncementService) GetAll() ([]model.AnnouncementGet, error) {
	return s.repo.GetAll()
}

func (s *AnnouncementService) GetById(announcementId int) (model.AnnouncementGetById, error) {
	return s.repo.GetById(announcementId)
}

package service

import (
	"errors"
	"hazar_tracking/internal/model"
	"hazar_tracking/internal/repository"
)

type EmailService struct {
	repo repository.Email
}

func NewEmailsService(repo repository.Email) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) Validate(email string) (int, error) {
	return s.repo.Validate(email)
}

func (s *EmailService) UpdateForgotCode(userId int, code string) error {
	return s.repo.UpdateForgotCode(userId, code)
}

func (s *EmailService) TestCode(code string, userId int) error {
	return s.repo.TestCode(code, userId)
}

func (s *EmailService) UpdateUsersPassword(input model.UpdatePassword, userId int) error {
	if input.Password_confirm != input.Password {
		return errors.New("password not confirmed")
	}
	return s.repo.UpdateUsersPassword(input, userId)
}

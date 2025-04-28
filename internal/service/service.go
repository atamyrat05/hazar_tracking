package service

import (
	"hazar_tracking/internal/model"
	"hazar_tracking/internal/repository"
)

type Profile interface {
	GetById(userId int) (model.UpdateProfileInput, error)
	GetAll() ([]model.User, error)
	Delete(userId int) error
	Update(userId int, input model.UpdateProfileInput) error
}

type Authorization interface {
	Create(user model.User) (int, error)
	GenerateToken(input model.SignIn) (string, error)
	ParseToken(token string) (int, error)
}

type Orders interface {
	Create(order model.OrdersInput, userId int) (int, error)
	GetAll(userId int) ([]model.OrdersGet, error)
	GetById(userId, orderId int) (model.OrdersGet, []model.OrderTrackingSteps, error)
	Search(userId int, input string) (model.OrdersGet, error)
	GetAllPoints() ([]model.Points, error)
	Update(orderId int, input model.UpdateOrderInput) error
	CreatePoints(input model.OrderTrackingStepsInput) (int, error)
	UpdatePoints(input model.UpdateOrderTrackingStepsInput) error
}

type Email interface {
	Validate(email string) (int, error)
	UpdateForgotCode(userId int, code string) error
	TestCode(code string, userId int) error
	UpdateUsersPassword(input model.UpdatePassword, userId int) error
}

type Announcement interface {
	Create(input model.AnnouncementInput, userId int) (int, error)
	GetAll() ([]model.AnnouncementGet, error)
	GetById(announcementId int) (model.AnnouncementGetById, error)
}

type Service struct {
	Authorization
	Profile
	Orders
	Email
	Announcement
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Profile:       NewProfileService(repo.Profile),
		Orders:        NewOrdersService(repo.Orders),
		Email:         NewEmailsService(repo.Email),
		Announcement:  NewAnnouncementService(repo.Announcement),
	}
}

package repository

import (
	"hazar_tracking/internal/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user model.User) (int, error)
	Get(email, password string) (model.User, error)
}

type Profile interface {
	GetById(userId int) (model.UpdateProfileInput, error)
	GetAll() ([]model.User, error)
	Delete(userId int) error
	Update(userId int, input model.UpdateProfileInput) error
}

type Orders interface {
	Create(order model.OrdersInput, userId int, qrcode_url string) (int, error)
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

type Repository struct {
	Authorization
	Profile
	Orders
	Email
	Announcement
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Profile:       NewProfileRepository(db),
		Orders:        NewOrdersRepository(db),
		Email:         NewEmailsRepository(db),
		Announcement:  NewAnnouncementRepository(db),
	}
}

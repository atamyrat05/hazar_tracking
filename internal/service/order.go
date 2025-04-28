package service

import (
	"errors"
	"hazar_tracking/internal/model"
	"hazar_tracking/internal/repository"
	"hazar_tracking/utilits"
)

type OrderService struct {
	repo repository.Orders
}

func NewOrdersService(repo repository.Orders) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(order model.OrdersInput, userId int) (int, error) {
	qrcode_url, err := utilits.GenerateQRCode()
	if err != nil {
		return 0, errors.New("can not generate qrcode")
	}
	return s.repo.Create(order, userId, qrcode_url)
}

func (s *OrderService) GetAll(userId int) ([]model.OrdersGet, error) {
	return s.repo.GetAll(userId)
}

func (s *OrderService) GetById(userId, orderId int) (model.OrdersGet, []model.OrderTrackingSteps, error) {
	return s.repo.GetById(userId, orderId)
}
func (s *OrderService) Search(userId int, input string) (model.OrdersGet, error) {
	return s.repo.Search(userId, input)
}

func (s *OrderService) GetAllPoints() ([]model.Points, error) {
	return s.repo.GetAllPoints()
}

func (s OrderService) Update(orderId int, input model.UpdateOrderInput) error {
	return s.repo.Update(orderId, input)
}

func (s OrderService) CreatePoints(input model.OrderTrackingStepsInput) (int, error) {
	return s.repo.CreatePoints(input)
}

func (s *OrderService) UpdatePoints(input model.UpdateOrderTrackingStepsInput) error {
	return s.repo.UpdatePoints(input)
}

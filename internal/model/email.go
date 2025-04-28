package model

type Email struct {
	Email string `json:"email" db:"email" binding:"required"`
}

type Code struct {
	Code string `json:"code" db:"forgot_pas_code" binding:"required"`
}

type UpdatePassword struct {
	Password         string `json:"password" db:"password" binding:"required"`
	Password_confirm string `json:"password_confirm" db:"password_confirm" binding:"required"`
}

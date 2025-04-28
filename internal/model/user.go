package model

type User struct {
	Id               int     `json:"-" db:"id"`
	First_name       string  `json:"first_name" db:"first_name" binding:"required"`
	Last_name        string  `json:"last_name" db:"last_name" binding:"required"`
	Email            string  `json:"email" db:"email" binding:"required"`
	PhoneNumber      string  `json:"phone_number" db:"phone_number" binding:"required"`
	Password         string  `json:"password" db:"password" binding:"required"`
	Password_confirm string  `json:"password_confirm" db:"password_confirm" binding:"required"`
	Code             *string `json:"code" db:"forgot_pas_code"`
	Image_Url        *string `json:"image_url" db:"image_urls"`
}

type UpdateProfileInput struct {
	Id          int     `json:"id" db:"id"`
	First_name  string  `json:"first_name" db:"first_name"`
	Last_name   string  `json:"last_name" db:"last_name"`
	Email       string  `json:"email" db:"email"`
	PhoneNumber string  `json:"phone_number" db:"phone_number"`
	Image_Url   *string `json:"image_url" db:"image_urls"`
}

type SignIn struct {
	Email    string `json:"email" db:"email"  binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

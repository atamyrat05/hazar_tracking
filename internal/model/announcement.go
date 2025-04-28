package model

type AnnouncementInput struct {
	Id          int    `json:"id" db:"id"`
	Category    string `json:"category" db:"category" binding:"required"`
	Time        string `json:"time" db:"time" binding:"required"`
	From_where  int    `json:"from_where" db:"from_where" binding:"required"`
	Where_to    int    `json:"where_to" db:"where_to" binding:"required"`
	Text        string `json:"text" db:"text" binding:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number" binding:"required"`
	Name        string `json:"name" db:"name" binding:"required"`
}

type AnnouncementGet struct {
	Id         int    `json:"id" db:"id"`
	Category   string `json:"category" db:"category"`
	Time       string `json:"time" db:"time"`
	From_where string `json:"from_where" db:"from_where"`
	Where_to   string `json:"where_to" db:"where_to"`
	Text       string `json:"text" db:"text"`
}

type AnnouncementGetById struct {
	Id          int    `json:"id" db:"id"`
	Category    string `json:"category" db:"category"`
	Time        string `json:"time" db:"time"`
	From_where  string `json:"from_where" db:"from_where"`
	Where_to    string `json:"where_to" db:"where_to"`
	Text        string `json:"text" db:"text"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Name        string `json:"name" db:"name"`
}

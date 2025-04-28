package model

type OrdersInput struct {
	Id              int    `json:"-" db:"id"`
	Senders_name    string `json:"senders_name" db:"senders_name" binding:"required"`
	Buyers_name     string `json:"buyers_name" db:"buyers_name" binding:"required"`
	From_where      int    `json:"from_where" db:"from_where" binding:"required"`
	Where_to        int    `json:"where_to" db:"where_to" binding:"required"`
	Type_of_service int    `json:"type_of_service" db:"type_of_service" binding:"required"`
	Weight          string `json:"weight" db:"weight" binding:"required"`
	Status          int    `json:"status" db:"status"`
	UsersId         int    `json:"-" db:"users_id"`
	Seria_id        string `json:"seria_id" db:"seria_id"`
}

type OrdersGet struct {
	Id                  int     `json:"id" db:"id"`
	Senders_name        string  `json:"senders_name" db:"senders_name" binding:"required"`
	Buyers_name         string  `json:"buyers_name" db:"buyers_name" binding:"required"`
	From_where          string  `json:"from_where" db:"from_where" binding:"required"`
	Where_to            string  `json:"where_to" db:"where_to" binding:"required"`
	Type_of_service     int     `json:"type_of_service" db:"type_of_service" binding:"required"`
	Weight              string  `json:"weight" db:"weight" binding:"required"`
	Status              int     `json:"status" db:"status"`
	UsersId             int     `json:"-" db:"users_id"`
	Seria_id            string  `json:"seria_id" db:"seria_id"`
	Started_time        *string `json:"started_time" db:"started_time"`
	Finished_time       *string `json:"finished_time" db:"finished_time"`
	Current_location    *string `json:"current_location" db:"name"`
	Total_steps         int     `json:"total_steps" db:"total_steps"`
	Current_step_number int     `json:"current_step_number" db:"current_step_number"`
}

type UpdateOrderInput struct {
	Id              int     `json:"id" db:"id"`
	Senders_name    string  `json:"senders_name" db:"senders_name"`
	Buyers_name     string  `json:"buyers_name" db:"buyers_name"`
	From_where      string  `json:"from_where" db:"from_where"`
	Where_to        string  `json:"where_to" db:"where_to"`
	Type_of_service int     `json:"type_of_service" db:"type_of_service"`
	Weight          string  `json:"weight" db:"weight"`
	Status          int     `json:"status" db:"status"`
	Seria_id        string  `json:"seria_id" db:"seria_id"`
	Started_time    *string `json:"started_time" db:"started_time"`
	Finished_time   *string `json:"finished_time" db:"finished_time"`
}

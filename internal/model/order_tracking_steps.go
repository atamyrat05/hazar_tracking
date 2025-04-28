package model

type OrderTrackingSteps struct {
	Location string  `json:"location" db:"name"`
	Date     *string `json:"date" db:"step_date"`
}

type OrderTrackingStepsInput struct {
	OrderId  int `json:"order_id" db:"order_id" binding:"required"`
	Location int `json:"location" db:"location" binding:"required"`
}

type UpdateOrderTrackingStepsInput struct {
	OrderId  int    `json:"order_id" db:"order_id"`
	Location int    `json:"location" db:"location" binding:"required"`
	StepDate string `json:"step_date" db:"step_date"`
}

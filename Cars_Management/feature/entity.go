package feature

import (
	"time"
)

type Cars struct {
	Car_id     uint32 `gorm:"primarykey"`
	Car_Name   string
	Day_rate   float32
	Month_rate float32
	Image      string
}

type Orders struct {
	Order_id         uint32 `gorm:"primarykey"`
	Car_id           uint32
	Order_date       *time.Time
	Pickup_date      *time.Time
	Dropoff_date     *time.Time
	Pickup_location  string
	Dropoff_location string
}

type ReqCarsDetail struct {
	Car_id uint32 `json:"car_id"`
}

type ResponseCarsDetail struct {
	Status string
	Data   Cars
}

type ReqInputCar struct {
	Car_Name   string  `json:"car_name"`
	Day_rate   float32 `json:"day_rate"`
	Month_rate float32 `json:"month_rate"`
	Image      string  `json:"image"`
}

type ReqOrders struct {
	Order_id         uint32 `json:"order_id"`
	Car_id           uint32 `json:"car_id"`
	Order_date       string `json:"order_date"`
	Pickup_date      string `json:"pickup_date"`
	Dropoff_date     string `json:"dropoff_date"`
	Pickup_location  string `json:"pickup_location"`
	Dropoff_location string `json:"dropoff_location"`
}

type DataInterface interface {
	GetCarsDetail(id uint32) (ResponseCarsDetail, error)
	InputCars(req Cars) error
	MakeOrder(req Orders) error
	UpdateOrder(req Orders) error
	GetOrderrsDetail(id uint32) (Orders, error)
}

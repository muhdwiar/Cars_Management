package data

import (
	"fmt"
	"project/cars-shop/feature"

	"gorm.io/gorm"
)

type ApiServer struct {
	db *gorm.DB
}

func New(db *gorm.DB) feature.DataInterface {
	return &ApiServer{
		db: db,
	}
}

func (d *ApiServer) InputCars(req feature.Cars) error {

	carData := feature.Cars{
		Car_Name:   req.Car_Name,
		Day_rate:   req.Day_rate,
		Month_rate: req.Month_rate,
		Image:      req.Image,
	}
	fmt.Println(carData)
	if err := d.db.Create(&carData).Error; err != nil {
		return err
	}

	return nil
}

func (d *ApiServer) GetCarsDetail(id uint32) (feature.ResponseCarsDetail, error) {

	resp := feature.ResponseCarsDetail{
		Status: "Sukses",
	}

	dataDetail := &feature.Cars{}
	if err := d.db.Where("car_id= ?", id).First(dataDetail).Error; err != nil {
		return feature.ResponseCarsDetail{}, err
	}

	resp.Data.Car_Name = dataDetail.Car_Name
	resp.Data.Car_id = dataDetail.Car_id
	resp.Data.Day_rate = dataDetail.Day_rate
	resp.Data.Image = dataDetail.Image
	resp.Data.Month_rate = dataDetail.Month_rate

	return resp, nil
}

func (d *ApiServer) MakeOrder(req feature.Orders) error {

	orderData := feature.Orders{
		Car_id:          req.Car_id,
		Order_date:      req.Order_date,
		Pickup_location: req.Pickup_location,
	}
	fmt.Println(orderData)
	if err := d.db.Create(&orderData).Error; err != nil {
		return err
	}

	return nil
}

func (d *ApiServer) UpdateOrder(req feature.Orders) error {

	orderData := feature.Orders{
		Order_id:         req.Order_id,
		Car_id:           req.Car_id,
		Order_date:       req.Order_date,
		Pickup_location:  req.Pickup_location,
		Dropoff_date:     req.Dropoff_date,
		Dropoff_location: req.Dropoff_location,
	}
	fmt.Println(orderData)
	if err := d.db.Save(&orderData).Error; err != nil {
		return err
	}

	return nil
}

func (d *ApiServer) GetOrderrsDetail(id uint32) (feature.Orders, error) {

	dataDetail := &feature.Orders{}
	if err := d.db.Where("order_id= ?", id).First(dataDetail).Error; err != nil {
		return feature.Orders{}, err
	}

	return *dataDetail, nil
}

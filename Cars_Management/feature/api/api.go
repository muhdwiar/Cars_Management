package api

import (
	"fmt"
	"net/http"
	"project/cars-shop/feature"

	"github.com/labstack/echo"
)

type ApiServer struct {
	dataInterface feature.DataInterface
}

func New(e *echo.Echo, d feature.DataInterface) {
	handler := &ApiServer{
		dataInterface: d,
	}

	e.POST("/cars/input", handler.InputCars)
	e.GET("/cars/detail", handler.GetCarsDetail)
	e.POST("/order/input", handler.InputOrder)
	e.POST("/order/update", handler.EndOrder)
}

func (d *ApiServer) InputCars(ctx echo.Context) error {
	req := &feature.ReqInputCar{}
	ctx.Bind(req)

	carData := feature.Cars{
		Car_Name:   req.Car_Name,
		Day_rate:   req.Day_rate,
		Month_rate: req.Month_rate,
		Image:      req.Image,
	}
	fmt.Println(carData)
	if err := d.dataInterface.InputCars(carData); err != nil {
		return ctx.JSON(http.StatusInternalServerError, Fail_Resp("Error Input Data Car"))
	}

	return ctx.JSON(http.StatusOK, Succes_Resp())
}

func (d *ApiServer) GetCarsDetail(ctx echo.Context) error {
	req := &feature.ReqCarsDetail{}
	ctx.Bind(req)

	resp := feature.ResponseCarsDetail{
		Status: "Sukses",
	}

	dataDetail, err := d.dataInterface.GetCarsDetail(req.Car_id)

	if err != nil {
		return ctx.JSON(http.StatusNonAuthoritativeInfo, Fail_Resp("Error Get Data From DB"))
	}
	resp = dataDetail

	return ctx.JSON(http.StatusOK, resp)
}

func (d *ApiServer) InputOrder(ctx echo.Context) error {
	req := &feature.ReqOrders{}
	ctx.Bind(req)

	dataDetail, err := d.dataInterface.GetCarsDetail(req.Car_id)

	if err != nil {
		return ctx.JSON(http.StatusNonAuthoritativeInfo, Fail_Resp("Error Get Data From DB"))
	}

	if (dataDetail == feature.ResponseCarsDetail{}) {
		return ctx.JSON(http.StatusInternalServerError, Fail_Resp("Car Data is not found"))
	}

	order_date, err := DateConvert(req.Order_date)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, Fail_Resp("Error Convert Order Date"))
	}

	orderData := feature.Orders{
		Car_id:          req.Car_id,
		Order_date:      &order_date,
		Pickup_location: req.Pickup_location,
	}
	fmt.Println(orderData)
	if err := d.dataInterface.MakeOrder(orderData); err != nil {
		return ctx.JSON(http.StatusInternalServerError, Fail_Resp("Error Input Data Car"))
	}

	return ctx.JSON(http.StatusOK, Succes_Resp())
}

func (d *ApiServer) EndOrder(ctx echo.Context) error {
	req := &feature.ReqOrders{}
	ctx.Bind(req)

	dataDetail, err := d.dataInterface.GetOrderrsDetail(req.Order_id)

	if err != nil {
		return ctx.JSON(http.StatusNonAuthoritativeInfo, Fail_Resp("Error Get Data From DB"))
	}

	if (dataDetail == feature.Orders{}) {
		return ctx.JSON(http.StatusInternalServerError, Fail_Resp("Order Data is not found"))
	}

	dropoff_date, err := DateConvert(req.Dropoff_date)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, Fail_Resp("Error Convert Order Date"))
	}

	orderData := dataDetail
	orderData.Dropoff_date = &dropoff_date
	orderData.Dropoff_location = req.Dropoff_location
	fmt.Println(orderData)
	if err := d.dataInterface.UpdateOrder(orderData); err != nil {
		return ctx.JSON(http.StatusInternalServerError, Fail_Resp("Error Input Data Car"))
	}

	return ctx.JSON(http.StatusOK, Succes_Resp())
}

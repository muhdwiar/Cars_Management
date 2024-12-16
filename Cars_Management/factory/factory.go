package factory

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"

	// "project/cars-shop/feature"
	"project/cars-shop/feature/api"
	"project/cars-shop/feature/data"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	dataInterface := data.New(db)
	api.New(e, dataInterface)
}

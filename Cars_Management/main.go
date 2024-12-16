package main

import (
	"fmt"
	"project/cars-shop/factory"
	"project/cars-shop/feature/data"
	"project/cars-shop/migration"

	"github.com/labstack/echo"
)

func main() {
	db := data.InitDB()
	migration.InitMigrate(db)
	e := echo.New()

	factory.InitFactory(e, db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8000)))
}

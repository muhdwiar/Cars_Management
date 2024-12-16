package migration

import (
	"project/cars-shop/feature"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&feature.Cars{})
	db.AutoMigrate(&feature.Orders{})

}

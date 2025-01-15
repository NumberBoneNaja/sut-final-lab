package entity

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name string
	Price float64 `valid:"range(1|100)~Price must be between 1 and 1000 ,required~Price is required"`
	SKU string  

	
}
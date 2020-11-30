package main

import "github.com/jinzhu/gorm"

type Goods struct {
	gorm.Model
	Name   string
	Price  float64
	Number uint32
}


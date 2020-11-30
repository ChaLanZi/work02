package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

type GoodsService struct {
	ID     uint32  `from:"id" json:"id" binding:"min=0"`
	Name   string  `from:"name" json:"name" binding:"min=0,max=100"`
	Price  float64 `from:"price" json:"price" binding:"min=0"`
	Number uint32   `from:"number" json:"number" binding:"min=0"`
}

func (s *GoodsService) Create() error {
	goods := Goods{
		Name:   s.Name,
		Price:  s.Price,
		Number: s.Number,
	}
	err := DB.Create(&goods).Error
	return errors.Wrap(err, "定位错误位置")
}

func (s *GoodsService) FindOne() (*Goods, error) {
	var goods *Goods

	err := DB.Where("id=?", s.ID).First(&goods).Error
	// sql.ErrNoRows  不需要定位，直接返回
	if err == sql.ErrNoRows {
		return nil, err
	}
	return goods, errors.Wrap(err, "定位返回nil的其他错误")

}

func (s *GoodsService) FindAll() ([]Goods, error) {
	var goods []Goods
	err := DB.Find(&goods).Error
	// sql.ErrNoRows  不需要定位，直接返回
	if err == sql.ErrNoRows {
		return nil, err
	}
	return goods, errors.Wrap(err, "定位返回nil的其他错误")

}
func (s *GoodsService) Update() error {
	goods := Goods{
		Name:   s.Name,
		Price:  s.Price,
		Number: s.Number,
	}
	return errors.Wrap(DB.Save(&goods).Error, "定位错误")
}

func (s *GoodsService) Delete() error {
	goods := Goods{
		Name: s.Name,
	}
	return errors.Wrap(DB.Delete(&goods).Error, "定位错误")
}

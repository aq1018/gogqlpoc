package db

import (
	"github.com/PeerStreet/aqgqlpoc/graph/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init(dbstr string) (err error) {
	DB, err = gorm.Open("postgres", dbstr)
	return err
}

func AutoMigrate() {
	DB.AutoMigrate(&model.Loan{}, &model.Property{})
}

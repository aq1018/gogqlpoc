package db

import (
	"log"
	"os"

	"github.com/aq1018/gogqlpoc/graph/model"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var Redis *redis.Pool

func init() {
	Redis = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(os.Getenv("REDIS_URL"))
		},
	}

	var err error
	DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&model.Loan{}, &model.Property{})
}

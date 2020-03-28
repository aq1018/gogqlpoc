package db

import (
	"log"
	"os"

	"github.com/aq1018/gogqlpoc/graph/model"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const REDIS_WORKER_NAMESPACE = "gogqlpoc"

func NewDB() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}
	return db
}

func NewRedis() *redis.Pool {
	return &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(os.Getenv("REDIS_URL"))
		},
	}
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Loan{}, &model.Property{})
}

package worker

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aq1018/gogqlpoc/db"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

type Worker struct {
	pool *work.WorkerPool
}

type Context struct {
	DB    *gorm.DB
	Redis *redis.Pool
}

func New() *Worker {
	context := Context{DB: db.DB, Redis: db.Redis}
	pool := work.NewWorkerPool(context, 10, "gogqlpoc", db.Redis)
	return &Worker{pool}
}

func (worker *Worker) Run() {
	worker.registerMiddlewares()
	worker.registerJobs()

	worker.pool.Start()
	fmt.Println("Worker Started.")

	waitSigKill()
	worker.pool.Stop()
	fmt.Println("Worker Stopped.")
}

func waitSigKill() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
}

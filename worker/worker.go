package worker

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/aq1018/gogqlpoc/db"
	"github.com/aq1018/gogqlpoc/operation"
	"github.com/gocraft/work"
)

type Worker struct {
	pool *work.WorkerPool
}

type Context struct {
	op *operation.Operation
}

func NewWorker() *Worker {
	context := Context{op: operation.NewOperation()}
	pool := work.NewWorkerPool(context, 10, db.REDIS_WORKER_NAMESPACE, context.op.Redis)
	return &Worker{pool: pool}
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

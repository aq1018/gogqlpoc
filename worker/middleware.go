package worker

import (
	"fmt"

	"github.com/gocraft/work"
)

func (worker *Worker) registerMiddlewares() {
	worker.pool.Middleware(Log)
}

func Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

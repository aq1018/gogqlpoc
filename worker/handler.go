package worker

import (
	"fmt"
	"time"

	"github.com/gocraft/work"
)

func (worker *Worker) registerJobs() {
	worker.pool.Job("send_email", (*Context).SendEmail)
}

func (context *Context) SendEmail(job *work.Job) error {
	loanID := job.ArgInt64("loanID")

	fmt.Println("Start sending email for LoanID: ", loanID)

	// fake sending email
	time.Sleep(200 * time.Millisecond)

	fmt.Println("Email sent!")

	return nil
}

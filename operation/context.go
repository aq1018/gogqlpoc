package operation

import (
	"context"
	"net/http"

	"github.com/aq1018/gogqlpoc/db"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

var operationCtxKey = contextKey{"operation"}

type contextKey struct {
	name string
}

type Operation struct {
	DB     *gorm.DB
	Redis  *redis.Pool
	Worker *work.Enqueuer
}

func NewOperation() *Operation {
	conn := db.NewDB()
	redis := db.NewRedis()
	worker := work.NewEnqueuer(db.REDIS_WORKER_NAMESPACE, redis)
	return &Operation{DB: conn, Redis: redis, Worker: worker}
}

func Middleware(operation *Operation) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), operationCtxKey, operation)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *Operation {
	raw, _ := ctx.Value(operationCtxKey).(*Operation)
	return raw
}

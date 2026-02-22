package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

type traceIdKey struct{}

func newTraceId() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

func SetTraceId(ctx context.Context, traceId int) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}

func GetTraceId(ctx context.Context) int {
	id := ctx.Value(traceIdKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

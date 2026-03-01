package common

import (
	"context"
	"net/http"
)

type traceIdKey struct{}

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

type userNameKey struct{}

func GetUserName(ctx context.Context) string {
	id := ctx.Value(userNameKey{})

	if userNameStr, ok := id.(string); ok {
		return userNameStr
	}
	return ""
}

func SetUserName(req *http.Request, name string) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, userNameKey{}, name)
	req = req.WithContext(ctx)
	return req
}

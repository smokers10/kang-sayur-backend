package helper

import (
	"context"
	"time"
)

func InitCtxTimeout() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	return ctx, cancel
}

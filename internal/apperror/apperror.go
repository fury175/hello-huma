package apperror

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/fury175/hello-huma/internal/logger"
	"go.uber.org/zap"
)

type humaCtxKey struct{}

func GetHumaCtx(ctx context.Context) huma.Context {
	v, _ := ctx.Value(humaCtxKey{}).(huma.Context)
	return v
}

func MiddleWare(humaCtx huma.Context, next func(huma.Context)) {
	next(huma.WithValue(humaCtx, humaCtxKey{}, humaCtx))
}

func log() *zap.SugaredLogger {
	return logger.Log.WithOptions(zap.AddCallerSkip(1))
}

func logFields(ctx context.Context, err error) []any {
	humaCtx := GetHumaCtx(ctx)
	return []any{
		"method", humaCtx.Method(),
		"path", humaCtx.URL().Path,
		"error", err.Error(),
	}
}

func InternalServerError(ctx context.Context, err error) error {
	log().Errorw("Internal server error", logFields(ctx, err)...)
	return huma.Error500InternalServerError("something went wrong, try again")
}

func NotFoundError(ctx context.Context, err error) error {
	log().Warnw("not found", logFields(ctx, err)...)
	return huma.Error404NotFound(err.Error())
}

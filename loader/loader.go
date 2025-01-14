package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"

	"github.com/noissefnoc/go-Intern-Diary/service"
)

type Loaders interface {
	Attach(context.Context) context.Context
}

func New(app service.DiaryApp) Loaders {
	return &loaders{
		batchFuncs: map[string]dataloader.BatchFunc{
		// ...
		},
	}
}

type loaders struct {
	batchFuncs map[string]dataloader.BatchFunc
}

func (c *loaders) Attach(ctx context.Context) context.Context {
	for key, batchFn := range c.batchFuncs {
		ctx = context.WithValue(ctx, key, dataloader.NewBatchedLoader(batchFn))
	}
	return ctx
}

func getLoader(ctx context.Context, key string) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(key).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader from the request context", key)
	}
	return ldr, nil
}

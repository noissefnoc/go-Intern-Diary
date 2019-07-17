package resolver

import (
	"context"

	"github.com/noissefnoc/go-Intern-Diary/model"
	"github.com/noissefnoc/go-Intern-Diary/service"
)

type Resolver interface {
	// ...
}

func newResolver(app service.DiaryApp) Resolver {
	return &resolver{app: app}
}

type resolver struct {
	app service.DiaryApp
}

func currentUser(ctx context.Context) *model.User {
	return ctx.Value("user").(*model.User)
}

package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/oudrag/server/internal/domain/events"
	"github.com/oudrag/server/internal/platform/app"
	"github.com/oudrag/server/internal/platform/gqlcore"
)

type Resolver struct {
	ioc app.Container
}

func NewResolver(ioc app.Container) *Resolver { return &Resolver{ioc: ioc} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() gqlcore.MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() gqlcore.QueryResolver { return &queryResolver{r} }

func (r *Resolver) initializeCustomResolver(i interface{}) error {
	if needInit, ok := i.(app.HasInit); ok {
		return needInit.Init(r.ioc)
	}

	return nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Today(ctx context.Context, showDone *bool) ([]*events.Event, error) {
	panic("not implemented")
}

func (r *queryResolver) Event(ctx context.Context, id *string) (*events.Event, error) {
	panic("not implemented")
}

func (r *queryResolver) ThreeDays(ctx context.Context, showDone *bool) ([]*events.Event, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (m mutationResolver) NewEvent(ctx context.Context, input *gqlcore.NewEventData) (*events.Event, error) {
	panic("implement me")
}

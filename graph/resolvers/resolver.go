package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/ShintaNakama/sn_gqlserver/graph/generated"
	"github.com/ShintaNakama/sn_gqlserver/graph/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.ProjectInput) (bool, error) {
	return true, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	d1 := "testdis1"
	d2 := "testdis2"
	p1 := &model.Project{
		ID:          "1",
		Name:        "test1",
		Description: &d1,
		CreatedAt:   "",
		UpdatedAt:   "",
		Completed:   model.CompletedTypeComplete,
	}
	p2 := &model.Project{
		ID:          "2",
		Name:        "test2",
		Description: &d2,
		CreatedAt:   "20200725",
		UpdatedAt:   "",
		Completed:   model.CompletedTypeNotcomplete,
	}

	return []*model.Project{p1, p2}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

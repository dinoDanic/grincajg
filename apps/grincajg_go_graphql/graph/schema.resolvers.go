package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"grincajg/graph/controllers"
	"grincajg/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return controllers.CreateUser(ctx, input)
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (*model.Session, error) {
	return controllers.LoginUser(ctx, input)
}

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
	return controllers.CreateOrganization(ctx, input)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return controllers.Me(ctx)
}

// Organization is the resolver for the organization field.
func (r *userResolver) Organization(ctx context.Context, obj *model.User) (*model.Organization, error) {
	return controllers.PreloadOrganization(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

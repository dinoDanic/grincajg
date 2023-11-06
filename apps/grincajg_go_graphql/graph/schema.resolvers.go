package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"grincajg/graph/controllers"
	"grincajg/graph/model"
)

// Children is the resolver for the children field.
func (r *categoryResolver) Children(ctx context.Context, obj *model.Category) ([]*model.Category, error) {
	return controllers.GetCategoryChildrens(ctx, obj)
}

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

// CreateStore is the resolver for the createStore field.
func (r *mutationResolver) CreateStore(ctx context.Context, input model.CreateStoreInput) (*model.Store, error) {
	return controllers.CreateStore(ctx, input)
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	return controllers.CreateCategory(ctx, input)
}

// Stores is the resolver for the stores field.
func (r *organizationResolver) Stores(ctx context.Context, obj *model.Organization) ([]*model.Store, error) {
	return controllers.GetStores(ctx)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return controllers.Me(ctx)
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return controllers.GetCategories(ctx)
}

// GetOrganizationsOnMap is the resolver for the getOrganizationsOnMap field.
func (r *queryResolver) GetOrganizationsOnMap(ctx context.Context, input model.GetOrganizationsOnMapInput) ([]*model.Organization, error) {
	return controllers.GetOrganizationsOnMap(ctx, input)
}

// Organization is the resolver for the organization field.
func (r *userResolver) Organization(ctx context.Context, obj *model.User) (*model.Organization, error) {
	return controllers.GetMeOrganization(ctx)
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Organization returns OrganizationResolver implementation.
func (r *Resolver) Organization() OrganizationResolver { return &organizationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type categoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

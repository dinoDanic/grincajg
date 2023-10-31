package controllers

import (
	"context"
	"grincajg/database"
	"grincajg/graph/model"
	"grincajg/middleware"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetStores(ctx context.Context) ([]*model.Store, error) {

	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	database.DB.Preload("Organization.Stores").First(user)
	return user.Organization.Stores, nil
}

func CreateStore(ctx context.Context, input model.CreateStoreInput) (*model.Store, error) {
	user, err := middleware.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	database.DB.Preload("Organization").First(user)

	store := model.Store{
		Name:           input.Name,
		Address:        input.Address,
		OrganizationID: user.Organization.ID,
	}

	createdStore, err := store.SaveStore()

	if err != nil {
		return nil, gqlerror.Errorf("Failed crating store")
	}

	return createdStore, nil

}

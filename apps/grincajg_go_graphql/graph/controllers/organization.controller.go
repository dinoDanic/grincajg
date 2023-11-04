package controllers

import (
	"context"
	"grincajg/database"
	"grincajg/graph/model"
	"grincajg/middleware"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CreateOrganization(ctx context.Context, input model.CreateOrganizationInput) (*model.Organization, error) {
	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	newOrganization := model.Organization{
		AdminUserID: user.ID,
		Name:        input.Name,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
	}

	orgEntry, error := newOrganization.SaveOrganization()

	if error != nil {
		return nil, gqlerror.Errorf(error.Error())
	}

	user.OrganizationID = &orgEntry.ID

	result := database.DB.Save(&user)

	if result.Error != nil {
		return nil, gqlerror.Errorf(result.Error.Error())
	}

	return orgEntry, nil
}

func GetMeOrganization(ctx context.Context) (*model.Organization, error) {
	user, err := middleware.GetUser(ctx)

	if err != nil {
		return nil, err
	}

	database.DB.Preload("Organization").First(user)

	return user.Organization, nil
}

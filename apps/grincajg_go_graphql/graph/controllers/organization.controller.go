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

func GetOrganizationsOnMap(ctx context.Context, input model.GetOrganizationsOnMapInput) ([]*model.Organization, error) {
	ne := input.Northeast
	sw := input.Southwest

	// Ensure that the longitude and latitude values are in the correct order
	if ne.Latitude < sw.Latitude {
		ne.Latitude, sw.Latitude = sw.Latitude, ne.Latitude
	}
	if ne.Longitude < sw.Longitude {
		ne.Longitude, sw.Longitude = sw.Longitude, ne.Longitude
	}

	var organizations []*model.Organization
	err := database.DB.
		Where("latitude BETWEEN ? AND ?", sw.Latitude, ne.Latitude).
		Where("longitude BETWEEN ? AND ?", sw.Longitude, ne.Longitude).
		Find(&organizations).Error

	if err != nil {
		return nil, gqlerror.Errorf("Error retrieving organizations: %v", err)
	}

	return organizations, nil
}

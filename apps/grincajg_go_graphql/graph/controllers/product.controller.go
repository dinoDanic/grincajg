package controllers

import (
	"context"
	"grincajg/database"
	"grincajg/graph/model"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetProductsByCategoryOnMap(ctx context.Context, input model.ProductsByCategoryOnMapInput) ([]*model.Product, error) {
	ne := input.Northeast
	sw := input.Southwest

	categoryID := input.CategoryID

	// Ensure that the longitude and latitude values are in the correct order
	if ne.Latitude < sw.Latitude {
		ne.Latitude, sw.Latitude = sw.Latitude, ne.Latitude
	}
	if ne.Longitude < sw.Longitude {
		ne.Longitude, sw.Longitude = sw.Longitude, ne.Longitude
	}

	products := []*model.Product{}

	err := database.DB.Find(&products, "category_id = ?", categoryID).Error

	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	productsWithinScreen := []*model.Product{}

	for _, product := range products {
		database.DB.Preload("Organization").Find(product)
		if product.Organization.Latitude >= sw.Latitude &&
			product.Organization.Latitude <= ne.Latitude &&
			product.Organization.Longitude >= sw.Longitude &&
			product.Organization.Longitude <= ne.Longitude {
			productsWithinScreen = append(productsWithinScreen, product)
		}
	}

	return productsWithinScreen, nil
}

func GetProductCategory(obj *model.Product) (*model.Category, error) {
	database.DB.Preload("Category").Find(obj)
	return obj.Category, nil
}

func GetProductOrganization(obj *model.Product) (*model.Organization, error) {
	return obj.Organization, nil
}

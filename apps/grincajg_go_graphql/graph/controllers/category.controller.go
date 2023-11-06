package controllers

import (
	"context"
	"grincajg/database"
	"grincajg/graph/model"
	"grincajg/middleware"
	"strconv"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	_, err := middleware.GetSuperUser(ctx)

	if err != nil {
		return nil, err
	}

	category := model.Category{
		Name: input.Name,
	}

	if input.ParentCategoryID != nil {
		categoryID, err := strconv.Atoi(*input.ParentCategoryID)
		if err != nil {
			return nil, gqlerror.Errorf("Invalid ParentCategoryID: %s", err)
		}
		category.CategoryID = &categoryID
	}

	createdCategory, err := category.Save()

	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return createdCategory, nil
}

func GetCategories(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category

	database.DB.Where("category_id IS NULL").Find(&categories)
	// database.DB.Find(&categories)

	return categories, nil

}

func GetCategoryChildrens(ctx context.Context, obj *model.Category) ([]*model.Category, error) {
	database.DB.Preload("Childrens").Find(obj)
	return obj.Childrens, nil
}

package controllers

import (
	"context"
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

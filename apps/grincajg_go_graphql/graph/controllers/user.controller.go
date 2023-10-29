package controllers

import (
	"context"
	"grincajg/graph/model"
	"grincajg/middleware"
)

func Me(ctx context.Context) (*model.User, error) {
	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

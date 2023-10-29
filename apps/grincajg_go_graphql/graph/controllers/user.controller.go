package controllers

import (
	"context"
	"grincajg/graph/model"
	"grincajg/middleware"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Me(ctx context.Context) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return &model.User{}, gqlerror.Errorf("Not authorized")
	}
	return user, nil
}

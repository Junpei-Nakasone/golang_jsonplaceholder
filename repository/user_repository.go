package repository

import (
	"context"
	"golang_jsonplaceholder_without_framework/model"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User)
	Update(ctx context.Context, user model.User)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) (model.User, error)
	FindAll(ctx context.Context) []model.User
}

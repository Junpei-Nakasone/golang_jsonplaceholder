package service

import (
	"context"
	"golang_jsonplaceholder_without_framework/data/request"
	"golang_jsonplaceholder_without_framework/data/response"
)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest)
	Update(ctx context.Context, request request.UserUpdateRequest)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
}

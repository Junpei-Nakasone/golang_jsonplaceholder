package service

import (
	"context"
	"golang_jsonplaceholder_without_framework/data/request"
	"golang_jsonplaceholder_without_framework/data/response"
	"golang_jsonplaceholder_without_framework/helper"
	"golang_jsonplaceholder_without_framework/model"
	"golang_jsonplaceholder_without_framework/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

// Create implements UserService
func (u *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) {
	user := model.User{
		Name:     request.Name,
		UserName: request.UserName,
		Email:    request.Email,
	}
	u.UserRepository.Save(ctx, user)
}

// Delete implements UserService
func (u *UserServiceImpl) Delete(ctx context.Context, userId int) {
	user, err := u.UserRepository.FindById(ctx, userId)
	helper.PanicIfError(err)
	u.UserRepository.Delete(ctx, user.Id)
}

// FindAll implements UserService
func (u *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	users := u.UserRepository.FindAll(ctx)

	var userResp []response.UserResponse

	for _, value := range users {
		user := response.UserResponse{Id: value.Id, Name: value.Name, UserName: value.UserName, Email: value.Email}
		userResp = append(userResp, user)
	}
	return userResp
}

// FindById implements UserService
func (u *UserServiceImpl) FindById(ctx context.Context, userId int) response.UserResponse {
	user, err := u.UserRepository.FindById(ctx, userId)
	helper.PanicIfError(err)
	return response.UserResponse(user)
}

// Update implements UserService
func (u *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) {
	user, err := u.UserRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	user.Name = request.Name
	user.UserName = request.UserName
	user.Email = request.Email
	u.UserRepository.Update(ctx, user)

}

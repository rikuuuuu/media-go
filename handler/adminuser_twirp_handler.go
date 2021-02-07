package handler

import (
	"context"
	pb "mediago/pb"
	"mediago/repository"
)

type AdminUserService struct {
	contextProvider ContextProvider
}

type AdminUserHandler interface {
	GetMe(ctx context.Context, req *pb.Empty) (resp *pb.AdminUser, err error)
	GetAll(ctx context.Context, req *pb.Empty) (resp *pb.AdminUserList, err error)
	Create(ctx context.Context, req *pb.CreateAdminUserRequest) (resp *pb.AdminUser, err error)
	Update(ctx context.Context, req *pb.UpdateAdminUserRequest) (resp *pb.AdminUser, err error)
}

func NewAdminUserService() AdminUserHandler {
	c := NewContextProvider()
	return &AdminUserService{contextProvider: c}
}

func (a *AdminUserService) GetMe(ctx context.Context, req *pb.Empty) (resp *pb.AdminUser, err error) {
	
	user, err := repository.GetMe(ctx, req)

	return user, err
}

func (a *AdminUserService) GetAll(ctx context.Context, req *pb.Empty) (resp *pb.AdminUserList, err error) {

	users, err := repository.AdminUserGetAll(ctx, req)

	return users, err
}

func (a *AdminUserService) Create(ctx context.Context, req *pb.CreateAdminUserRequest) (resp *pb.AdminUser, err error) {

	user, err := repository.AdminUserCreate(ctx, req)

	return user, err
}

func (a *AdminUserService) Update(ctx context.Context, req *pb.UpdateAdminUserRequest) (resp *pb.AdminUser, err error) {

	user, err := repository.AdminUserUpdate(ctx, req)

	return user, err
}

func (a *AdminUserService) Delete(ctx context.Context, req *pb.AdminUserID) (resp *pb.Empty, err error) {

	empty, err := repository.AdminUserDelete(ctx, req)

	return empty, err
}
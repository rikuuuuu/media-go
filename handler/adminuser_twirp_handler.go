package handler

import (
	"context"
	"fmt"
	pb "mediago/pb"
	"mediago/repository"
)

type AdminUserService struct {
	contextProvider ContextProvider
}

func (a *AdminUserService) GetMe(ctx context.Context, req *pb.Empty) (resp *pb.AdminUser, err error) {

	fmt.Printf("%v", "getme")

	// uid := a.contextProvider.GetUID(ctx)
	// fmt.Printf("%v", uid)

	user, err := repository.GetMe(ctx, req)

	return user, err
}

func (a *AdminUserService) GetAll(ctx context.Context, req *pb.Empty) (resp *pb.AdminUserList, err error) {

	fmt.Printf("%v", "adminuser getall")

	users, err := repository.AdminUserGetAll(ctx, req)

	return users, err
}

func (a *AdminUserService) Create(ctx context.Context, req *pb.CreateAdminUserRequest) (resp *pb.AdminUser, err error) {

	fmt.Printf("%v", "article create")

	user, err := repository.AdminUserCreate(ctx, req)

	return user, err
}

func (a *AdminUserService) Update(ctx context.Context, req *pb.UpdateAdminUserRequest) (resp *pb.AdminUser, err error) {

	fmt.Printf("%v", "article update")

	user, err := repository.AdminUserUpdate(ctx, req)

	return user, err
}

func (a *AdminUserService) Delete(ctx context.Context, req *pb.AdminUserID) (resp *pb.Empty, err error) {

	fmt.Printf("%v", "article delete")

	empty, err := repository.AdminUserDelete(ctx, req)

	return empty, err
}
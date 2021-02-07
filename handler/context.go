package handler

import (
	"mediago/model"
	"context"
	// "fmt"
	// "log"
)

const authUidStoreKey = "__auth_admin_uid_store_key__"

type ContextProvider interface {
	WithAuthUID(ctx context.Context, uid model.UserID) (context.Context, error)
	GetUID(ctx context.Context) model.UserID
}

type contextProvider struct {}

func NewContextProvider() ContextProvider {
	return &contextProvider{}
}


func (u *contextProvider) WithAuthUID(ctx context.Context, uid model.UserID) (context.Context, error) {
	return context.WithValue(ctx, authUidStoreKey, uid), nil
}

func (u *contextProvider) GetUID(ctx context.Context) model.UserID {
	uid, _ := ctx.Value(authUidStoreKey).(model.UserID)
	// if !ok {
	// 	panic(domain.NewError(domain.ErrorTypeInternal, "server error"))
	// }
	return uid
}


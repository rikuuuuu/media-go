package handler

import (
	"mediago/model"
	"context"
	"fmt"
)

const authUidStoreKey = "__auth_admin_uid_store_key__"

type ContextProvider interface {
	WithAuthUID(ctx context.Context, uid model.UserID) (context.Context, error)
	GetUID(ctx context.Context) model.UserID
}

type contextProvider struct {}

// func NewContextProvider() ContextProvider {
// 	return &contextProvider{}
// }


func (u *contextProvider) WithAuthUID(ctx context.Context, uid model.UserID) (context.Context, error) {
	fmt.Printf("%v", "WithAuthUID")
	fmt.Printf("%v", uid)
	return context.WithValue(ctx, authUidStoreKey, uid), nil
}

func (u *contextProvider) GetUID(ctx context.Context) model.UserID {
	fmt.Printf("%v", "GetUID")
	fmt.Printf("%v", ctx)
	uid, _ := ctx.Value(authUidStoreKey).(model.UserID)
	fmt.Printf("%v", "uid")
	fmt.Printf("%v", uid)
	// if !ok {
	// 	panic(domain.NewError(domain.ErrorTypeInternal, "server error"))
	// }
	return uid
}


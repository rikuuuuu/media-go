package repository

import (
	// "time"
	"context"
	"log"
	"mediago/model"
	"cloud.google.com/go/firestore"
	pb "mediago/pb"
)

const authUidStoreKey = "__auth_admin_uid_store_key__"

func GetMe(ctx context.Context, req *pb.Empty) (resp *pb.AdminUser, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	uid, _ := ctx.Value(authUidStoreKey).(model.UserID)

	dsnap, err := client.Collection("user").Doc(uid.String()).Get(ctx)
	if err != nil {
        return nil, err
	}
	dsnap.DataTo(&resp)

	defer client.Close()

	return &pb.AdminUser{
		Id: resp.Id,
		Email: resp.Email,
		Name: resp.Name,
	}, nil
}

func AdminUserGetAll(ctx context.Context, req *pb.Empty) (resp *pb.AdminUserList, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	allData := client.Collection("user").Documents(ctx)
	docs, err := allData.GetAll()
	if err != nil {
		log.Fatalf("Failed adding getAll: %v", err)
	}

	users := make([]*pb.AdminUser, 0, len(docs))
	for _, doc := range docs {
		a := new(pb.AdminUser)
		mapToStruct(doc.Data(), &a)
		users = append(users, a)
	}

	return &pb.AdminUserList{
				Items: users,
			}, nil
}

func AdminUserCreate(ctx context.Context, req *pb.CreateAdminUserRequest) (resp *pb.AdminUser, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	uid, _ := ctx.Value(authUidStoreKey).(model.UserID)

	user := &model.CreateAdminUser{
		Id: uid.String(),
		Email: req.GetEmail(),
		Password: req.GetPassword(),
		Name: "",
	}

	_, err = client.Collection("user").Doc(uid.String()).Set(ctx, user)

	defer client.Close()

	return &pb.AdminUser{
				Email: req.GetEmail(),
			}, nil
}

func AdminUserUpdate(ctx context.Context, req *pb.UpdateAdminUserRequest) (resp *pb.AdminUser, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	uid, _ := ctx.Value(authUidStoreKey).(model.UserID)

	_, err = client.Collection("user").Doc(uid.String()).Set(ctx, map[string]interface{}{
		"name": req.GetName(),
	}, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	dsnap, err := client.Collection("user").Doc(uid.String()).Get(ctx)
	if err != nil {
        return nil, err
	}
	dsnap.DataTo(&resp)

	defer client.Close()

	return &pb.AdminUser{
		Id: resp.Id,
		Name: resp.Name,
		Email: resp.Email,
	}, nil
}

func AdminUserDelete(ctx context.Context, req *pb.AdminUserID) (resp *pb.Empty, err error) {

	client, err := firebaseInit(ctx)
		if err != nil {
			log.Fatal(err)
	}

	uid, _ := ctx.Value(authUidStoreKey).(model.UserID)

	_, err = client.Collection("user").Doc(uid.String()).Delete(ctx)
	if err != nil {
			log.Printf("An error has occurred: %s", err)
	}

	defer client.Close()

	return &pb.Empty{}, nil
}
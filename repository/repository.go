package repository

import (
	"log"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
	"encoding/json"
)

func firebaseInit(ctx context.Context) (*firestore.Client, error) {

	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return client, nil
}

func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
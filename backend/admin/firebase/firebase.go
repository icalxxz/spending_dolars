package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client
var FirestoreClient *firestore.Client

func InitFirebase() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("service.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	AuthClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}

	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

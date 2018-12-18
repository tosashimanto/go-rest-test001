package firebase

import (
	"firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

var App firebase.App

func init() {
	opt := option.WithCredentialsFile("google-services.json")
	println(opt)
	config := &firebase.Config{ProjectID: "marclay-84a44"}
	fmt.Println("ProjectID=", config.ProjectID)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	App = *app
}

func VerifyIDToken(ctx context.Context, idToken string) {
	fmt.Println("idToken=", idToken)

	client, err := App.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	log.Printf("Verified ID token: %v\n", token)
}

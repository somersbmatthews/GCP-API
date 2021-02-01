package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// type FirebaseAuth struct {
// 	firebase             *firebase.App
// 	defaultAuthenticator *auth.Client
// }

type UserInfo struct {
	UserId
	Name
	Email
}

func AuthenticateFirebaseUserWithToken(token string) (authenticated bool) {

}

func GetEmailFromFirebaseToken(token string) (email string) {

}

func GetUserInfo(token string) {

}

func newAuthenticator(cfg firebase.Config) *auth.Client {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	// Create a new authenticator for the app.
	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("app.Auth: %v", err)

		panic("app.Auth: %v")
	}
	return client
}

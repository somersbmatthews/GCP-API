package fba

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func VerifyToken(ctx context.Context, tokenStr string) bool {
	client := newAuth()

	_, err := client.VerifyIDToken(ctx, tokenStr)
	if err != nil {
		return false
	}

	return true
}

func newAuth() *auth.Client {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
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

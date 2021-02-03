package fba

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"github.com/dgrijalva/jwt-go"
)

// type FirebaseAuth struct {
// 	firebase             *firebase.App
// 	defaultAuthenticator *auth.Client
// }

type claims = jwt.MapClaims

func VerifyToken(ctx context.Context, tokenStr string) bool {
	client := newAuth()

	_, err := client.VerifyIDToken(ctx, tokenStr)
	if err != nil {
		return false
	}

	return true
}

func GetUserIdFromToken(ctx context.Context, tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, nil)
	if token == nil {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	// claims are actually a map[string]interface{}
	userId := claims["UID"].(string)
	return userId, nil
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

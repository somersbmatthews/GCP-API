package fba

import (
	"context"
	"encoding/base64"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func VerifyToken(ctx context.Context, tokenStr string) (string, bool) {
	client := newAuth()

	idToken, ok := parseBearerAuth(tokenStr)
	if !ok {
		return "", false
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", false
	}

	return token.UID, true
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

func parseBearerAuth(auth string) (token string, ok bool) {
	const prefix = "Bearer "
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return "", false
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return "", false
	}
	return string(c), true
}

package fba

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func VerifyToken(ctx context.Context, tokenStr string) (string, bool) {
	client := newAuth()

	// log.Printf("this is token str: %v", tokenStr)

	// idToken, ok := parseBearerAuth(tokenStr)
	// if !ok {
	// 	log.Printf("parseBearerAuth failed, this is idToken: %v", idToken)
	// 	return "", false
	// }

	authToken, err := client.VerifyIDToken(ctx, tokenStr)
	if err != nil {
		log.Printf("this is verify token err: %v", err)
		return "", false
	}

	// log.Printf("here is authToken: %v", authToken)

	return authToken.UID, true
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

// func parseBearerAuth(auth string) (token string, ok bool) {
// 	const prefix = "Bearer "
// 	// if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
// 	// 	log.Printf("len prefix fail: \n auth: \n %v \n prefix: \n %v \n", auth, prefix)
// 	// 	return "", false
// 	// }
// 	log.Printf("THIS IS BYTE 988: \n %v \n", auth[988])
// 	log.Printf("This is byte 987: \n %v \n", auth[987])
// 	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
// 	if err != nil {
// 		log.Printf("this is error from base64 decoding: \n %v", err)
// 		return "", false
// 	}
// 	return string(c[:]), true
// }

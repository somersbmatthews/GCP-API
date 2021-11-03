package auth

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

//	secretmanager "cloud.google.com/go/secretmanager/apiv1"
//	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

var oldDataToken string

var tokenSecret string

func init() {
	// oldDataTokenValue, err := accessOldDataToken()
	// if err != nil {
	// 	panic(err)
	// }
	// oldDataToken = oldDataTokenValue

	// tokenSecretValue, err := accessTokenSecret()
	// if err != nil {
	// 	panic(err)
	// }
	// tokenSecret = tokenSecretValue
}

// func VerifyOldDataToken(context.Context, string) bool {
// 	token, err := accessOldDataToken()

// }

func accessOldDataToken() (string, error) {
	name := "projects/gircapp/secrets/GIRC_OLD_DATA_AUTH_TOKEN/versions/latest"
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}
	return string(result.Payload.Data), nil
}

func VerifyOldDataToken(token string) bool {
	if token == oldDataToken {
		return true
	} else {
		return false
	}
}

// func VerifyToken(context.Context, string) bool {
// 	token, err := accessTokenSecret()

// }

// func accessTokenSecret() (string, error) {
// 	name := "projects/gircapp/secrets/TOKEN_SECRET/versions/latest"
// 	ctx := context.Background()
// 	client, err := secretmanager.NewClient(ctx)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
// 	}
// 	req := &secretmanagerpb.AccessSecretVersionRequest{
// 		Name: name,
// 	}
// 	result, err := client.AccessSecretVersion(ctx, req)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to access secret version: %v", err)
// 	}
// 	return string(result.Payload.Data), nil
// }

// func generateRefreshToken() (string, error) {

// }

// func generateAccessToken() (string, error) {

// }

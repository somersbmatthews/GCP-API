package pg

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/gircapp/api/secret"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var encyptionKey string

var iv = []byte("1234567812345679")

func init() {
	key, err := secret.GetSecret("ENCRYPTIONKEY")
	if err != nil {
		panic(err)
	}
	encyptionKey = key
}

func getEncryptionKey() (string, error) {
	name := "projects/gircapp/secrets/ENCRYPTIONKEY/versions/latest"
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

func encryptUserID(userID string) (string, error) {
	key := []byte(encyptionKey)
	bytes := []byte(userID)
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create the AES cipher: %v)", err)
		panic(errMsg)
	}
	stream := cipher.NewCTR(blockCipher, iv)
	stream.XORKeyStream(bytes, bytes)
	return string(bytes[:]), nil
}

func decryptUserID(encryptedUserID string) (string, error) {
	key := []byte(encyptionKey)
	bytes := []byte(encryptedUserID)
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	stream := cipher.NewCTR(blockCipher, iv)
	stream.XORKeyStream(bytes, bytes)
	return string(bytes[:]), nil
}

package fba

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/gdexlab/go-render/render"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

const uid string = "1234567890"

// const uid string = "LdXjx5tA0zXj1FqaghkiAl9rkuL2"
const baduid string = "1234567790"

const badToken string = "alkdjfalksjfa87wet235"

var bearerToken string

func TestCreateFirebaseUser(t *testing.T) {
	ctx := context.Background()

	ok := createFirebaseUser(ctx, uid)
	if !ok {
		t.Error("create firebase user failed.")
	}
}

func TestCreateBearerToken(t *testing.T) {

	token, ok := createBearerToken()
	if !ok {
		t.Error("create firebase user failed.")
	}

	bearerToken = token
}

func createBearerToken() (string, bool) {
	ctx := context.Background()
	idToken, err := getIDTokenForUser(ctx, uid)
	if err != nil {
		return "", false
	}
	// TODO write code to make bearertoken

	bearerToken := fmt.Sprintf("%s%s", "Bearer ", idToken)

	idTokenBytes := []byte(bearerToken)
	return base64.StdEncoding.EncodeToString(idTokenBytes), true

}

func TestAuthenticateFirebaseUser(t *testing.T) {
	ctx := context.Background()

	client := newAuth()

	token, err := client.VerifyIDToken(ctx, bearerToken)
	if err != nil {
		t.Error("verify firebase user failed.")
	}

	if token.UID == uid {
		t.Errorf("token.UID %v does not equal uid %v", token.UID, uid)
	}

	// ok = VerifyToken(ctx, badToken)
	// if ok {
	// 	t.Error("expected verify firebase user have failed, but it succeeded.")
	// }
}

func TestDeleteFirebaseUser(t *testing.T) {
	ctx := context.Background()

	ok := deleteFirebaseUser(ctx, uid)
	if !ok {
		t.Error("delete firebase user failed.")
	}

	// ok = deleteFirebaseUser(ctx, baduid)
	// if ok {
	// 	t.Error("expected verify firebase user have failed, but it succeeded.")
	// }
}

// func TestFirebaseUserDeleted(t *testing.T) {
// 	ctx := context.Background()

// 	ok := deleteFirebaseUser(ctx, uid)
// 	if ok {
// 		t.Error("expected verify firebase user have failed, but it succeeded.")
// 	}
// }

func getIDTokenForUser(ctx context.Context, uid string) (string, error) {

	client := newAuth()

	customToken, err := client.CustomToken(ctx, uid)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	apiKey, ok := os.LookupEnv("FIREBASEAPIKEY")
	if !ok {
		return "", errors.New("can't find api key env var")
	}

	booleanTrue := true

	trueStr := strconv.FormatBool(booleanTrue)

	reqBody, err := json.Marshal(map[string]string{
		"token":             customToken,
		"returnSecureToken": trueStr,
	})
	if err != nil {
		return "", errors.New("could not marshal json")
	}

	urlStr := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s", apiKey)

	resp, err := http.Post(urlStr, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", errors.New("cannot post request")
	}
	defer resp.Body.Close()

	response := make(map[string]interface{})

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("read resp.Body as a slice of byte")
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", errors.Errorf("cannot unmarshal json, error: %v", err)
	}

	idTokenInterface := response["idToken"]
	if idTokenInterface == nil {
		return "", errors.New("cannot read idToken field from json")
	}

	idTokenStr := idTokenInterface.(string)

	return idTokenStr, nil
}

func createFirebaseUser(ctx context.Context, uid string) bool {
	client := newAuth()

	params := (&auth.UserToCreate{}).
		UID(uid).
		Email("somersbmatthews2@gmail.com").
		EmailVerified(true).
		Password("123456")
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
		return false
	}
	log.Printf("Successfully created user: %v\n", u)
	return true
}

func deleteFirebaseUser(ctx context.Context, uid string) bool {
	client := newAuth()

	err := client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
		return false
	}
	log.Printf("Successfully deleted user: %s\n", uid)
	return true
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	client := newAuth()

	iter := client.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Errorf("error listing users: %s\n", err)
		}
		fmt.Printf("read user user: %v\n", render.Render(user))
	}

}

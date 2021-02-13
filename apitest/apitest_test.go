package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"testing"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gdexlab/go-render/render"
	"github.com/pkg/errors"
)

type body map[string]interface{}

const uid string = "1234567890"

const urlstr string = "http://127.0.0.1:8080/v2"

var token string

func init() {
	ctx := context.Background()
	tokenStr, err := getIDTokenForUser(ctx, uid)
	if err != nil {
		panic(err)
	}
	token = tokenStr
}

func TestRegisterUser(t *testing.T) {

	reqBody := body{
		"userId":     "123456789022",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
	}

	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/user", urlstr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	t.Log("THIS CODE IS RUNNING")

	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to register user, error: %v", err)
	}

	json := getBody(*resp)
	t.Log("THIS CODE IS RUNNING TOO")
	t.Log(render.Render(json))

	want := body{
		"userId":     "123456789022",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
		"created":    true,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

// func TestGetUser(t *testing.T) {
// 	reqBody := body{
// 		"userId": "1234567890",
// 	}
// 	req := newRequest()
// 	reqBodyData, err := setBody(reqBody)
// 	if err != nil {
// 		t.Errorf("could not set json into readcloser %v", err)
// 	}
// 	req.Body = reqBodyData
// 	req.Method = "GET"
// 	req.URL, _ = url.Parse(urlstr)
// 	client := newClient()
// 	resp, err := client.Do(&req)
// 	if err != nil {
// 		t.Errorf("failed to register user, error: %v", err)
// 	}

// 	json := getBody(*resp)

// 	want := body{
// 		"userId":     "1234567890",
// 		"name":       "Tee Bow",
// 		"email":      "tbow@gmail.com",
// 		"speciality": "otolaryngologist",
// 		"degree":     "MD",
// 		"verified":   false,
// 	}

// 	if !reflect.DeepEqual(want, json) {
// 		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
// 	}
// }

// // TODO: fix update user
// func TestUpdateUser(t *testing.T) {

// 	reqBody := body{
// 		"userId":     "1234567890",
// 		"name":       "Tee Bow",
// 		"email":      "tbow@gmail.com",
// 		"speciality": "otolaryngologist",
// 		"degree":     "MD",
// 	}
// 	req := newRequest()
// 	reqBodyData, err := setBody(reqBody)
// 	if err != nil {
// 		t.Errorf("could not set json into readcloser %v", err)
// 	}
// 	req.Body = reqBodyData
// 	req.Method = "POST"
// 	req.URL, _ = url.Parse(urlstr)
// 	client := newClient()
// 	resp, err := client.Do(&req)
// 	if err != nil {
// 		t.Errorf("failed to register user, error: %v", err)
// 	}

// 	json := getBody(*resp)

// 	want := body{
// 		"userId":     "1234567890",
// 		"name":       "Tee Bow",
// 		"email":      "tbow@gmail.com",
// 		"speciality": "otolaryngologist",
// 		"degree":     "MD",
// 		"created":    true,
// 	}

// 	if !reflect.DeepEqual(want, json) {
// 		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
// 	}
// }

// func TestDeleteUser(t *testing.T) {
// 	reqBody := body{
// 		"userId": "1234567890",
// 	}
// 	req := newRequest()
// 	reqBodyData, err := setBody(reqBody)
// 	if err != nil {
// 		t.Errorf("could not set json into readcloser %v", err)
// 	}
// 	req.Body = reqBodyData
// 	req.Method = "DELETE"
// 	req.URL, _ = url.Parse(urlstr)
// 	client := newClient()
// 	_, err = client.Do(&req)
// 	if err != nil {
// 		t.Errorf("failed to delete user, error: %v", err)
// 	}
// }

// // TODO: fix verify user
// func TestVerifyUser(t *testing.T) {

// 	reqBody := body{
// 		"userId":   "1234567890",
// 		"verified": true,
// 	}
// 	req := newRequest()
// 	reqBodyData, err := setBody(reqBody)
// 	if err != nil {
// 		t.Errorf("could not set json into readcloser %v", err)
// 	}
// 	req.Body = reqBodyData
// 	req.Method = "PATCH"
// 	req.URL, _ = url.Parse(urlstr)
// 	client := newClient()
// 	resp, err := client.Do(&req)
// 	if err != nil {
// 		t.Errorf("failed to register user, error: %v", err)
// 	}

// 	json := getBody(*resp)

// 	want := body{
// 		"userId":     "1234567890",
// 		"name":       "Tee Bow",
// 		"email":      "tbow@gmail.com",
// 		"speciality": "otolaryngologist",
// 		"degree":     "MD",
// 		"created":    true,
// 	}

// 	if !reflect.DeepEqual(want, json) {
// 		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
// 	}

// }

// func TestCreateIncident(t *testing.T) {

// 	reqBody := body{
// 		"ID":                            "1234567890",
// 		"DateOfIncident":                "12/20/2020",
// 		"ApproximatePatientAge":         "34",
// 		"Gender":                        "female",
// 		"LongTermPrognosis":             "dead",
// 		"IncidentDescription":           "choking",
// 		"Anterior":                      "someurl@url.com",
// 		"ObjectConsistency":             "rough",
// 		"ObjectBasicShape":              "round",
// 		"WhatMaterialIsTheObjectMadeOf": "plastic",
// 		"TheObjectIs":                   "small",
// 		"LargestLength":                 "23",
// 	}
// 	req := newRequest()
// 	reqBodyData, err := setBody(reqBody)
// 	if err != nil {
// 		t.Errorf("could not set json into readcloser %v", err)
// 	}
// 	req.Body = reqBodyData
// 	req.Method = "POST"
// 	req.URL, _ = url.Parse(urlstr)
// 	client := newClient()
// 	resp, err := client.Do(&req)
// 	if err != nil {
// 		t.Errorf("failed to register user, error: %v", err)
// 	}

// 	json := getBody(*resp)

// 	want := body{
// 		"ID":                            "1234567890",
// 		"DateOfIncident":                "12/20/2020",
// 		"ApproximatePatientAge":         "34",
// 		"Gender":                        "female",
// 		"LongTermPrognosis":             "dead",
// 		"IncidentDescription":           "choking",
// 		"Anterior":                      "someurl@url.com",
// 		"ObjectConsistency":             "rough",
// 		"ObjectBasicShape":              "round",
// 		"WhatMaterialIsTheObjectMadeOf": "plastic",
// 		"TheObjectIs":                   "small",
// 		"LargestLength":                 "23",
// 		"Created":                       true,
// 	}

// 	if !reflect.DeepEqual(want, json) {
// 		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
// 	}
// }

// func TestGetIncidents(t *testing.T) {// func TestUpdateIncident(t *testing.T) {

// }
// func TestUpdateIncident(t *testing.T) {

// }

// func TestDeleteIncident(t *testing.T) {

// }

func setBody(body body) ([]byte, error) {

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// datareader := bytes.NewBuffer(data)

	// datacloser := ioutil.NopCloser(datareader)

	return data, nil
}

func getBody(req http.Response) body {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	jsondata := body{}
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		panic(err)
	}

	return jsondata
}

func newHeader() http.Header {
	return http.Header{}
}

func newRequest() http.Request {
	return http.Request{}
}

func newClient() http.Client {
	return http.Client{}
}

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

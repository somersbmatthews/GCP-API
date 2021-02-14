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
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gdexlab/go-render/render"
	"github.com/pkg/errors"
)

type body map[string]interface{}

const uid string = "1234567890"

// const urlstr string = "http://127.0.0.1:8080/v2"

// const urlstr string = "https://35.208.147.207/v2"

const urlstr string = "https://girc.app/v2"

var token string

func init() {
	ctx := context.Background()
	tokenStr, err := getIDTokenForUser(ctx, uid)
	if err != nil {

		fmt.Println(err)
		panic(err)
	}
	token = tokenStr
}

func TestRegisterUser(t *testing.T) {
	t.Logf("this is token: %v", token)
	t.Logf("this is running")
	reqBody := body{
		"userId":     "1234567890",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
	}
	t.Logf("this is running 2")
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}
	t.Logf("this is running 3")
	url := fmt.Sprintf("%v/user", urlstr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}
	t.Logf("this is running 4")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to register user, error: %v", err)
	}
	t.Logf("this is running 5")
	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
		"created":    true,
	}
	t.Logf("this is running 6")
	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

func TestGetUser(t *testing.T) {
	reqBody := body{
		"userId": "1234567890",
	}
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/user", urlstr)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to get user, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
		"verified":   false,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

// // TODO: fix update user
func TestUpdateUser(t *testing.T) {
	reqBody := body{
		"userId":     "1234567890",
		"name":       "Tee H.W. Bow",
		"email":      "tbowSD@gmail.com",
		"speciality": "spin doctor",
		"degree":     "SD",
	}
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/user", urlstr)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to update user, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
		"name":       "Tee H.W. Bow",
		"email":      "tbowSD@gmail.com",
		"speciality": "spin doctor",
		"degree":     "SD",
		"verified":   false,
		"updated":    true,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

func TestVerifyUser(t *testing.T) {

	reqBody := body{
		"userId":   "1234567890",
		"verified": true,
	}
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/verify", urlstr)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to verify user, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
		"name":       "Tee H.W. Bow",
		"email":      "tbowSD@gmail.com",
		"speciality": "spin doctor",
		"degree":     "SD",
		"verified":   true,
		"updated":    true,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}

}

func TestDeleteUser(t *testing.T) {
	reqBody := body{
		"userId": "1234567890",
	}
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/user", urlstr)

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to delete user, error: %v", err)
	}
	json := getBody(*resp)

	want := body{
		"userId":  "1234567890",
		"deleted": true,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

// // TODO: fix verify user

func TestCreateIncident(t *testing.T) {
	reqBody := body{
		"ID":                                  "1234567890",
		"Date_of_Incident":                    "12/20/2020",
		"Approximate_Patient_Age":             "34",
		"Gender":                              "female",
		"Long-term_prognosis":                 "dead",
		"Incident_Description":                "choking",
		"Anterior":                            "someurl@url.com",
		"Object_Consistency":                  "rough",
		"Object_Basic_Shape":                  "round",
		"What_material_is_the_object_made_of": "plastic",
		"The_object_is":                       "small",
		"Largest_Length":                      "23",
	}
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/incident", urlstr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to create incident, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"ID":                                  "1234567890",
		"Date_of_Incident":                    "12/20/2020",
		"Approximate_Patient_Age":             "34",
		"Gender":                              "female",
		"Long-term_prognosis":                 "dead",
		"Incident_Description":                "choking",
		"Anterior":                            "someurl@url.com",
		"Object_Consistency":                  "rough",
		"Object_Basic_Shape":                  "round",
		"What_material_is_the_object_made_of": "plastic",
		"The_object_is":                       "small",
		"Largest_Length":                      "23",
		"Created":                             true,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

func TestUpdateIncident(t *testing.T) {
	reqBody := body{
		"ID":                                  "1234567890",
		"Date_of_Incident":                    "12/20/2020",
		"Approximate_Patient_Age":             "34",
		"Gender":                              "female",
		"Long-term_prognosis":                 "dead",
		"Incident_Description":                "choking",
		"Anterior":                            "someurl@url.com",
		"Object_Consistency":                  "rough",
		"Object_Basic_Shape":                  "round",
		"What_material_is_the_object_made_of": "plastic",
		"The_object_is":                       "small",
		"Largest_Length":                      "23",
	}
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/incident", urlstr)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to create incident, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"ID":                                  "1234567890",
		"Date_of_Incident":                    "12/20/2020",
		"Approximate_Patient_Age":             "34",
		"Gender":                              "female",
		"Long-term_prognosis":                 "dead",
		"Incident_Description":                "choking",
		"Anterior":                            "someurl@url.com",
		"Object_Consistency":                  "rough",
		"Object_Basic_Shape":                  "round",
		"What_material_is_the_object_made_of": "plastic",
		"The_object_is":                       "small",
		"Largest_Length":                      "23",
		"Updated":                             true,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

func TestDeleteIncident(t *testing.T) {
	reqBody := body{
		"ID": "1234567890",
	}
	data, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not convert reqBody map[string]interface to []byte, error: %v", err)
	}

	url := fmt.Sprintf("%v/incident", urlstr)

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("could not make new request %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := newClient()
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("failed to create incident, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"ID":      "1234567890",
		"Deleted": true,
	}

	if !reflect.DeepEqual(want, json) {
		t.Errorf("response json: \n %v \n does not equal json in request: \n %v \n.", render.Render(json), render.Render(want))
	}
}

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

func newClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	return client
}

func getIDTokenForUser(ctx context.Context, uid string) (string, error) {

	client := newAuth()

	customToken, err := client.CustomToken(ctx, uid)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	apiKey, ok := os.LookupEnv("FIREBASEAPIKEY")
	if !ok {
		return "", errors.Errorf("can't find api key env var, error: %v", err)
	}

	booleanTrue := true

	trueStr := strconv.FormatBool(booleanTrue)

	reqBody, err := json.Marshal(map[string]string{
		"token":             customToken,
		"returnSecureToken": trueStr,
	})
	if err != nil {
		return "", errors.Errorf("could not marshal json, error: %v", err)
	}

	urlStr := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s", apiKey)

	resp, err := http.Post(urlStr, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", errors.Errorf("cannot post request, error: %v", err)
	}
	defer resp.Body.Close()

	response := make(map[string]interface{})

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Errorf("read resp.Body as a slice of byte, error: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", errors.Errorf("cannot unmarshal json, error: %v", err)
	}

	idTokenInterface := response["idToken"]
	if idTokenInterface == nil {
		return "", errors.Errorf("cannot read idToken field from json, error: %v", err)
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

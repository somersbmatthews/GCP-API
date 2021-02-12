package test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/gdexlab/go-render/render"
)

type body map[string]interface{}

const urlstr string = ""

func TestRegisterUser(t *testing.T) {

	reqBody := body{
		"userId":     "1234567890",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
	}
	req := newRequest()
	reqBodyData, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not set json into readcloser %v", err)
	}
	req.Body = reqBodyData
	req.Method = "POST"
	req.URL, _ = url.Parse(urlstr)
	client := newClient()
	resp, err := client.Do(&req)
	if err != nil {
		t.Errorf("failed to register user, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
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

func TestGetUser(t *testing.T) {
	reqBody := body{
		"userId":     "1234567890",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
	}
	req := newRequest()
	reqBodyData, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not set json into readcloser %v", err)
	}
	req.Body = reqBodyData
	req.Method = "POST"
	req.URL, _ = url.Parse(urlstr)
	client := newClient()
	resp, err := client.Do(&req)
	if err != nil {
		t.Errorf("failed to register user, error: %v", err)
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

// TODO: fix update user
func TestUpdateUser(t *testing.T) {

	reqBody := body{
		"userId":     "1234567890",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
	}
	req := newRequest()
	reqBodyData, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not set json into readcloser %v", err)
	}
	req.Body = reqBodyData
	req.Method = "POST"
	req.URL, _ = url.Parse(urlstr)
	client := newClient()
	resp, err := client.Do(&req)
	if err != nil {
		t.Errorf("failed to register user, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
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

func TestDeleteUser(t *testing.T) {
	reqBody := body{
		"userId": "1234567890",
	}
	req := newRequest()
	reqBodyData, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not set json into readcloser %v", err)
	}
	req.Body = reqBodyData
	req.Method = "DELETE"
	req.URL, _ = url.Parse(urlstr)
	client := newClient()
	_, err = client.Do(&req)
	if err != nil {
		t.Errorf("failed to delete user, error: %v", err)
	}
}

// TODO: fix verify user
func TestVerifyUser(t *testing.T) {

	reqBody := body{
		"userId":     "1234567890",
		"name":       "Tee Bow",
		"email":      "tbow@gmail.com",
		"speciality": "otolaryngologist",
		"degree":     "MD",
	}
	req := newRequest()
	reqBodyData, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not set json into readcloser %v", err)
	}
	req.Body = reqBodyData
	req.Method = "POST"
	req.URL, _ = url.Parse(urlstr)
	client := newClient()
	resp, err := client.Do(&req)
	if err != nil {
		t.Errorf("failed to register user, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
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

func TestCreateIncident(t *testing.T) {

	reqBody := body{
		"ID":                            "1234567890",
		"DateOfIncident":                "12/20/2020",
		"ApproximatePatientAge":         "34",
		"Gender":                        "female",
		"LongTermPrognosis":             "dead",
		"IncidentDescription":           "choking",
		"Anterior":                      "someurl@url.com",
		"ObjectConsistency":             "rough",
		"ObjectBasicShape":              "round",
		"WhatMaterialIsTheObjectMadeOf": "plastic",
		"TheObjectIs":                   "small",
		"LargestLength":                 "23",
	}
	req := newRequest()
	reqBodyData, err := setBody(reqBody)
	if err != nil {
		t.Errorf("could not set json into readcloser %v", err)
	}
	req.Body = reqBodyData
	req.Method = "POST"
	req.URL, _ = url.Parse(urlstr)
	client := newClient()
	resp, err := client.Do(&req)
	if err != nil {
		t.Errorf("failed to register user, error: %v", err)
	}

	json := getBody(*resp)

	want := body{
		"userId":     "1234567890",
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

// func TestGetIncidents(t *testing.T) {// func TestUpdateIncident(t *testing.T) {

// }
// func TestUpdateIncident(t *testing.T) {

// }

// func TestDeleteIncident(t *testing.T) {

// }

func setBody(body body) (io.ReadCloser, error) {

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	datareader := bytes.NewBuffer(data)

	datacloser := ioutil.NopCloser(datareader)

	return datacloser, nil
}

func getBody(req http.Response) map[string]interface{} {

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

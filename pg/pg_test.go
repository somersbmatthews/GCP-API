package pg

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gdexlab/go-render/render"
	"github.com/gircapp/api/models"

	"gorm.io/gorm"
)

const testUserID string = "1234567890"

const badUserID string = "1234567790"

// func TestBytea(t *testing.T) {
// 	encryptedUserID, err := encryptUserID(testUserID)
// 	if err != nil {
// 		panic(err)
// 	}

// 	bytea := getByteaFromString(encryptedUserID)

// 	encryptedUserIDFromBytea := getStringFromBytea(bytea)

// 	if encryptedUserID != encryptedUserIDFromBytea {
// 		t.Errorf("error getting string from bytea \n %v \n does not equal \n %v \n", encryptedUserID, encryptedUserIDFromBytea)
// 	}
// }
func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	user := User{
		UserID:    testUserID,
		Email:     "DrJimBob@jimbobclinic.com",
		Name:      "Jim Bob",
		Specialty: "Otolaryngologist",
		Degree:    "MD",
		Verified:  false,
	}
	payload, ok := CreateUser(ctx, user)
	if !ok {
		t.Errorf("postgres create user \n %v \n failed", user)
	}

	if payload.Degree != user.Degree ||
		payload.Name != user.Name ||
		payload.Email != user.Email ||
		*payload.UserID != testUserID ||
		payload.Specialty != user.Specialty {
		t.Errorf("postgres create user failed, \n payload %v \n does not match user \n %v", render.Render(payload), render.Render(user))
	}
}

func TestGetUser(t *testing.T) {

	user := User{
		UserID:    "",
		Email:     "DrJimBob@jimbobclinic.com",
		Name:      "Jim Bob",
		Specialty: "Otolaryngologist",
		Degree:    "MD",
		Verified:  false,
	}

	ctx := context.Background()

	payload, ok := GetUser(ctx, testUserID)

	if !ok {
		t.Errorf("postgres get user by user id: %v failed", testUserID)
	}
	if *payload.Degree != user.Degree ||
		*payload.Name != user.Name ||
		*payload.Email != user.Email ||
		*payload.UserID != testUserID ||
		*payload.Specialty != user.Specialty {
		t.Errorf("postgres get user failed, \n payload %v \n does not match user \n %v", render.Render(payload), render.Render(user))
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()

	user := User{
		UserID:    testUserID,
		Email:     "DrJimBob@jimbobclinic.com",
		Name:      "Jim Bob",
		Degree:    "MD",
		Specialty: "Otolaryngologist",
	}

	updateWithUser := User{
		UserID:    testUserID,
		Email:     "DrJimBobNewEmail@jimbobclinic.com",
		Name:      "Jim W. Bob",
		Degree:    "DO",
		Specialty: "Otolaryngologist",
	}

	badUser := User{
		UserID:    badUserID,
		Email:     "DrJimBobNewEmail@jimbobclinic.com",
		Name:      "Jim W. Bob",
		Degree:    "DO",
		Specialty: "SpinDoctor",
	}

	updateWithEmptyFields := User{
		UserID:    testUserID,
		Email:     "DrJimHWBobNewEmail@jimbobclinic.com",
		Name:      "Jim HW. Bob",
		Degree:    "",
		Specialty: "",
	}

	payload, ok := UpdateUser(ctx, updateWithUser)
	if !ok {
		t.Errorf("postgres user update failed for user: %v \n with updatewithuser: %v", user, updateWithUser)
	}

	if *payload.Degree != updateWithUser.Degree ||
		*payload.Name != updateWithUser.Name ||
		*payload.Email != updateWithUser.Email ||
		*payload.UserID != testUserID ||
		*payload.Specialty != updateWithUser.Specialty {
		t.Errorf("postgres update user failed, \n payload %v \n does not match updatewithuser \n %v", render.Render(payload), updateWithUser)
	}

	db, conn := Open()
	defer conn.Close()
	err := db.First(&User{}, "user_id = ?", badUserID).Omit("user_id").Updates(badUser).Error
	if err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		t.Errorf("expected postgres update user to have failed, \n by baduserID: %v", badUser.UserID)
	}

	_, ok = UpdateUser(ctx, updateWithEmptyFields)
	if !ok {
		t.Errorf("postgres user update with empty fields failed for user: %v \n with updateWithEmptyFields: %v", user, updateWithEmptyFields)
	}

	payloadAfterEmptyUpdates, ok := GetUser(ctx, testUserID)

	if *payloadAfterEmptyUpdates.Degree != updateWithUser.Degree ||
		*payloadAfterEmptyUpdates.Name != updateWithEmptyFields.Name ||
		*payloadAfterEmptyUpdates.Email != updateWithEmptyFields.Email ||
		*payloadAfterEmptyUpdates.UserID != testUserID ||
		*payloadAfterEmptyUpdates.Specialty != updateWithUser.Specialty {
		t.Errorf("postgres update user with empty fields failed, \n payload %v \n does not match updateWithEmptyFields \n %v", render.Render(payloadAfterEmptyUpdates), updateWithEmptyFields)
	}

}

func TestVerifyUser(t *testing.T) {
	ctx := context.Background()

	booleanTrue := true

	verify := models.Verify{
		Verified: &booleanTrue,
	}

	_, ok := VerifyUser(ctx, verify, testUserID)
	if !ok {
		t.Errorf("postgres verify user failed, verified with verify: %v", render.Render(verify))
	}

	payload := User{}

	db, conn := Open()
	defer conn.Close()

	err := db.First(&payload, "user_id = ?", testUserID).Error
	if err == gorm.ErrRecordNotFound {
		t.Errorf("could not find user: %v after verifying", testUserID)
	} else if err != nil {
		t.Error(err)
	}

	if payload.Verified == false {
		t.Errorf("postgres verify user failed, \n payload %v \n does not match updatewithincident \n %v", payload, render.Render(verify))
	}

	if payload.Verified != true {
		t.Errorf("postgres verify user failed, \n payload %v \n does not match updatewithincident \n %v", payload, render.Render(verify))
	}

}

func TestDeleteUser(t *testing.T) {

	ctx := context.Background()
	_, ok := DeleteUser(ctx, testUserID)
	if !ok {
		t.Error("error deleted user")
	}

	_, ok = GetUser(ctx, testUserID)
	if ok {
		t.Error("expected get deleted user to have failed, but it succeeded")
	}
}

func TestCreateIncidents(t *testing.T) {
	ctx := context.Background()

	ID := "1234567890"
	ID2 := "1234567790"

	incident1 := models.CreateIncident{
		ID:                            &ID,
		DateOfIncident:                "30/80/2021",
		ApproximatePatientAge:         "34",
		Gender:                        "female",
		LongTermPrognosis:             "dead",
		IncidentDescription:           "choking",
		Anterior:                      "someurl@url.com",
		ObjectConsistency:             "rough",
		ObjectBasicShape:              "round",
		WhatMaterialIsTheObjectMadeOf: "plastic",
		TheObjectIs:                   "small",
		LargestLength:                 "23",
		LocationOfObject:              "throat",
	}

	incident2 := models.CreateIncident{
		ID:                            &ID2,
		DateOfIncident:                "30/80/2021",
		ApproximatePatientAge:         "34",
		Gender:                        "female",
		LongTermPrognosis:             "dead",
		IncidentDescription:           "choking",
		Anterior:                      "someurl@url.com",
		ObjectConsistency:             "rough",
		ObjectBasicShape:              "round",
		WhatMaterialIsTheObjectMadeOf: "plastic",
		TheObjectIs:                   "small",
		LargestLength:                 "23",
		LocationOfObject:              "throat",
	}

	_ = CreateIncident(ctx, incident1, testUserID)

	_ = CreateIncident(ctx, incident2, testUserID)

}

func TestGetBytea(t *testing.T) {
	db, conn := Open()
	defer conn.Close()

	encryptedUserID, err := encryptUserID(testUserID)
	if err != nil {
		panic(err)
	}

	bytea := getByteaFromString(encryptedUserID)

	incident := Incident{}

	sql := "SELECT * FROM incidents WHERE encrypted_user_id = ? LIMIT 1"

	err = db.Raw(sql, bytea).Scan(&incident).Error

	// err = db.Where(&incident, "encrypted_user_id = ?", bytea).First(&incident).Error
	if err == gorm.ErrRecordNotFound {
		t.Errorf("could not find encrypted_user_id: \n %v \n after trying to get an incident", encryptedUserID)
	} else if err != nil {
		t.Error(err)
	}
	t.Logf("this is encrypted user ID from database: \n %v \n ", incident.EncryptedUserID)
	if incident.EncryptedUserID != encryptedUserID {
		t.Errorf("the queried encrypted user id \n %v \n does not equal the one originally inserted \n %v \n", incident.EncryptedUserID, encryptedUserID)
	}
}

func TestGetIncidents(t *testing.T) {
	ctx := context.Background()

	incidents, ok := GetIncidents(ctx, testUserID)
	if !ok {
		t.Error("failed to get incidents")
	}

	ID := "1234567890"
	ID2 := "1234567790"

	incident1 := models.CreateIncident{
		ID:                            &ID,
		DateOfIncident:                "30/80/2021",
		ApproximatePatientAge:         "34",
		Gender:                        "female",
		LongTermPrognosis:             "dead",
		IncidentDescription:           "choking",
		Anterior:                      "someurl@url.com",
		ObjectConsistency:             "rough",
		ObjectBasicShape:              "round",
		WhatMaterialIsTheObjectMadeOf: "plastic",
		TheObjectIs:                   "small",
		LargestLength:                 "23",
		LocationOfObject:              "throat",
	}

	incident2 := models.CreateIncident{
		ID:                            &ID2,
		DateOfIncident:                "30/80/2021",
		ApproximatePatientAge:         "34",
		Gender:                        "female",
		LongTermPrognosis:             "dead",
		IncidentDescription:           "choking",
		Anterior:                      "someurl@url.com",
		ObjectConsistency:             "rough",
		ObjectBasicShape:              "round",
		WhatMaterialIsTheObjectMadeOf: "plastic",
		TheObjectIs:                   "small",
		LargestLength:                 "23",
		LocationOfObject:              "throat",
	}

	if *incidents.UserID != testUserID {
		t.Errorf("userID's from get incidents do not match: \n %v \n %v", *incidents.UserID, testUserID)
	}

	ok = reflect.DeepEqual(incidents.Incidents[0], incident1)
	if !ok {

	}

	ok = reflect.DeepEqual(incidents.Incidents[1], incident2)
	if !ok {

	}

}

func TestUpdateIncident(t *testing.T) {
	ctx := context.Background()

	ID := "1234567890"
	badID := "1234567790"

	updateWithIncident := models.UpdateIncident{
		ID:                            &ID,
		DateOfIncident:                "30/80/2021",
		ApproximatePatientAge:         "34",
		Gender:                        "female",
		LongTermPrognosis:             "dead",
		IncidentDescription:           "choking",
		Anterior:                      "someurl@url.com",
		ObjectConsistency:             "rough",
		ObjectBasicShape:              "round",
		WhatMaterialIsTheObjectMadeOf: "plastic",
		TheObjectIs:                   "small",
		LargestLength:                 "23",
		LocationOfObject:              "throat",
	}

	updateWithBadIncident := models.UpdateIncident{
		ID:                            &badID,
		DateOfIncident:                "30/80/2021",
		ApproximatePatientAge:         "34",
		Gender:                        "female",
		LongTermPrognosis:             "dead",
		IncidentDescription:           "choking",
		Anterior:                      "someurl@url.com",
		ObjectConsistency:             "rough",
		ObjectBasicShape:              "round",
		WhatMaterialIsTheObjectMadeOf: "plastic",
		TheObjectIs:                   "small",
		LargestLength:                 "23",
		LocationOfObject:              "throat",
	}

	_, ok := UpdateIncident(ctx, updateWithIncident)
	if !ok {
		t.Errorf("postgres update incident failed, updated with incident: %v", updateWithIncident)
	}

	db, conn := Open()
	defer conn.Close()

	payload := Incident{}
	err := db.First(&payload, "id = ?", updateWithIncident.ID).Error
	if err == gorm.ErrRecordNotFound {
		t.Errorf("could not find incident: %v after updating", updateWithIncident)
	} else if err != nil {
		t.Error(err)
	}

	decryptedUserID, err := decryptUserID(payload.EncryptedUserID)
	if err != nil {
		t.Error("could not decrypted encrypted userID")
	}

	if payload.ID != *updateWithIncident.ID ||
		payload.DateOfIncident != updateWithIncident.DateOfIncident ||
		payload.ApproximatePatientAge != updateWithIncident.ApproximatePatientAge ||
		payload.Gender != updateWithIncident.Gender ||
		payload.LongTermPrognosis != updateWithIncident.LongTermPrognosis ||
		payload.IncidentDescription != updateWithIncident.IncidentDescription ||
		payload.Anterior != updateWithIncident.Anterior ||
		payload.ObjectConsistency != updateWithIncident.ObjectConsistency ||
		payload.ObjectBasicShape != updateWithIncident.ObjectBasicShape ||
		payload.WhatMaterialIsTheObjectMadeOf != updateWithIncident.WhatMaterialIsTheObjectMadeOf ||
		payload.TheObjectIs != updateWithIncident.TheObjectIs ||
		payload.LargestLength != updateWithIncident.LargestLength ||
		payload.LocationOfObject != updateWithIncident.LocationOfObject ||
		decryptedUserID != testUserID {
		t.Errorf("postgres update incident failed, \n payload %v \n does not match updatewithincident \n %v", payload, render.Render(updateWithIncident))
	}

	err = db.Where(&Incident{ID: *updateWithBadIncident.ID}).Updates(updateWithBadIncident).Error
	if err == gorm.ErrRecordNotFound {
		t.Errorf("could not find incident: %v after updating", updateWithIncident)
	} else if err == nil {
		t.Error(err)
	}

}

func TestDeleteIncident(t *testing.T) {
	ctx := context.Background()

	db, conn := Open()
	defer conn.Close()

	incidentID := "1234567890"
	incidentID2 := "1234567790"
	badIncidentID := "1234567770"

	err := db.First(&Incident{}, "id = ?", badIncidentID).Delete(Incident{}).Error
	if err != gorm.ErrRecordNotFound {
		t.Errorf("expected delete incident to fail with ErrRecordNotFound with bad incident id: %v", badIncidentID)
	} else if err == nil {
		t.Errorf("expected delete incident to fail with bad incident id: %v \n but there was no error", badIncidentID)
	}

	payload, ok := DeleteIncident(ctx, incidentID)
	if !ok {
		t.Errorf("postgres delete incident failed, deleted with incidentID: %v", incidentID)
	}

	booleanTrue := true

	if *payload.ID != incidentID ||
		*payload.Deleted != booleanTrue {
		t.Errorf("postgres delete incident failed, deleted with incidentID: %v", incidentID)
	}

	payload, ok = DeleteIncident(ctx, incidentID2)
	if !ok {
		t.Errorf("postgres delete incident failed, deleted with incidentID: %v", incidentID)
	}

	if *payload.ID != incidentID2 ||
		*payload.Deleted != booleanTrue {
		t.Errorf("postgres delete incident failed, deleted with incidentID: %v", incidentID)
	}
}

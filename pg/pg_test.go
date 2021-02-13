package pg

import (
	"context"
	"fmt"
	"testing"

	"github.com/gdexlab/go-render/render"
	"github.com/gircapp/api/models"

	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	user := User{
		UserID:     "1234567890",
		Email:      "DrJimBob@jimbobclinic.com",
		Name:       "Jim Bob",
		Speciality: "Otolaryngologist",
		Degree:     "MD",
		Verified:   false,
	}
	payload, ok := CreateUser(ctx, user)
	if !ok {
		t.Errorf("postgres create user \n %v \n failed", user)
	}

	if payload.Degree != user.Degree ||
		payload.Name != user.Name ||
		payload.Email != user.Email ||
		payload.UserID != user.UserID ||
		payload.Speciality != user.Speciality {
		t.Errorf("postgres create user failed, \n payload %v \n does not match user \n %v", payload, user)
	}
}

func TestGetUser(t *testing.T) {

	user := User{
		UserID:     "1234567890",
		Email:      "DrJimBob@jimbobclinic.com",
		Name:       "Jim Bob",
		Speciality: "Otolaryngologist",
		Degree:     "MD",
		Verified:   false,
	}

	ctx := context.Background()
	userID := "1234567890"
	payload, ok := GetUser(ctx, userID)

	if !ok {
		t.Errorf("postgres get user by user id: %v failed", userID)
	}
	if *payload.Degree != user.Degree ||
		*payload.Name != user.Name ||
		*payload.Email != user.Email ||
		*payload.UserID != user.UserID ||
		*payload.Speciality != user.Speciality {
		t.Errorf("postgres get user failed, \n payload %v \n does not match user \n %v", render.Render(payload), render.Render(user))
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()

	user := User{
		UserID:     "1234567890",
		Email:      "DrJimBob@jimbobclinic.com",
		Name:       "Jim Bob",
		Degree:     "MD",
		Speciality: "Otolaryngologist",
	}

	updateWithUser := User{
		UserID:     "1234567890",
		Email:      "DrJimBobNewEmail@jimbobclinic.com",
		Name:       "Jim W. Bob",
		Degree:     "DO",
		Speciality: "Otolaryngologist",
	}

	badUser := User{
		UserID:     "1234567790",
		Email:      "DrJimBobNewEmail@jimbobclinic.com",
		Name:       "Jim W. Bob",
		Degree:     "DO",
		Speciality: "SpinDoctor",
	}

	updateWithEmptyFields := User{
		UserID:     "1234567890",
		Email:      "DrJimHWBobNewEmail@jimbobclinic.com",
		Name:       "Jim HW. Bob",
		Degree:     "",
		Speciality: "",
	}

	payload, ok := UpdateUser(ctx, updateWithUser)
	if !ok {
		t.Errorf("postgres user update failed for user: %v \n with updatewithuser: %v", user, updateWithUser)
	}

	if *payload.Degree != updateWithUser.Degree ||
		*payload.Name != updateWithUser.Name ||
		*payload.Email != updateWithUser.Email ||
		*payload.UserID != updateWithUser.UserID ||
		*payload.Speciality != updateWithUser.Speciality {
		t.Errorf("postgres update user failed, \n payload %v \n does not match updatewithuser \n %v", render.Render(payload), updateWithUser)
	}

	db := Open()
	// TODO: fix update
	err := db.First(&User{}, "user_id = ?", badUser.UserID).Omit("user_id").Updates(badUser).Error
	if err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		t.Errorf("expected postgres update user to have failed, \n by baduserID: %v", badUser.UserID)
	}

	_, ok = UpdateUser(ctx, updateWithEmptyFields)
	if !ok {
		t.Errorf("postgres user update with empty fields failed for user: %v \n with updateWithEmptyFields: %v", user, updateWithEmptyFields)
	}
	// TODO - check to see if empty fields were not updated

	payloadAfterEmptyUpdates, ok := GetUser(ctx, updateWithUser.UserID)

	if *payloadAfterEmptyUpdates.Degree != updateWithUser.Degree ||
		*payloadAfterEmptyUpdates.Name != updateWithEmptyFields.Name ||
		*payloadAfterEmptyUpdates.Email != updateWithEmptyFields.Email ||
		*payloadAfterEmptyUpdates.UserID != updateWithEmptyFields.UserID ||
		*payloadAfterEmptyUpdates.Speciality != updateWithUser.Speciality {
		t.Errorf("postgres update user with empty fields failed, \n payload %v \n does not match updateWithEmptyFields \n %v", render.Render(payloadAfterEmptyUpdates), updateWithEmptyFields)
	}

}

// TODO: write test for delete user

func TestVerifyUser(t *testing.T) {
	ctx := context.Background()

	userID := "1234567890"

	// TODO finish this

	booleanTrue := true

	verify := models.Verify{
		UserID:   &userID,
		Verified: &booleanTrue,
	}

	_, ok := VerifyUser(ctx, verify)
	if !ok {
		t.Errorf("postgres verify user failed, verified with verify: %v", render.Render(verify))
	}

	payload := User{}

	db := Open()

	err := db.First(&payload, "user_id = ?", verify.UserID).Error
	if err == gorm.ErrRecordNotFound {
		t.Errorf("could not find user: %v after verifying", verify.UserID)
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
	userID := "1234567890"
	ctx := context.Background()
	_, ok := DeleteUser(ctx, userID)
	if !ok {
		t.Error("error deleted user")
	}

	_, ok = GetUser(ctx, userID)
	if ok {
		t.Error("expected get deleted user to have failed, but it succeeded")
	}
}

func TestCreateIncident(t *testing.T) {
	ctx := context.Background()

	ID := "12345678900"

	incident := models.CreateIncident{
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

	_ = CreateIncident(ctx, incident)

}

func TestUpdateIncident(t *testing.T) {
	ctx := context.Background()

	ID := "12345678900"
	badID := "12345677900"

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

	db := Open()

	payload := Incident{}
	// TODO: fix update
	err := db.First(&payload, "id = ?", updateWithIncident.ID).Error
	if err == gorm.ErrRecordNotFound {
		t.Errorf("could not find incident: %v after updating", updateWithIncident)
	} else if err != nil {
		t.Error(err)
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
		payload.LocationOfObject != updateWithIncident.LocationOfObject {
		t.Errorf("postgres update incident failed, \n payload %v \n does not match updatewithincident \n %v", payload, render.Render(updateWithIncident))
	}

	// fmt.Println(payload.ID)
	// fmt.Println(*updateWithIncident.ID)
	// fmt.Println(payload.DateOfIncident)
	// fmt.Println(updateWithIncident.DateOfIncident)
	// fmt.Println(payload.ApproximatePatientAge)
	// fmt.Println(updateWithIncident.ApproximatePatientAge)
	// fmt.Println(payload.Gender)
	// fmt.Println(updateWithIncident.Gender)
	// fmt.Println(payload.LongTermPrognosis)
	// fmt.Println(updateWithIncident.LongTermPrognosis)
	// fmt.Println(payload.IncidentDescription)
	// fmt.Println(updateWithIncident.IncidentDescription)
	// fmt.Println(payload.Anterior)
	// fmt.Println(updateWithIncident.Anterior)
	// fmt.Println(payload.ObjectConsistency)
	// fmt.Println(updateWithIncident.ObjectConsistency)
	// fmt.Println(payload.ObjectBasicShape)
	// fmt.Println(updateWithIncident.ObjectBasicShape)
	// fmt.Println(payload.WhatMaterialIsTheObjectMadeOf)
	// fmt.Println(updateWithIncident.WhatMaterialIsTheObjectMadeOf)
	// fmt.Println(payload.TheObjectIs)
	// fmt.Println(updateWithIncident.TheObjectIs)
	// fmt.Println(payload.LargestLength)
	// fmt.Println(updateWithIncident.LargestLength)
	// fmt.Println(payload.LocationOfObject)
	// fmt.Println(updateWithIncident.LocationOfObject)

	err = db.Where(&Incident{ID: *updateWithBadIncident.ID}).Updates(updateWithBadIncident).Error
	if err == gorm.ErrRecordNotFound {
		t.Errorf("could not find incident: %v after updating", updateWithIncident)
	} else if err == nil {
		t.Error(err)
	}

}

func TestDeleteIncident(t *testing.T) {
	ctx := context.Background()

	db := Open()

	incidentID := "12345678900"
	badIncidentID := "12345677900"
	// TODO: fix delete
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

	// fmt.Println(*payload.IncidentID)
	// fmt.Println(incidentID)
	// fmt.Println(payload.Deleted)
	// fmt.Println(&booleanTrue)

}

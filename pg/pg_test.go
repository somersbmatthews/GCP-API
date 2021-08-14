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

// func TestUpdateUser(t *testing.T) {
// 	ctx := context.Background()

// 	user := User{
// 		UserID:    testUserID,
// 		Email:     "DrJimBob@jimbobclinic.com",
// 		Name:      "Jim Bob",
// 		Degree:    "MD",
// 		Specialty: "Otolaryngologist",
// 	}

// 	updateWithUser := User{
// 		UserID:    testUserID,
// 		Email:     "DrJimBobNewEmail@jimbobclinic.com",
// 		Name:      "Jim W. Bob",
// 		Degree:    "DO",
// 		Specialty: "Otolaryngologist",
// 	}

// 	badUser := User{
// 		UserID:    badUserID,
// 		Email:     "DrJimBobNewEmail@jimbobclinic.com",
// 		Name:      "Jim W. Bob",
// 		Degree:    "DO",
// 		Specialty: "SpinDoctor",
// 	}

// 	updateWithEmptyFields := User{
// 		UserID:    testUserID,
// 		Email:     "DrJimHWBobNewEmail@jimbobclinic.com",
// 		Name:      "Jim HW. Bob",
// 		Degree:    "",
// 		Specialty: "",
// 	}

// 	payload, ok := UpdateUser(ctx, updateWithUser)
// 	if !ok {
// 		t.Errorf("postgres user update failed for user: %v \n with updatewithuser: %v", user, updateWithUser)
// 	}

// 	if *payload.Degree != updateWithUser.Degree ||
// 		*payload.Name != updateWithUser.Name ||
// 		*payload.Email != updateWithUser.Email ||
// 		*payload.UserID != testUserID ||
// 		*payload.Specialty != updateWithUser.Specialty {
// 		t.Errorf("postgres update user failed, \n payload %v \n does not match updatewithuser \n %v", render.Render(payload), updateWithUser)
// 	}

// 	db, conn := Open()
// 	defer conn.Close()
// 	err := db.First(&User{}, "user_id = ?", badUserID).Omit("user_id").Updates(badUser).Error
// 	if err != gorm.ErrRecordNotFound {
// 		fmt.Println(err)
// 		t.Errorf("expected postgres update user to have failed, \n by baduserID: %v", badUser.UserID)
// 	}

// 	_, ok = UpdateUser(ctx, updateWithEmptyFields)
// 	if !ok {
// 		t.Errorf("postgres user update with empty fields failed for user: %v \n with updateWithEmptyFields: %v", user, updateWithEmptyFields)
// 	}

// 	payloadAfterEmptyUpdates, ok := GetUser(ctx, testUserID)

// 	if *payloadAfterEmptyUpdates.Degree != updateWithUser.Degree ||
// 		*payloadAfterEmptyUpdates.Name != updateWithEmptyFields.Name ||
// 		*payloadAfterEmptyUpdates.Email != updateWithEmptyFields.Email ||
// 		*payloadAfterEmptyUpdates.UserID != testUserID ||
// 		*payloadAfterEmptyUpdates.Specialty != updateWithUser.Specialty {
// 		t.Errorf("postgres update user with empty fields failed, \n payload %v \n does not match updateWithEmptyFields \n %v", render.Render(payloadAfterEmptyUpdates), updateWithEmptyFields)
// 	}

// }

// func TestVerifyUser(t *testing.T) {
// 	ctx := context.Background()

// 	booleanTrue := true

// 	verify := models.Verify{
// 		Verified: &booleanTrue,
// 	}

// 	_, ok := VerifyUser(ctx, verify, testUserID)
// 	if !ok {
// 		t.Errorf("postgres verify user failed, verified with verify: %v", render.Render(verify))
// 	}

// 	payload := User{}

// 	db, conn := Open()
// 	defer conn.Close()

// 	err := db.First(&payload, "user_id = ?", testUserID).Error
// 	if err == gorm.ErrRecordNotFound {
// 		t.Errorf("could not find user: %v after verifying", testUserID)
// 	} else if err != nil {
// 		t.Error(err)
// 	}

// 	if payload.Verified == false {
// 		t.Errorf("postgres verify user failed, \n payload %v \n does not match updatewithincident \n %v", payload, render.Render(verify))
// 	}

// 	if payload.Verified != true {
// 		t.Errorf("postgres verify user failed, \n payload %v \n does not match updatewithincident \n %v", payload, render.Render(verify))
// 	}

// }

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

var ID = "1234567890"
var ID2 = "1234567790"

var incident1 = models.Incident{
	AceticAcid:                  "no",
	AdditionalImagingAndSurgery: "no",
	Anesthesia:                  "no",
	AnteriorPhoto:               "https://www.photourl.com/anterior",
	BatteryImprintCode:          "12345",
	BatteryLocation:             "throat",
	Complications:               "death",
	CustomMagnetType:            "button",
	Dimensionality:              "3D",
	EaseOfRemoval:               "easy",
	Gender:                      "female",
	Honey:                       "yes",
	ID:                          ID,
	IncidentDescription:         "the girl swallowed a lego",
	IncidentYear:                "2021",
	LargestDepth:                "34.31",
	LargestLength:               "32.67",
	LargestWidth:                "15.23",
	LateralPhoto:                "https://www.photourl.com/lateral",
	LengthOfHospitalStay:        "two weeks",
	LifeThreatening:             "yes",
	Location:                    "United States",
	LocationOfObjects:           "stomach",
	LongTermPrognosis:           "death",
	MagnetType:                  "9V",
	MagneticPoleDirection:       "double",
	NumberOfObjects:             "1",
	NumberOfPieces:              "one",
	ObjectBasicShape:            "square",
	ObjectCharacteristics:       "round",
	ObjectConsistency:           "solid",
	ObjectMaterial:              "wood",
	ObjectsIntact:               "yes",
	Other:                       "no",
	OtherShape:                  "cylinder",
	PatientAge:                  "24",
	PosteriorPhoto:              "https://www.photourl.com/posterior",
	RemovalStrategy:             "forceps",
	SettingOfRemoval:            "examining room",
	Sucralfate:                  "no",
	SymptomSeverity:             "severe",
	SymptomsPresent:             "coughing, choking",
	TimeUntilRemoval:            "45",
	XrayOpacity:                 "opaque",
}

var incident2 = models.Incident{
	AceticAcid:                  "yes",
	AdditionalImagingAndSurgery: "yes",
	Anesthesia:                  "yes",
	AnteriorPhoto:               "https://www.photourl.com/anterior2",
	BatteryImprintCode:          "1234567",
	BatteryLocation:             "stomach",
	Complications:               "choking",
	CustomMagnetType:            "watch",
	Dimensionality:              "2D",
	EaseOfRemoval:               "hard",
	Gender:                      "male",
	Honey:                       "no",
	ID:                          ID2,
	IncidentDescription:         "the boy swallowed a lego",
	IncidentYear:                "2020",
	LargestDepth:                "34.21",
	LargestLength:               "32.47",
	LargestWidth:                "15.13",
	LateralPhoto:                "https://www.photourl.com/lateral2",
	LengthOfHospitalStay:        "three weeks",
	LifeThreatening:             "no",
	Location:                    "United States",
	LocationOfObjects:           "throat",
	LongTermPrognosis:           "good",
	MagnetType:                  "AA",
	MagneticPoleDirection:       "single",
	NumberOfObjects:             "2",
	NumberOfPieces:              "three",
	ObjectBasicShape:            "cylinder",
	ObjectCharacteristics:       "smooth",
	ObjectConsistency:           "hollow",
	ObjectMaterial:              "plastic",
	ObjectsIntact:               "no",
	Other:                       "yes",
	OtherShape:                  "cube",
	PatientAge:                  "34",
	PosteriorPhoto:              "https://www.photourl.com/posterior2",
	RemovalStrategy:             "manual",
	SettingOfRemoval:            "operating room",
	Sucralfate:                  "yes",
	SymptomSeverity:             "mild",
	SymptomsPresent:             "vomiting, coughing",
	TimeUntilRemoval:            "180",
	XrayOpacity:                 "transparent",
}

var sameIDIncident1 = models.Incident{
	AceticAcid:                  "no",
	AdditionalImagingAndSurgery: "no",
	Anesthesia:                  "no",
	AnteriorPhoto:               "https://www.photourl.com/anterior",
	BatteryImprintCode:          "12345",
	BatteryLocation:             "throat",
	Complications:               "death",
	CustomMagnetType:            "button",
	Dimensionality:              "3D",
	EaseOfRemoval:               "easy",
	Gender:                      "female",
	Honey:                       "yes",
	ID:                          ID,
	IncidentDescription:         "the girl swallowed a lego",
	IncidentYear:                "2021",
	LargestDepth:                "34.31",
	LargestLength:               "32.67",
	LargestWidth:                "15.23",
	LateralPhoto:                "https://www.photourl.com/lateral",
	LengthOfHospitalStay:        "two weeks",
	LifeThreatening:             "yes",
	Location:                    "United States",
	LocationOfObjects:           "stomach",
	LongTermPrognosis:           "death",
	MagnetType:                  "9V",
	MagneticPoleDirection:       "double",
	NumberOfObjects:             "1",
	NumberOfPieces:              "one",
	ObjectBasicShape:            "square",
	ObjectCharacteristics:       "round",
	ObjectConsistency:           "solid",
	ObjectMaterial:              "wood",
	ObjectsIntact:               "yes",
	Other:                       "no",
	OtherShape:                  "cylinder",
	PatientAge:                  "24",
	PosteriorPhoto:              "https://www.photourl.com/posterior",
	RemovalStrategy:             "forceps",
	SettingOfRemoval:            "examining room",
	Sucralfate:                  "no",
	SymptomSeverity:             "severe",
	SymptomsPresent:             "coughing, choking",
	TimeUntilRemoval:            "45",
	XrayOpacity:                 "opaque",
}

// var badIncident = models.Incident{}

func TestCreateIncidents(t *testing.T) {
	ctx := context.Background()

	incidents := []*models.Incident{&incident1, &incident2}

	_, ok := CreateIncidents(ctx, incidents, testUserID)
	if !ok {
		t.Error("couldn't create test incidents")
	}
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
	// if incident.EncryptedUserID != encryptedUserID {
	// 	t.Errorf("the queried encrypted user id \n %v \n does not equal the one originally inserted \n %v \n", incident.EncryptedUserID, encryptedUserID)
	// }
}

func TestGetIncidents(t *testing.T) {
	ctx := context.Background()

	incidents, ok := GetIncidents(ctx, testUserID)
	if !ok {
		t.Error("failed to get incidents")
	}

	if *incidents.UserID != testUserID {
		t.Errorf("userID's from get incidents do not match: \n %v \n %v", *incidents.UserID, testUserID)
	}

	ok = reflect.DeepEqual(incidents.Incidents[0], &incident2)
	if !ok {
		fmt.Println("this is incident 1")
		fmt.Println(render.Render(incident2))
		fmt.Println("this is first incident printed from database")
		fmt.Println(render.Render(incidents.Incidents[0]))
		// fmt.Println("this is incident 2 from database")
		// fmt.Println(render.Render(incident2))
		t.Error("incident 1 does not match first incident retrieved")
	}

	ok = reflect.DeepEqual(incidents.Incidents[1], &incident1)
	if !ok {
		t.Error("incident 2 does not match second incident retrieved")
	}

}

// func TestUpdateIncident(t *testing.T) {
// 	ctx := context.Background()

// 	ID := "1234567890"
// 	badID := "1234567790"

// 	updateWithIncident := models.UpdateIncident{
// 		ID:                            &ID,
// 		DateOfIncident:                "30/80/2021",
// 		ApproximatePatientAge:         "34",
// 		Gender:                        "female",
// 		LongTermPrognosis:             "dead",
// 		IncidentDescription:           "choking",
// 		Anterior:                      "someurl@url.com",
// 		ObjectConsistency:             "rough",
// 		ObjectBasicShape:              "round",
// 		WhatMaterialIsTheObjectMadeOf: "plastic",
// 		TheObjectIs:                   "small",
// 		LargestLength:                 "23",
// 		LocationOfObject:              "throat",
// 	}

// 	updateWithBadIncident := models.UpdateIncident{
// 		ID:                            &badID,
// 		DateOfIncident:                "30/80/2021",
// 		ApproximatePatientAge:         "34",
// 		Gender:                        "female",
// 		LongTermPrognosis:             "dead",
// 		IncidentDescription:           "choking",
// 		Anterior:                      "someurl@url.com",
// 		ObjectConsistency:             "rough",
// 		ObjectBasicShape:              "round",
// 		WhatMaterialIsTheObjectMadeOf: "plastic",
// 		TheObjectIs:                   "small",
// 		LargestLength:                 "23",
// 		LocationOfObject:              "throat",
// 	}

// 	_, ok := UpdateIncident(ctx, updateWithIncident)
// 	if !ok {
// 		t.Errorf("postgres update incident failed, updated with incident: %v", updateWithIncident)
// 	}

// 	db, conn := Open()
// 	defer conn.Close()

// 	payload := Incident{}
// 	err := db.First(&payload, "id = ?", updateWithIncident.ID).Error
// 	if err == gorm.ErrRecordNotFound {
// 		t.Errorf("could not find incident: %v after updating", updateWithIncident)
// 	} else if err != nil {
// 		t.Error(err)
// 	}

// 	decryptedUserID, err := decryptUserID(payload.EncryptedUserID)
// 	if err != nil {
// 		t.Error("could not decrypted encrypted userID")
// 	}

// 	if payload.ID != *updateWithIncident.ID ||
// 		payload.DateOfIncident != updateWithIncident.DateOfIncident ||
// 		payload.ApproximatePatientAge != updateWithIncident.ApproximatePatientAge ||
// 		payload.Gender != updateWithIncident.Gender ||
// 		payload.LongTermPrognosis != updateWithIncident.LongTermPrognosis ||
// 		payload.IncidentDescription != updateWithIncident.IncidentDescription ||
// 		payload.Anterior != updateWithIncident.Anterior ||
// 		payload.ObjectConsistency != updateWithIncident.ObjectConsistency ||
// 		payload.ObjectBasicShape != updateWithIncident.ObjectBasicShape ||
// 		payload.WhatMaterialIsTheObjectMadeOf != updateWithIncident.WhatMaterialIsTheObjectMadeOf ||
// 		payload.TheObjectIs != updateWithIncident.TheObjectIs ||
// 		payload.LargestLength != updateWithIncident.LargestLength ||
// 		payload.LocationOfObject != updateWithIncident.LocationOfObject ||
// 		decryptedUserID != testUserID {
// 		t.Errorf("postgres update incident failed, \n payload %v \n does not match updatewithincident \n %v", payload, render.Render(updateWithIncident))
// 	}

// 	err = db.Where(&Incident{ID: *updateWithBadIncident.ID}).Updates(updateWithBadIncident).Error
// 	if err == gorm.ErrRecordNotFound {
// 		t.Errorf("could not find incident: %v after updating", updateWithIncident)
// 	} else if err == nil {
// 		t.Error(err)
// 	}

// }

func TestDeleteIncident(t *testing.T) {
	ctx := context.Background()

	db, conn := Open()
	defer conn.Close()

	// incidentID := "1234567890"
	// incidentID2 := "1234567790"
	badIncidentID := "1234567770"

	err := db.First(&Incident{}, "id = ?", badIncidentID).Delete(Incident{}).Error
	if err != gorm.ErrRecordNotFound {
		t.Errorf("expected delete incident to fail with ErrRecordNotFound with bad incident id: %v", badIncidentID)
	} else if err == nil {
		t.Errorf("expected delete incident to fail with bad incident id: %v \n but there was no error", badIncidentID)
	}

	_, ok := DeleteIncidents(ctx, testUserID)
	if !ok {
		t.Errorf("postgres delete incidents failed for user id;, %s", testUserID)
	}

	// booleanTrue := true

	// if *payload.ID != incidentID ||
	// 	*payload.Deleted != booleanTrue {
	// 	t.Errorf("postgres delete incident failed, deleted with incidentID: %v", incidentID)
	// }

	// payload, ok = DeleteIncident(ctx, incidentID2)
	// if !ok {
	// 	t.Errorf("postgres delete incident failed, deleted with incidentID: %v", incidentID)
	// }

	// if *payload.ID != incidentID2 ||
	// 	*payload.Deleted != booleanTrue {
	// 	t.Errorf("postgres delete incident failed, deleted with incidentID: %v", incidentID)
	// }
}

package pg

import (
	"context"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gircapp/api/models"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	UserID    string `gorm:"unique; not null"`
	Email     string
	Name      string
	Specialty string
	Degree    string
	Verified  bool
}

type Incident struct {
	CreatedAt                   int64
	ID                          string `gorm:"unique; not null"`
	EncryptedUserID             string `gorm:"type:bytea"`
	Location                    string
	LocationOfObjects           string
	LongTermPrognosis           string
	SymptomsPresent             string
	AnteriorPhoto               string
	LateralPhoto                string
	PosteriorPhoto              string
	IncidentYear                string
	ObjectConsistency           string
	Gender                      string
	PatientAge                  string
	LargestLength               string
	RemovalStrategy             string
	SettingOfRemoval            string
	LengthOfHospitalStay        string
	LifeThreatening             string
	IncidentDescription         string
	TimeUntilRemoval            string
	EaseOfRemoval               string
	ObjectMaterial              string
	ObjectBasicShape            string
	Anesthesia                  string
	SymptomSeverity             string
	XrayOpacity                 string
	AceticAcid                  string
	Other                       string
	Dimensionality              string
	AdditionalImagingAndSurgery string
	NumberOfPieces              string
	ObjectsIntact               string
	ObjectCharacteristics       string
	BatteryLocation             string
	MagneticPoleDirection       string
	Complications               string
	LargestDepth                string
	Sucralfate                  string
	BatteryImprintCode          string
	OtherShape                  string
	LargestWidth                string
	MagnetType                  string
	NumberOfObjects             string
	CustomMagnetType            string
	Honey                       string
}

func CreateIncidents(ctx context.Context, incidents []*models.Incident, userID string) (*models.GoodResponse, bool) {

	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	bytea := getByteaFromString(encryptedUserID)
	for _, incident := range incidents {
		incidentModel := Incident{
			CreatedAt:                   time.Now().UnixNano(),
			ID:                          getIDFromPhotoURL(incident.AnteriorPhoto),
			EncryptedUserID:             bytea,
			Location:                    incident.Location,
			LocationOfObjects:           incident.LocationOfObjects,
			LongTermPrognosis:           incident.LongTermPrognosis,
			SymptomsPresent:             incident.SymptomsPresent,
			AnteriorPhoto:               incident.AnteriorPhoto,
			LateralPhoto:                incident.LateralPhoto,
			PosteriorPhoto:              incident.PosteriorPhoto,
			IncidentYear:                incident.IncidentYear,
			ObjectConsistency:           incident.ObjectConsistency,
			Gender:                      incident.Gender,
			PatientAge:                  incident.PatientAge,
			LargestLength:               incident.LargestLength,
			RemovalStrategy:             incident.RemovalStrategy,
			SettingOfRemoval:            incident.SettingOfRemoval,
			LengthOfHospitalStay:        incident.LengthOfHospitalStay,
			LifeThreatening:             incident.LifeThreatening,
			IncidentDescription:         incident.IncidentDescription,
			TimeUntilRemoval:            incident.TimeUntilRemoval,
			EaseOfRemoval:               incident.EaseOfRemoval,
			ObjectMaterial:              incident.ObjectMaterial,
			ObjectBasicShape:            incident.ObjectBasicShape,
			Anesthesia:                  incident.Anesthesia,
			SymptomSeverity:             incident.SymptomSeverity,
			XrayOpacity:                 incident.XrayOpacity,
			AceticAcid:                  incident.AceticAcid,
			Other:                       incident.Other,
			Dimensionality:              incident.Dimensionality,
			AdditionalImagingAndSurgery: incident.AdditionalImagingAndSurgery,
			NumberOfPieces:              incident.NumberOfPieces,
			ObjectsIntact:               incident.ObjectsIntact,
			ObjectCharacteristics:       incident.ObjectCharacteristics,
			BatteryLocation:             incident.BatteryLocation,
			MagneticPoleDirection:       incident.MagneticPoleDirection,
			Complications:               incident.Complications,
			LargestDepth:                incident.LargestDepth,
			Sucralfate:                  incident.Sucralfate,
			BatteryImprintCode:          incident.BatteryImprintCode,
			OtherShape:                  incident.OtherShape,
			LargestWidth:                incident.LargestWidth,
			MagnetType:                  incident.MagnetType,
			NumberOfObjects:             incident.NumberOfObjects,
			CustomMagnetType:            incident.CustomMagnetType,
			Honey:                       incident.Honey,
		}
		_ = db.Create(incidentModel).Error
		// if err != nil {
		// 	return nil, false
		// }
	}

	response := models.GoodResponse{
		Message: fmt.Sprintf("incidents created for user id: %s", userID),
	}
	return &response, true
}

func CreateUser(ctx context.Context, user User) (*models.CreateUserGoodResponse, bool) {

	err := db.Create(user).Error
	if err != nil {
		return nil, false
	}
	booleanTrue := true
	return &models.CreateUserGoodResponse{
		UserID:    &user.UserID,
		Email:     user.Email,
		Specialty: user.Specialty,
		Degree:    user.Degree,
		Created:   &booleanTrue,
		Name:      user.Name,
	}, true
}

func GetIncidents(ctx context.Context, userID string) (*models.GetIncidentsGoodResponse, bool) {

	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	sql := "SELECT * FROM incidents WHERE encrypted_user_id = ? ORDER BY created_at DESC"

	// bytea := stringToBin(encryptedUserID)

	bytea := getByteaFromString(encryptedUserID)

	incidents := []Incident{}

	// fields := []string{"id", "long_term_prognosis", "what_material_is_the_object_made_of", "anterior", "date_of_incident", "object_consistency", "gender", "approximate_patient_age", "location_of_object", "incident_description", "largest_length", "object_basic_shape", "the_object_is"}

	// TODO: make a sql query as per website with function for making the bytea to text
	// OR: TODO find out what that big fucking space is in my gorm query or replace with raw sql
	// FIRST: make sql query in psql to convert the bytea value to text so you can see what it is
	err = db.Raw(sql, bytea).Scan(&incidents).Error

	//	err = db.Where(Incident{}, "encrypted_user_id = ?", bytea).Find(&incidents).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	incidentResponses := []*models.Incident{}

	for _, incident := range incidents {

		incidentResponses = append(incidentResponses, &models.Incident{
			ID:                          incident.ID,
			Location:                    incident.Location,
			LocationOfObjects:           incident.LocationOfObjects,
			LongTermPrognosis:           incident.LongTermPrognosis,
			SymptomsPresent:             incident.SymptomsPresent,
			AnteriorPhoto:               incident.AnteriorPhoto,
			LateralPhoto:                incident.LateralPhoto,
			PosteriorPhoto:              incident.PosteriorPhoto,
			IncidentYear:                incident.IncidentYear,
			ObjectConsistency:           incident.ObjectConsistency,
			Gender:                      incident.Gender,
			PatientAge:                  incident.PatientAge,
			LargestLength:               incident.LargestLength,
			RemovalStrategy:             incident.RemovalStrategy,
			SettingOfRemoval:            incident.SettingOfRemoval,
			LengthOfHospitalStay:        incident.LengthOfHospitalStay,
			LifeThreatening:             incident.LifeThreatening,
			IncidentDescription:         incident.IncidentDescription,
			TimeUntilRemoval:            incident.TimeUntilRemoval,
			EaseOfRemoval:               incident.EaseOfRemoval,
			ObjectMaterial:              incident.ObjectMaterial,
			ObjectBasicShape:            incident.ObjectBasicShape,
			Anesthesia:                  incident.Anesthesia,
			SymptomSeverity:             incident.SymptomSeverity,
			XrayOpacity:                 incident.XrayOpacity,
			AceticAcid:                  incident.AceticAcid,
			Other:                       incident.Other,
			Dimensionality:              incident.Dimensionality,
			AdditionalImagingAndSurgery: incident.AdditionalImagingAndSurgery,
			NumberOfPieces:              incident.NumberOfPieces,
			ObjectsIntact:               incident.ObjectsIntact,
			ObjectCharacteristics:       incident.ObjectCharacteristics,
			BatteryLocation:             incident.BatteryLocation,
			MagneticPoleDirection:       incident.MagneticPoleDirection,
			Complications:               incident.Complications,
			LargestDepth:                incident.LargestDepth,
			Sucralfate:                  incident.Sucralfate,
			BatteryImprintCode:          incident.BatteryImprintCode,
			OtherShape:                  incident.OtherShape,
			LargestWidth:                incident.LargestWidth,
			MagnetType:                  incident.MagnetType,
			NumberOfObjects:             incident.NumberOfObjects,
			CustomMagnetType:            incident.CustomMagnetType,
			Honey:                       incident.Honey,
		})
	}

	response := &models.GetIncidentsGoodResponse{
		Incidents: incidentResponses,
		UserID:    &userID,
	}

	// log.Printf(render.Render(response))

	return response, true
}

// func UpdateIncident(ctx context.Context, incident models.UpdateIncident) (*models.UpdateIncidentGoodResponse, bool) {
// 	db, conn := Open()
// 	defer conn.Close()
// 	updateWithModel := Incident{
// 		ID:                            *incident.ID,
// 		DateOfIncident:                incident.DateOfIncident,
// 		ApproximatePatientAge:         incident.ApproximatePatientAge,
// 		Gender:                        incident.Gender,
// 		LongTermPrognosis:             incident.LongTermPrognosis,
// 		IncidentDescription:           incident.IncidentDescription,
// 		Anterior:                      incident.Anterior,
// 		ObjectConsistency:             incident.ObjectConsistency,
// 		ObjectBasicShape:              incident.ObjectBasicShape,
// 		WhatMaterialIsTheObjectMadeOf: incident.WhatMaterialIsTheObjectMadeOf,
// 		TheObjectIs:                   incident.TheObjectIs,
// 		LargestLength:                 incident.LargestLength,
// 		LocationOfObject:              incident.LocationOfObject,
// 	}
// 	err := db.First(&Incident{}, "id = ?", incident.ID).Updates(updateWithModel).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return nil, false
// 	} else if err != nil {
// 		panic(err)
// 	}
// 	booleanTrue := true
// 	return &models.UpdateIncidentGoodResponse{
// 			ID:                            incident.ID,
// 			DateOfIncident:                incident.DateOfIncident,
// 			ApproximatePatientAge:         incident.ApproximatePatientAge,
// 			Gender:                        incident.Gender,
// 			LongTermPrognosis:             incident.LongTermPrognosis,
// 			IncidentDescription:           incident.IncidentDescription,
// 			Anterior:                      incident.Anterior,
// 			ObjectConsistency:             incident.ObjectConsistency,
// 			ObjectBasicShape:              incident.ObjectBasicShape,
// 			WhatMaterialIsTheObjectMadeOf: incident.WhatMaterialIsTheObjectMadeOf,
// 			TheObjectIs:                   incident.TheObjectIs,
// 			LargestLength:                 incident.LargestLength,
// 			Updated:                       &booleanTrue,
// 		},
// 		true
// }

func DeleteIncidents(ctx context.Context, userID string) (*models.GoodResponse, bool) {

	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	bytea := getByteaFromString(encryptedUserID)

	incidents := []Incident{}
	// TODO CHANGE TO DELETE
	sql := "DELETE FROM incidents WHERE encrypted_user_id = ?"
	//err := db.Where(&Incident{}, "id = ?", incidentID).Delete(Incident{}).Error
	err = db.Raw(sql, bytea).Scan(&incidents).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	return &models.GoodResponse{
			Message: fmt.Sprintf("all incidents deleted for user id: %s", userID),
		},
		true
}

// func GetUser(ctx context.Context, userId string) (*models.GetUserGoodResponse, bool) {

// 	model := User{}
// 	err := db.First(&model, "user_id = ?", userId).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return &models.GetUserGoodResponse{}, false
// 	} else if err != nil {
// 		panic(err)
// 	}

// 	return &models.GetUserGoodResponse{
// 			Name:      &model.Name,
// 			Degree:    &model.Degree,
// 			Verified:  &model.Verified,
// 			Email:     &model.Email,
// 			Specialty: &model.Specialty,
// 			UserID:    &model.UserID,
// 		},
// 		true
// }

// func UpdateUser(ctx context.Context, user User) (*models.UpdateUserGoodResponse, bool) {
// 	db, conn := Open()
// 	defer conn.Close()
// 	model := user
// 	err := db.First(&User{}, "user_id = ?", user.UserID).Updates(user).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return nil, false
// 	} else if err != nil {
// 		panic(err)
// 	}
// 	booleanTrue := true
// 	return &models.UpdateUserGoodResponse{
// 			UserID:    &model.UserID,
// 			Name:      &model.Name,
// 			Email:     &model.Email,
// 			Degree:    &model.Degree,
// 			Specialty: &model.Specialty,
// 			Updated:   &booleanTrue,
// 			Verified:  model.Verified,
// 		},
// 		true
// }

func DeleteUser(ctx context.Context, userID string) (*models.DeleteUserGoodResponse, bool) {

	err := db.First(&User{}, "user_id = ?", userID).Delete(User{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	booleanTrue := true

	return &models.DeleteUserGoodResponse{
			UserID:  userID,
			Deleted: &booleanTrue,
		},
		true
}

// func VerifyUser(ctx context.Context, verify models.Verify, userID string) (*models.UpdateUserGoodResponse, bool) {
// 	db, conn := Open()
// 	defer conn.Close()
// 	model := User{}
// 	err := db.First(&model, "user_id = ?", userID).Update("verified", verify.Verified).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return nil, false
// 	} else if err != nil {
// 		panic(err)
// 	}
// 	booleanTrue := true
// 	return &models.UpdateUserGoodResponse{
// 			UserID:    &userID,
// 			Name:      &model.Name,
// 			Email:     &model.Email,
// 			Degree:    &model.Degree,
// 			Specialty: &model.Specialty,
// 			Updated:   &booleanTrue,
// 			Verified:  booleanTrue,
// 		},
// 		true
// }

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func comparePasswords(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		panic(err)
	}
	return true
}

// func stringToBin(s string) (binString string) {
// 	for _, c := range s {
// 		binString = fmt.Sprintf("%s%b", binString, c)
// 	}
// 	return
// }

func getByteaFromString(encryptedUserID string) string {

	bytes := []byte(encryptedUserID)

	hexStr := hex.EncodeToString(bytes)

	hexStrWithBackslash := "\\x" + hexStr

	return hexStrWithBackslash

}

func getStringFromBytea(bytea string) string {

	// hexStr = strings.TrimPrefix(bytea, "\\x")

	bytes, err := hex.DecodeString(bytea)
	if err != nil {
		panic(err)
	}

	str := string(bytes[:])

	return str
}

func getIDFromPhotoURL(url string) string {
	const regex = `[^\/]+$`
	var unformattedString = `\/v0\/b\/chd-backend.appspot.com\/o\/image-uploads\/D7FE3971-1270-4894-979E-D28D085B69E8.jpg`
	re := regexp.MustCompile(regex)
	formattedStringSlice := re.FindStringSubmatch(unformattedString)
	//	fmt.Println(formattedStringSlice)
	formattedString := strings.Join(formattedStringSlice, "")
	// fmt.Println(formattedString)
	withoutjpeg := strings.TrimSuffix(formattedString, ".jpg")
	// fmt.Println(withoutjpeg)
	return withoutjpeg
}

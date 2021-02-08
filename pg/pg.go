package pg

import (
	"context"

	"github.com/somersbmatthews/gircapp2/models"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	UserId     string
	Email      string
	Name       string
	Speciality string
	Degree     string
	Verified   bool
}

type Incident struct {
	ID                                  string
	Long_term_prognosis                 string
	What_material_is_the_object_made_of string
	Anterior                            string
	Date_of_Incident                    string
	Object_Consistency                  string
	Gender                              string
	Approximate_Patient_Age             string
	Location_of_object                  string
	Incident_Description                string
	Largest_Length                      string
	Object_Basic_Shape                  string
	The_object_is                       string
}

func init() {
	db := open()
	err := db.AutoMigrate(&User{}, &Incident{}).Error
	if err != nil {
		panic(err)
	}
}

func open() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=gorm password=gorm DB.name=postgres port=5432 sslmode=disable",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func CreateIncident(ctx context.Context, incident models.CreateIncident) *models.CreateIncidentGoodResponse {
	db := open()

	incidentModel := Incident{
		ID:                                  *incident.ID,
		Date_of_Incident:                    incident.DateOfIncident,
		Approximate_Patient_Age:             incident.ApproximatePatientAge,
		Gender:                              incident.Gender,
		Long_term_prognosis:                 incident.LongTermPrognosis,
		Incident_Description:                incident.IncidentDescription,
		Anterior:                            incident.Anterior,
		Object_Consistency:                  incident.ObjectConsistency,
		Object_Basic_Shape:                  incident.ObjectBasicShape,
		What_material_is_the_object_made_of: incident.WhatMaterialIsTheObjectMadeOf,
		The_object_is:                       incident.TheObjectIs,
		Largest_Length:                      incident.LargestLength,
		Location_of_object:                  incident.LocationOfObject,
	}

	err := db.Create(incidentModel)
	if err != nil {
		panic(err)
	}

	response := models.CreateIncidentGoodResponse{

		ID:                            incident.ID,
		DateOfIncident:                incident.DateOfIncident,
		ApproximatePatientAge:         incident.ApproximatePatientAge,
		Gender:                        incident.Gender,
		LongTermPrognosis:             incident.LongTermPrognosis,
		IncidentDescription:           incident.IncidentDescription,
		Anterior:                      incident.Anterior,
		ObjectConsistency:             incident.ObjectConsistency,
		ObjectBasicShape:              incident.ObjectBasicShape,
		WhatMaterialIsTheObjectMadeOf: incident.WhatMaterialIsTheObjectMadeOf,
		TheObjectIs:                   incident.TheObjectIs,
		LargestLength:                 incident.LargestLength,

		Created: true,
	}

	return &response
}

func CreateUser(ctx context.Context, user User) (*models.CreateUserGoodResponse, bool) {
	db := open()

	err := db.Create(user).Error
	if err != nil {
		panic(err)
	}

	booleanTrue := true

	return &models.CreateUserGoodResponse{
		UserID:     user.UserId,
		Email:      user.Email,
		Speciality: user.Speciality,
		Degree:     user.Degree,
		Created:    &booleanTrue,
		Name:       user.Name,
	}, true
}

// func GetIncidents(ctx context.Context, userId string) (*models.GetIncidentsGoodResponse, bool) {
// 	db := open()

// 	err := db.Model(Incident{}).Where("userId = ?", userId).Error

// }

func UpdateIncident(ctx context.Context, incident models.UpdateIncident) (*models.UpdateIncidentGoodResponse, bool) {
	db := open()

	// TODO: fix swagger.yaml so update incident and all other types below that do not have allof

	model := Incident{}

	err := db.Model(&model).Where("ID = ?", incident.ID).Updates(incident).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	}

	booleanTrue := true

	return &models.UpdateIncidentGoodResponse{
			ID:                            incident.ID,
			DateOfIncident:                incident.DateOfIncident,
			ApproximatePatientAge:         incident.ApproximatePatientAge,
			Gender:                        incident.Gender,
			LongTermPrognosis:             incident.LongTermPrognosis,
			IncidentDescription:           incident.IncidentDescription,
			Anterior:                      incident.Anterior,
			ObjectConsistency:             incident.ObjectConsistency,
			ObjectBasicShape:              incident.ObjectBasicShape,
			WhatMaterialIsTheObjectMadeOf: incident.WhatMaterialIsTheObjectMadeOf,
			TheObjectIs:                   incident.TheObjectIs,
			LargestLength:                 incident.LargestLength,
			Updated:                       &booleanTrue,
		},
		true
}

func DeleteIncident(ctx context.Context, incidentID string) (*models.DeleteIncidentGoodResponse, bool) {
	db := open()

	model := Incident{}

	err := db.Where("ID = ?", incidentID).Delete(model).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	}

	booleanTrue := true

	return &models.DeleteIncidentGoodResponse{
			Deleted:    &booleanTrue,
			IncidentID: &incidentID,
		},
		true
}

func GetUser(ctx context.Context, userId string) (*models.GetUserGoodResponse, bool) {
	db := open()

	model := User{}

	err := db.Where(&User{UserId: userId}).First(model).Error
	if err == gorm.ErrRecordNotFound {
		return &models.GetUserGoodResponse{}, false
	}

	booleanTrue := true

	return &models.GetUserGoodResponse{
			Name:       &model.Name,
			Degree:     &model.Degree,
			Verified:   &booleanTrue,
			Email:      &model.Email,
			Speciality: &model.Speciality,
			UserID:     &model.UserId,
		},
		true
}

func UpdateUser(ctx context.Context, user User) (*models.UpdateUserGoodResponse, bool) {
	db := open()

	model := User{}

	err := db.Model(&model).Where("userId = ?", user.UserId).Updates(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	}

	booleanTrue := true

	return &models.UpdateUserGoodResponse{
			Name:       &model.Name,
			Email:      &model.Email,
			Degree:     &model.Degree,
			Speciality: &model.Speciality,
			Updated:    &booleanTrue,
			Verified:   model.Verified,
		},
		true

}

func VerifyUser(ctx context.Context, verify models.Verify) (*models.UpdateUserGoodResponse, bool) {
	db := open()

	model := User{}

	err := db.Where("userId = ?", verify.UserID).Update("verified", verify.Verified).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	}

	booleanTrue := true

	return &models.UpdateUserGoodResponse{
			Name:       &model.Name,
			Email:      &model.Email,
			Degree:     &model.Degree,
			Speciality: &model.Speciality,
			Updated:    &booleanTrue,
			Verified:   booleanTrue,
		},
		true
}

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
		return false
	}
	return true
}

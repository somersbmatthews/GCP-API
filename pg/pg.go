package pg

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/gircapp/api/models"

	"golang.org/x/crypto/bcrypt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// TESTING: uncomment cloud-sql proxy postgres
	// 	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
)

type User struct {
	UserID     string
	Email      string
	Name       string
	Speciality string
	Degree     string
	Verified   bool
}

type Incident struct {
	ID                            string
	LongTermPrognosis             string
	WhatMaterialIsTheObjectMadeOf string
	Anterior                      string
	DateOfIncident                string
	ObjectConsistency             string
	Gender                        string
	ApproximatePatientAge         string
	LocationOfObject              string
	IncidentDescription           string
	LargestLength                 string
	ObjectBasicShape              string
	TheObjectIs                   string
}

var postgrespassword string

func init() {
	db := Open()

	_ = db.AutoMigrate(&User{}, &Incident{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// TESTING: uncomment access postgres passsword

	// password, err := accessPostgresPassword()
	// if err != nil {
	// 	panic(err)
	// }

	// postgrespassword = password
}

func accessPostgresPassword() (string, error) {

	name := "projects/gircapp/secrets/POSTGRESPASSWORD/versions/latest"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}

	return string(result.Payload.Data), nil
}

func Open() *gorm.DB {

	// DSN := "host=project:region:instance user=postgres dbname=postgres password=%s sslmode=disable"
	DSN := "host=localhost user=gorm password=gorm database=postgres port=5432 sslmode=disable"
	// TESTING: uncomment driver name
	db, err := gorm.Open(postgres.New(postgres.Config{
		// DriverName: "cloudsqlpostgres",
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CreateIncident(ctx context.Context, incident models.CreateIncident) *models.CreateIncidentGoodResponse {
	db := Open()

	incidentModel := Incident{
		ID:                            *incident.ID,
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
		LocationOfObject:              incident.LocationOfObject,
	}

	err := db.Create(incidentModel).Error
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
		Created:                       true,
	}

	return &response
}

func CreateUser(ctx context.Context, user User) (*models.CreateUserGoodResponse, bool) {
	db := Open()

	err := db.Create(user).Error
	if err != nil {
		panic(err)
	}

	booleanTrue := true

	return &models.CreateUserGoodResponse{
		UserID:     user.UserID,
		Email:      user.Email,
		Speciality: user.Speciality,
		Degree:     user.Degree,
		Created:    &booleanTrue,
		Name:       user.Name,
	}, true
}

// func GetIncidents(ctx context.Context, userId string) (*models.GetIncidentsGoodResponse, bool) {
// 	db := Open()
//

// 	err := db.Model(Incident{}).Where("userId = ?", userId).Error

// }

func UpdateIncident(ctx context.Context, incident models.UpdateIncident) (*models.UpdateIncidentGoodResponse, bool) {
	db := Open()

	updateWithModel := Incident{
		ID:                            *incident.ID,
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
		LocationOfObject:              incident.LocationOfObject,
	}
	// TODO: fix update
	err := db.First(&Incident{}, "id = ?", incident.ID).Updates(updateWithModel).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
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
	db := Open()

	err := db.First(&Incident{}, "id = ?", incidentID).Delete(Incident{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	booleanTrue := true

	return &models.DeleteIncidentGoodResponse{
			Deleted:    &booleanTrue,
			IncidentID: &incidentID,
		},
		true
}

func GetUser(ctx context.Context, userId string) (*models.GetUserGoodResponse, bool) {
	db := Open()

	model := User{}

	err := db.First(&model, "user_id = ?", userId).Error

	if err == gorm.ErrRecordNotFound {
		return &models.GetUserGoodResponse{}, false
	} else if err != nil {
		panic(err)
		//	return nil, false
	}

	booleanTrue := true

	return &models.GetUserGoodResponse{
			Name:       &model.Name,
			Degree:     &model.Degree,
			Verified:   &booleanTrue,
			Email:      &model.Email,
			Speciality: &model.Speciality,
			UserID:     &model.UserID,
		},
		true
}

func UpdateUser(ctx context.Context, user User) (*models.UpdateUserGoodResponse, bool) {
	db := Open()

	model := user

	// TODO: change all db.Updates to use db.model instead of db.Where, so no db.where because that does not return an error
	// so I am going to have to get old and new value

	// aslo and idea: db.First(&user, "id = ?", "string_primary_key")

	// err := db.Model(&User{}).Where(&User{UserID: user.UserID}).Omit("user_id").Updates(user).Error
	err := db.First(&User{}, "user_id = ?", user.UserID).Omit("user_id").Updates(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
		//	return nil, false
	}

	booleanTrue := true

	return &models.UpdateUserGoodResponse{
			UserID:     &model.UserID,
			Name:       &model.Name,
			Email:      &model.Email,
			Degree:     &model.Degree,
			Speciality: &model.Speciality,
			Updated:    &booleanTrue,
			Verified:   model.Verified,
		},
		true

}

func DeleteUser(ctx context.Context, userID string) (*models.DeleteUserGoodResponse, bool) {
	db := Open()

	err := db.First(&User{}, "user_id = ?", userID).Delete(User{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	booleanTrue := true

	return &models.DeleteUserGoodResponse{
			Deleted: &booleanTrue,
			UserID:  &userID,
		},
		true
}

func VerifyUser(ctx context.Context, verify models.Verify) (*models.UpdateUserGoodResponse, bool) {
	db := Open()

	model := User{}

	err := db.First(&User{}, "user_id = ?", verify.UserID).Update("verified", verify.Verified).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
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
	}
	return true
}

package pg

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/gdexlab/go-render/render"
	"github.com/gircapp/api/models"
	"golang.org/x/crypto/bcrypt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	// LOCAL TESTING: uncomment cloud-sql proxy postgres
	// _ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	UserID     string `gorm: "unique"`
	Email      string
	Name       string
	Speciality string
	Degree     string
	Verified   bool
}

type Incident struct {
	CreatedAt                     int64
	ID                            string `gorm: "unique"`
	EncryptedUserID               string `gorm:"type:bytea"`
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

// type IncidentQuery struct {
// 	ID                            string `gorm: "unique"`
// 	LongTermPrognosis             string
// 	WhatMaterialIsTheObjectMadeOf string
// 	Anterior                      string
// 	DateOfIncident                string
// 	ObjectConsistency             string
// 	Gender                        string
// 	ApproximatePatientAge         string
// 	LocationOfObject              string
// 	IncidentDescription           string
// 	LargestLength                 string
// 	ObjectBasicShape              string
// 	TheObjectIs                   string
// }

var postgrespassword string

func init() {
	password, err := accessPostgresPassword()
	if err != nil {
		panic(err)
	}
	postgrespassword = password
	db := Open()
	ok := db.Migrator().HasTable(&User{})
	if !ok {
		_ = db.AutoMigrate(&User{})
	}
	ok = db.Migrator().HasTable(&Incident{})
	if !ok {
		_ = db.AutoMigrate(&Incident{})
	}
}

func accessPostgresPassword() (string, error) {
	name := "projects/gircapp/secrets/POSTGRESPASSWORD/versions/latest"
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}
	return string(result.Payload.Data), nil
}

// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of SQL Server.
func initSocketConnectionPool() (*sql.DB, error) {
	var (
		dbUser                 = "gorm"
		instanceConnectionName = "gircapp:us-central1:gircapppostgres"
		dbName                 = "postgres"
	)
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}
	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, postgrespassword, dbName, socketDir, instanceConnectionName)
	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	configureConnectionPool(dbPool)
	return dbPool, nil
}

// 10.88.176.3
// gircapp:us-central1:gircapppostgres

func initTCPConnectionPool() (*sql.DB, error) {
	var (
		dbUser    = "gorm"
		dbTcpHost = "10.88.176.3"
		dbPort    = "5432"
		dbName    = "postgres"
	)
	var dbURI string
	dbURI = fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTcpHost, dbUser, postgrespassword, dbPort, dbName)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	configureConnectionPool(dbPool)
	return dbPool, nil
}

func configureConnectionPool(dbPool *sql.DB) {
	dbPool.SetMaxIdleConns(5)
	dbPool.SetMaxOpenConns(7)
	dbPool.SetConnMaxLifetime(1800)
}

func Open() *gorm.DB {
	sqlDB, err := initTCPConnectionPool()
	if err != nil {
		errMsg := fmt.Sprintf("%v,::: %v", err, render.Render(sqlDB))
		panic(errMsg)
	}
	// dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTcpHost, dbUser, postgrespassword, dbPort, dbName)
	// dbPool, err := sql.Open("pgx", dbURI)

	// DSN := "host=localhost user=gorm password=gorm database=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		// DSN: DSN,
		Conn: sqlDB,
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		// QueryFields: true,
	})
	if err != nil {
		errMsg := fmt.Sprintf("%v,::: %v", err, render.Render(db))
		panic(errMsg)
	}
	return db
}

func CreateIncident(ctx context.Context, incident models.CreateIncident, userID string) *models.CreateIncidentGoodResponse {
	db := Open()
	// conn, err := db.DB()
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()
	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	// bytea := stringToBin(encryptedUserID)

	bytea := getByteaFromString(encryptedUserID)

	incidentModel := Incident{
		CreatedAt:                     time.Now().UnixNano(),
		ID:                            *incident.ID,
		EncryptedUserID:               bytea,
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
	}
	err = db.Create(incidentModel).Error
	if err != nil {
		panic(err)
	}
	response := models.CreateIncidentGoodResponse{
		UserID:                        userID,
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

func CreateUser(ctx context.Context, user User, userID string) (*models.CreateUserGoodResponse, bool) {
	db := Open()
	user.UserID = userID
	err := db.Create(user).Error
	if err != nil {
		return nil, false
	}
	booleanTrue := true
	return &models.CreateUserGoodResponse{
		UserID:     &userID,
		Email:      user.Email,
		Speciality: user.Speciality,
		Degree:     user.Degree,
		Created:    &booleanTrue,
		Name:       user.Name,
	}, true
}

func GetIncidents(ctx context.Context, userID string) (*models.GetIncidentsGoodResponse, bool) {
	db := Open()

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

	// err = db.Where(Incident{}, "encrypted_user_id = ?", bytea).Find(&incidents).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	incidentResponses := []*models.GetIncidentsIncident{}

	for _, incident := range incidents {
		incidentResponses = append(incidentResponses, &models.GetIncidentsIncident{
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
			LocationOfObject:              incident.LocationOfObject,
		})
	}

	response := &models.GetIncidentsGoodResponse{
		Incidents: incidentResponses,
		UserID:    &userID,
	}

	// log.Printf(render.Render(response))

	return response, true
}

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
	err := db.First(&Incident{}, "id = ?", incident.ID).Updates(updateWithModel).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
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
			Deleted: &booleanTrue,
			ID:      &incidentID,
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
	}

	return &models.GetUserGoodResponse{
			Name:       &model.Name,
			Degree:     &model.Degree,
			Verified:   &model.Verified,
			Email:      &model.Email,
			Speciality: &model.Speciality,
			UserID:     &model.UserID,
		},
		true
}

func UpdateUser(ctx context.Context, user User, userID string) (*models.UpdateUserGoodResponse, bool) {
	db := Open()
	model := user
	err := db.First(&User{}, "user_id = ?", userID).Omit("user_id").Updates(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	booleanTrue := true
	return &models.UpdateUserGoodResponse{
			UserID:     &userID,
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
			UserID:  userID,
			Deleted: &booleanTrue,
		},
		true
}

func VerifyUser(ctx context.Context, verify models.Verify, userID string) (*models.UpdateUserGoodResponse, bool) {
	db := Open()
	model := User{}
	err := db.First(&model, "user_id = ?", userID).Update("verified", verify.Verified).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	booleanTrue := true
	return &models.UpdateUserGoodResponse{
			UserID:     &userID,
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

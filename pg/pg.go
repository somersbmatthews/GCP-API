package pg

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/gdexlab/go-render/render"
	"github.com/gircapp/api/models"

	// "github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	// TESTING: uncomment cloud-sql proxy postgres
	// _ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	password, err := accessPostgresPassword()
	if err != nil {
		panic(err)
	}

	postgrespassword = password

	// var pgdriver postgres.Hstore
	// fmt.Println(pgdriver)

	// db := Open()

	// _ = db.AutoMigrate(&User{}, &Incident{})

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// TESTING: uncomment access postgres passsword

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

// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of SQL Server.
func initSocketConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_postgres_databasesql_create_socket]
	var (
		dbUser = "gorm" // e.g. 'my-db-user'
		// dbPwd                  = postgrespassword                      // e.g. 'my-db-password'
		instanceConnectionName = "gircapp:us-central1:gircapppostgres" // e.g. 'project:region:instance'
		dbName                 = "postgres"                            // e.g. 'my-database'
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

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, nil
	// [END cloud_sql_postgres_databasesql_create_socket]
}

// 10.88.176.3
// gircapp:us-central1:gircapppostgres

func initTCPConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_postgres_databasesql_create_tcp]
	var (
		dbUser    = "gorm"           // e.g. 'my-db-user'
		dbPwd     = postgrespassword // e.g. 'my-db-password'
		dbTcpHost = "10.88.176.3"    // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
		dbPort    = "5432"           // e.g. '5432'
		dbName    = "postgres"       // e.g. 'my-database'
	)

	var dbURI string
	dbURI = fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTcpHost, dbUser, dbPwd, dbPort, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, nil
	// [END cloud_sql_postgres_databasesql_create_tcp]
}

func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_postgres_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)

	// [END cloud_sql_postgres_databasesql_limit]

	// [START cloud_sql_postgres_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)

	// [END cloud_sql_postgres_databasesql_lifetime]
}

func Open() *gorm.DB {

	var (
		dbUser = "postgres" // e.g. 'my-db-user'
		// dbPwd     = postgrespassword // e.g. 'my-db-password'
		dbTcpHost = "10.88.176.3" // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
		dbPort    = "5432"        // e.g. '5432'
		dbName    = "postgres"    // e.g. 'my-database'
	)

	// sqlDB, err := initTCPConnectionPool()
	// sqlDB, err := initTCPConnectionPool()
	// if err != nil {
	// 	errMsg := fmt.Sprintf("%v,::: %v", err, render.Render(sqlDB))
	// 	panic(errMsg)
	// }
	// gormDB, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTcpHost, dbUser, postgrespassword, dbPort, dbName)
	// dbPool, err := sql.Open("pgx", dbURI)

	// DSN := fmt.Sprintf("host=10.88.176.3 user=postgres password=%s dbname=postgres port:5432 sslmode=disable", postgrespassword)
	// DSN := "host=localhost user=gorm password=gorm database=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbURI,
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	// db, err := gorm.Open("postgres", "host=%s port=%s user=%s dbname=%s password=%s", dbTcpHost, dbPort, dbUser, dbName, postgrespassword)

	if err != nil {
		errMsg := fmt.Sprintf("%v,::: %v", err, render.Render(db))
		panic(errMsg)
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
		//	return nil, false
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

	err := db.First(&model, "user_id = ?", verify.UserID).Update("verified", verify.Verified).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	booleanTrue := true

	return &models.UpdateUserGoodResponse{
			UserID:     &model.UserID,
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

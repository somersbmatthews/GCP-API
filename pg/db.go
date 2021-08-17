package pg

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	// LOCAL TESTING: uncomment cloud-sql proxy postgres
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/gdexlab/go-render/render"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgrespassword string

var db *gorm.DB
var conn *sql.DB

func init() {
	password, err := accessPostgresPassword()
	if err != nil {
		panic(err)
	}
	postgrespassword = password
	db, conn = Open()
	// TODO: add conn.Close for all other db := Open()
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
		// dbTcpHost = "127.0.0.1"
		dbPort = "5432"
		dbName = "postgres"
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

func Open() (*gorm.DB, *sql.DB) {
	//sqlDB, err := initSocketConnectionPool()
	sqlDB, err := initTCPConnectionPool()
	if err != nil {
		errMsg := fmt.Sprintf("%v,::: %v", err, render.Render(sqlDB))
		panic(errMsg)
	}
	// dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTcpHost, dbUser, postgrespassword, dbPort, dbName)
	// dbPool, err := sql.Open("pgx", dbURI)

	// DSN := "host=localhost user=gorm password=gorm database=postgres port=5432 sslmode=disable"
	postgresDB, err := gorm.Open(postgres.New(postgres.Config{
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
	conn, err := db.DB()
	if err != nil {
		panic(err)
	}
	db = postgresDB

	return db, conn
}

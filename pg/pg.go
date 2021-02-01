package pg

import (
	"errors"

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
	IncidentId      string
	EncryptedUserId string
	ChokingObject   ChokingObject `gorm:"foreignkey:IncidentId"`
}

type ChokingObject struct {
	ChokingObjectId string
	IncidentId      string
	Images          []Images `gorm:"foreignkey:ChokingObjectId"`
}

type Images struct {
	gorm.Model
	ChokingObjectId string
	Url             string
}

func init() {
	db := open()
	err := db.AutoMigrate(&User{}, &Incident{}, &ChokingObject{}, &Images{}).Error
	if err != nil {
		panic(err)
	}
}
func open() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=gorm password=gorm DB.name=postgres port=9920 sslmode=disable", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func CreateUser(user models.User, token string) models.RegisterUserResponse {
	db := open()

	model := User{
		UserId:     *user.UserID,
		Email:      user.Email,
		Name:       *user.Name,
		Speciality: *user.Speciality,
		Degree:     *user.Degree,
		Verified:   false,
	}

	if user.Password != "" {
		if user.Email == "" {
			err := errors.New("must have both email and password")
			panic(err)
		}
	}

	err := db.Create(model).Error
	if err != nil {
		panic(err)
	}
}

func CreateAndVerifyUser(user models.User) models.RegisterAndVerifyUserResponse {
	db := open()

	model := User{
		UserId:     *user.UserID,
		Email:      user.Email,
		Name:       *user.Name,
		Speciality: *user.Speciality,
		Degree:     *user.Degree,
		Verified:   true,
	}

	if user.Password != "" {
		if user.Email == "" {
			err := errors.New("must have both email and password")
			panic(err)
		}

	}

	err := db.Create(model).Error
	if err != nil {
		panic(err)
	}

}

func GetUser(userId string) (models.GetUserGoodResponse, bool) {
	db := open()

	model := User{}

	err := db.Where(&User{UserId: userId}).First(model).Error
	if err == gorm.ErrRecordNotFound {
		return models.GetUserGoodResponse{}, false
	}

	return models.GetUserGoodResponse{
			Name:       &model.Name,
			Degree:     &model.Degree,
			Verified:   model.Verified,
			Email:      &model.Email,
			Speciality: &model.Speciality,
			UserID:     &model.UserId,
		},
		true
}

func FindUserWithEmailAndPassword(email string, password string) (models.LoginGoodResponse, string, bool) {
	db := open()

	model := User{


	err := db.Where(
}

// func UpdateUser

// func CreateIncident

// func GetIncidents

// func UpdateIncident

// func DeleteIncident

func hashAndSalt(pwd string) string {
	bytePwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashedPassword := string(hash)
	ok := comparePasswords(hashedPassword, pwd)
	if !ok {
		panic(err)
	}

}

func comparePasswords(hashedPassword string, plainPassword string) bool {
	bytePassword := []byte(plainPassword)
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return false
	}
	return true
}

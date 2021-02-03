package pg

import (
	"context"

	"github.com/somersbmatthews/gircapp2/models"

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
		DSN:                  "host=localhost user=gorm password=gorm DB.name=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func CreateUser(ctx context.Context, user User) *models.CreateUserGoodResponse {
	db := open()

	err := db.Create(user).Error
	if err != nil {
		panic(err)
	}

	booleanFalse := true

	response := models.CreateUserGoodResponse{
		UserID:     user.UserId,
		Email:      user.Email,
		Speciality: user.Speciality,
		Degree:     user.Degree,
		Created:    &booleanFalse,
		Name:       user.Name,
	}
	return &response
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

	err := db.Model(&model).Where("userId = ?", user.UserId).Select("*").Updates(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	}

	err = db.Where("userId = ?", user.UserId).First(&model).Error

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

	err = db.Where("userId = ?", verify.UserID).First(&model).Error
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

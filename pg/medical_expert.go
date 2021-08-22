package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/gircapp/api/models"
	"gorm.io/gorm"
)

type Expert struct {
	ID                                           string `gorm:"unique; not null"`
	Email                                        string
	Name                                         string
	Expertise                                    string
	Degree                                       string
	DirectorVerified                             bool
	Registered                                   bool
	EmailConfirmed                               bool // perhaps remove depending on email confirmation strategy
	DeviceType                                   string
	FCMToken                                     string
	RegistrationAttempts                         int
	LastRegisterationAttemptTimeAfterMaxAttempts int64
	LastOpeningAppTime                           int64
	RegistrationTime                             int64
	Banned                                       bool
}

func CreateExpertWithAutoDirectorAndEmailVerification(ctx context.Context, expertRequestObject *models.Expert, userID string) (*models.GoodResponse, bool) {

	expert := Expert{
		ID:               userID,
		Email:            *expertRequestObject.Email,
		Name:             *expertRequestObject.Name,
		Expertise:        *expertRequestObject.Expertise,
		Degree:           *expertRequestObject.Degree,
		DirectorVerified: true,
		Registered:       true,
		EmailConfirmed:   true,
		DeviceType:       *expertRequestObject.DeviceType,
		Banned:           false,
		RegistrationTime: time.Now().Unix(),
	}

	err := db.Create(expert).Error
	if err != nil {
		return nil, false
	}
	message := fmt.Sprintf("Medical Expert with UserID: %s", userID)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func CreateExpertWithAutoDirectorWithoutEmailVerification(ctx context.Context, expertRequestObject *models.Expert, userID string) (*models.GoodResponse, bool) {

	expert := Expert{
		ID:               userID,
		Email:            *expertRequestObject.Email,
		Name:             *expertRequestObject.Name,
		Expertise:        *expertRequestObject.Expertise,
		Degree:           *expertRequestObject.Degree,
		DirectorVerified: true,
		Registered:       true,
		EmailConfirmed:   false,
		DeviceType:       *expertRequestObject.DeviceType,
		Banned:           false,
		RegistrationTime: time.Now().Unix(),
	}

	err := db.Create(expert).Error
	if err != nil {
		return nil, false
	}
	message := fmt.Sprintf("Medical Expert with UserID: %s", userID)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func CreateExpertNormally(ctx context.Context, expertRequestObject *models.Expert, userID string) (*models.GoodResponse, bool) {

	expert := Expert{
		ID:               userID,
		Email:            *expertRequestObject.Email,
		Name:             *expertRequestObject.Name,
		Expertise:        *expertRequestObject.Expertise,
		Degree:           *expertRequestObject.Degree,
		DirectorVerified: false,
		Registered:       true,
		EmailConfirmed:   false,
		DeviceType:       *expertRequestObject.DeviceType,
		Banned:           false,
		RegistrationTime: time.Now().Unix(),
	}

	err := db.Create(expert).Error
	if err != nil {
		return nil, false
	}
	message := fmt.Sprintf("Medical Expert with UserID: %s", userID)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func ConfirmEmail(ctx context.Context, email string, userID string) (*models.GoodResponse, bool) {
	model := Expert{
		ID: userID,
	}

	err := db.Model(model).Update("email", email)
	if err != nil {
		return nil, false
	}
	message := fmt.Sprintf("Medical Expert with UserID: %s, has confirmed email with %s", userID, email)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func GetMedicalExpert(ctx context.Context, userID string) (*models.GetExpertResponse, bool) {
	model := Expert{}
	err := db.First(&model, "id = ?", userID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	return &models.GetExpertResponse{
		ID:             &model.ID,
		Name:           &model.Name,
		Degree:         &model.Degree,
		Expertise:      &model.Expertise,
		Email:          &model.Email,
		Verified:       &model.DirectorVerified,
		EmailConfirmed: &model.EmailConfirmed,
		Banned:         &model.Banned,
	}, true

}

func UpdateMedicalExpert(ctx context.Context, expertRequestObject models.Expert, userId string) (*models.GoodResponse, bool) {
	expert := Expert{
		Name: *expertRequestObject.Name,
		//Email: *expertRequestObject.Email, // email is updated via emailed link with jwt
		Degree:    *expertRequestObject.Degree,
		Expertise: *expertRequestObject.Expertise,
	}
	model := Expert{}
	err := db.First(&model, "id = ?", userId).Updates(expert).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Medical Expert with id %s", userId)
	return &models.GoodResponse{
			Message: message,
		},
		true

}

func UpdateFCMToken(ctx context.Context, FCMRequestObject models.FCMToken, userId string) (*models.GoodResponse, bool) {

	err := db.First(&Expert{}, "id = ?", userId).Update("fcm_token", FCMRequestObject.FCMToken).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("FCMToken Updated for Medical Expert with %s", userId)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func VerifyExpert(ctx context.Context, VerifyRequestObject models.Verify, userId string) (*models.GoodResponse, bool) {

	err := db.First(&Expert{}, "id = ?", userId).Update("director_verified", VerifyRequestObject.Verified).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Medical Expert with id %s is verfied or unverified", userId)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func BanExpert(ctx context.Context, BanRequestObject models.Ban, userId string) (*models.GoodResponse, bool) {

	err := db.First(&Expert{}, "id = ?", userId).Update("banned", BanRequestObject.Banned).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Medical Expert with id %s is banned or unbanned", userId)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func DeleteMedicalExpert(ctx context.Context, userId string) (*models.GoodResponse, bool) {

	err := db.First(&Expert{}, "id = ?", userId).Delete(&Expert{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("Medical Expert with id %s is deleted", userId)
	return &models.GoodResponse{
		Message: message,
	}, true
}

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

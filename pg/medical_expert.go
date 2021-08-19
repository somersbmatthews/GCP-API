package pg

import (
	"context"
	"fmt"

	"github.com/gircapp/api/models"
	"gorm.io/gorm"
)

type Expert struct {
	ID               string `gorm:"unique"`
	Email            string
	Name             string
	Expertise        string
	Degree           string
	DirectorVerified bool
	EmailVerified    bool
	RegisteredDevice string
	Banned           bool
}

func CreateExpertWithAutoDirectorAndEmailVerification(ctx context.Context, expertRequestObject *models.Expert, userID string) (*models.GoodResponse, bool) {

	expert := Expert{
		ID:               userID,
		Email:            *expertRequestObject.Email,
		Name:             *expertRequestObject.Name,
		Expertise:        *expertRequestObject.Expertise,
		Degree:           *expertRequestObject.Degree,
		DirectorVerified: true,
		EmailVerified:    true,
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
		EmailVerified:    false,
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

	model := Expert{
		ID:               userID,
		Email:            *expertRequestObject.Email,
		Name:             *expertRequestObject.Name,
		Expertise:        *expertRequestObject.Expertise,
		Degree:           *expertRequestObject.Degree,
		DirectorVerified: false,
		EmailVerified:    false,
	}

	err := db.Create(model).Error
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
	err := db.First(&model, "user_id = ?", userID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	return &models.GetExpertResponse{
		Name:      &model.Name,
		Degree:    &model.Degree,
		Expertise: &model.Expertise,
		Email:     &model.Email,
		Verified:  model.DirectorVerified,
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

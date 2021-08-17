package pg

import (
	"context"
	"fmt"

	"github.com/gircapp/api/models"
	"gorm.io/gorm"
)

type Expert struct {
	UserID           string `gorm:"unique"`
	Email            string
	Name             string
	Expertise        string
	Degree           string
	DirectorVerified bool
	EmailVerified    string
	RegisteredDevice string
	Banned           bool
}

func CreateExpert(ctx context.Context, expert []*models.CreateExpert, userID string) (*models.GoodResponse, bool) {
	err := db.Create(expert).Error
	if err != nil {
		return nil, false
	}
	message := fmt.Sprintf("Medical Expert with UserID: %s", userID)
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
		ID:               model.UserID,
		Name:             model.Name,
		Degree:           model.Degree,
		Expertise:        model.Expertise,
		Email:            model.Email,
		DirectorVerified: model.DirectorVerified,
	}, true
}

// func UpdateExpert(ctx context.Context, expert )

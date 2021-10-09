package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/gircapp/api/emailer"
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
	EmailConfirmationTokenIsLive bool
	DirectorTokenIsLive bool
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
		EmailConfirmationTokenIsLive: false,
	    DirectorTokenIsLive: false,
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
		EmailConfirmationTokenIsLive: false,
	    DirectorTokenIsLive: false,
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
	// TODO: send confirmation email here
	expert := Expert{
		ID:               userID,
	//	Email:            *expertRequestObject.Email,
		Name:             *expertRequestObject.Name,
		Expertise:        *expertRequestObject.Expertise,
		Degree:           *expertRequestObject.Degree,
		DirectorVerified: false,
		Registered:       true,
		EmailConfirmed:   false,
		DeviceType:       *expertRequestObject.DeviceType,
		Banned:           false,
		RegistrationTime: time.Now().Unix(),
		EmailConfirmationTokenIsLive: false,
	    DirectorTokenIsLive: false,
	}



	//	fmt.Println("create expert normally is running")
	err := emailer.SendConfirmationEmailIfNotVerified(*expertRequestObject.Email, userID, *expertRequestObject.Name, *expertRequestObject.Expertise)
	if err != nil {
		panic(err)
	}

	err = db.Create(expert).Error
	if err != nil {
		return nil, false
	}
	ok := SetEmailConfirmationJWTLive(userID)
	if !ok {
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

	err := db.Model(model).Update("email", email).Update("email_confirmed", true).Error
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

	oldExpert := &Expert{
		ID: userId,
	}
	err := db.Find(oldExpert).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	// fullNameStrSlice := strings.Split(*expertRequestObject.Name, " ")

	// firstName := fullNameStrSlice[0]


	if oldExpert.Email != *expertRequestObject.Email {
		if oldExpert.DirectorVerified {
			ok := SetEmailConfirmationJWTLive(userId)
			if !ok {
				return nil, false
			}
			emailer.SendConfirmationEmailIfVerified(*expertRequestObject.Email, *expertRequestObject.Expertise, *expertRequestObject.Name, userId)
		} else {
			ok := SetEmailConfirmationJWTLive(userId)
			if !ok {
				return nil, false
			}
			emailer.SendConfirmationEmailIfNotVerified(*expertRequestObject.Email, userId, *expertRequestObject.Name, *expertRequestObject.Expertise)
		}
	}

	model := Expert{
		ID: userId,
	}
	err = db.Model(model).Updates(expert).Error
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
	expert := &Expert{
		ID: userId,
	}
	err := db.Model(expert).Update("fcm_token", FCMRequestObject.FCMToken).Error
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
	expert := &Expert{
		ID: userId,
	}
	err := db.Model(expert).Update("director_verified", VerifyRequestObject.Verified).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	model := &Expert{
		ID: userId,
	}
	err = db.Find(model).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	if model.DirectorTokenIsLive {
		emailer.SendUserVerificationWelcomeEmail(model.Email)
	}
	message := fmt.Sprintf("Medical Expert with id %s is verfied or unverified", userId)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func BanExpert(ctx context.Context, BanRequestObject models.Ban, userId string) (*models.GoodResponse, bool) {
	expert := &Expert{
		ID: userId,
	}
	err := db.Model(expert).Update("banned", BanRequestObject.Banned).Error
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
	expert := &Expert{
		ID: userId,
	}
	err := db.Delete(expert).Error
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

func SetEmailConfirmationJWTLive(userId string) bool {
	expert := &Expert{
		ID: userId,
	}
fmt.Println("set email confirmation jwt to live is being run")
	err := db.Model(expert).Update("email_confirmation_token_is_live", true).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

func SetDirectorJWTLive(userId string) bool {
	expert := &Expert{
		ID: userId,
	}
	err := db.Model(expert).Update("director_token_is_live", true).Error
	// err := db.Model(expert).Updates(map[string]interface{}{"director_token_is_live": true}).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

func KillEmailConfimationJWT(userId string) bool {
	fmt.Println("kill email confirmation token is being run")
	expert := &Expert{
		ID: userId,
	}
	err := db.Model(expert).Update("email_confirmation_token_is_live", false).Error
	// sql := "UPDATE experts SET email_confirmation_token_is_live = true WHERE id = ?"

	// err := db.Raw(sql, userId).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("record not found error")
		fmt.Println(err)
		return false
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return true
}

func KillDirectorJWT(userId string) bool {
	expert := &Expert{
		ID: userId,
	}
	err := db.Model(expert).Update("director_token_is_live", false).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

func CheckEmailConfirmationJWT(userID string) bool{
	expert := &Expert{
		ID: userID,
	}
	err := db.Find(expert).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		panic(err)
	}
	if expert.EmailConfirmationTokenIsLive {
		return true
	} else {
		return false
	}
}

func CheckDirectorJWT(userID string) bool {
	expert := &Expert{
		ID: userID,
	}
	err := db.Find(expert).Error
	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		panic(err)
	}
	if expert.DirectorTokenIsLive {
		return true
	} else {
		return false
	}
}

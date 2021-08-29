package pg

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gdexlab/go-render/render"
	"github.com/gircapp/api/models"
	"gorm.io/gorm"
)

// func TestDeleteMedicalExpert(t *testing.T) {
// 	ctx := context.Background()
// 	_, _ = DeleteMedicalExpert(ctx, testUserID)

// }

func TestCreateMedicalExpertWithAutoDirectorAndEmailVerification(t *testing.T) {
	ctx := context.Background()

	name := "Jim Bob"
	email := "jimbob@gmail.com"
	expertise := "Otolaryngology"
	degree := "MD"
	deviceType := "iPhone 6"
	FCMToken := "1234567890abcde"

	expert := models.Expert{
		Name:       &name,
		Email:      &email,
		Expertise:  &expertise,
		Degree:     &degree,
		DeviceType: &deviceType,
		FCMToken:   &FCMToken,
	}
	_, ok := CreateExpertWithAutoDirectorAndEmailVerification(ctx, &expert, testUserID)
	if !ok {
		t.Errorf("expected CreateExpertNormally to succeed")
	}
	// get expert

	got, ok := GetMedicalExpert(ctx, testUserID)

	booleanTrue := true
	booleanFalse := false
	want := models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &name,
		Email:          &email,
		Expertise:      &expertise,
		Degree:         &degree,
		Verified:       &booleanTrue,
		EmailConfirmed: &booleanTrue,
		Banned:         &booleanFalse,
	}

	ok = reflect.DeepEqual(*got, want)
	if !ok {
		t.Error("medical expert in db does not match medical expert inserted")
	}

	_, ok = GetMedicalExpert(ctx, badUserID)
	if ok {
		t.Error("getting medical expert with bad id that was never inserted should have failed")
	}

	// delete expert

	_, ok = DeleteMedicalExpert(ctx, testUserID)
	if !ok {
		t.Error("delete medical expert failed")
	}
	_, ok = GetMedicalExpert(ctx, testUserID)
	if ok {
		t.Error("get medical expert should have failed after deleting medical expert in db")
	}
}

func TestCreateMedicalExpertWithAutoDirectorWithoutEmailVerification(t *testing.T) {
	ctx := context.Background()

	name := "Jim Bob"
	email := "jimbob@gmail.com"
	expertise := "Otolaryngology"
	degree := "MD"
	deviceType := "iPhone 6"
	FCMToken := "1234567890abcde"

	expert := models.Expert{
		Name:       &name,
		Email:      &email,
		Expertise:  &expertise,
		Degree:     &degree,
		DeviceType: &deviceType,
		FCMToken:   &FCMToken,
	}
	_, ok := CreateExpertWithAutoDirectorWithoutEmailVerification(ctx, &expert, testUserID)

	// get expert
	got, ok := GetMedicalExpert(ctx, testUserID)

	booleanTrue := true
	booleanFalse := false
	want := models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &name,
		Email:          &email,
		Expertise:      &expertise,
		Degree:         &degree,
		Verified:       &booleanTrue,
		EmailConfirmed: &booleanFalse,
		Banned:         &booleanFalse,
	}
	ok = reflect.DeepEqual(*got, want)
	if !ok {
		fmt.Println("THIS IS GET EXPERT RESPOSE WANT")
		fmt.Println(render.Render(want))
		fmt.Println("this is get expert got")
		fmt.Println(render.Render(*got))
		t.Error("medical expert in db does not match medical expert inserted")
	}

	// delete expert

	_, ok = DeleteMedicalExpert(ctx, testUserID)
	if !ok {
		t.Error("delete medical expert failed")
	}
	_, ok = GetMedicalExpert(ctx, testUserID)
	if ok {
		t.Error("get medical expert should have failed after deleting medical expert in db")
	}
}

func TestCreateMedicalExpertNormally(t *testing.T) {
	ctx := context.Background()

	name := "Jim Bob"
	email := "jimbob@gmail.com"
	expertise := "Otolaryngology"
	degree := "MD"
	deviceType := "iPhone 6"
	FCMToken := "1234567890abcde"

	expert := models.Expert{
		Name:       &name,
		Email:      &email,
		Expertise:  &expertise,
		Degree:     &degree,
		DeviceType: &deviceType,
		FCMToken:   &FCMToken,
	}

	_, ok := CreateExpertNormally(ctx, &expert, testUserID)

	got, ok := GetMedicalExpert(ctx, testUserID)

	//	booleanTrue := true
	booleanFalse := false
	want := models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &name,
		Email:          &email,
		Expertise:      &expertise,
		Degree:         &degree,
		Verified:       &booleanFalse,
		EmailConfirmed: &booleanFalse,
		Banned:         &booleanFalse,
	}

	ok = reflect.DeepEqual(*got, want)
	if !ok {
		t.Error("medical expert in db does not match medical expert inserted")
	}

	// delete expert

	// _, ok = DeleteMedicalExpert(ctx, testUserID)
	// if !ok {
	// 	t.Error("delete medical expert failed")
	// }
	// _, ok = GetMedicalExpert(ctx, testUserID)
	// if ok {
	// 	t.Error("get medical expert should have failed after deleting medical expert in db")

}

func TestUpdateMedicalExpert(t *testing.T) {
	name := "Jim Bob"
	email := "jimbob@gmail.com"
	expertise := "Otolaryngology"
	degree := "MD"
	deviceType := "iPhone 6"
	FCMToken := "1234567890abcde"

	ctx := context.Background()
	booleanTrue := true
	booleanFalse := false
	verifyExpertObject := models.Verify{
		UserID:   &testUserID,
		Verified: &booleanTrue,
	}
	// verify expert
	_, ok := VerifyExpert(ctx, verifyExpertObject, testUserID)
	// get expert
	got, ok := GetMedicalExpert(ctx, testUserID)

	want := models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &name,
		Email:          &email,
		Expertise:      &expertise,
		Degree:         &degree,
		Verified:       &booleanTrue,
		EmailConfirmed: &booleanFalse,
		Banned:         &booleanFalse,
	}

	ok = reflect.DeepEqual(*got, want)
	if !ok {
		t.Error("medical expert in db does not match medical expert updated with verify")
	}

	// unverify expert

	unverifyExpertObject := models.Verify{
		UserID:   &testUserID,
		Verified: &booleanFalse,
	}
	// verify expert
	_, ok = VerifyExpert(ctx, unverifyExpertObject, testUserID)
	// get expert
	got, ok = GetMedicalExpert(ctx, testUserID)

	want = models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &name,
		Email:          &email,
		Expertise:      &expertise,
		Degree:         &degree,
		Verified:       &booleanFalse,
		EmailConfirmed: &booleanFalse,
		Banned:         &booleanFalse,
	}

	ok = reflect.DeepEqual(*got, want)
	if !ok {
		t.Error("medical expert in db does not match medical expert updated with verify")
	}

	// ban expert
	banRequestObject := models.Ban{
		Banned: &booleanTrue,
	}
	_, ok = BanExpert(ctx, banRequestObject, testUserID)
	// get expert
	got, ok = GetMedicalExpert(ctx, testUserID)

	want = models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &name,
		Email:          &email,
		Expertise:      &expertise,
		Degree:         &degree,
		Verified:       &booleanFalse,
		EmailConfirmed: &booleanFalse,
		Banned:         &booleanTrue,
	}

	ok = reflect.DeepEqual(*got, want)
	if !ok {
		t.Error("medical expert in db does not match medical expert updated with ban")
	}
	// unban expert
	banRequestObject = models.Ban{
		Banned: &booleanFalse,
	}
	_, ok = BanExpert(ctx, banRequestObject, testUserID)
	// get expert
	got, ok = GetMedicalExpert(ctx, testUserID)

	want = models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &name,
		Email:          &email,
		Expertise:      &expertise,
		Degree:         &degree,
		Verified:       &booleanFalse,
		EmailConfirmed: &booleanFalse,
		Banned:         &booleanFalse,
	}

	ok = reflect.DeepEqual(*got, want)
	if !ok {
		t.Error("medical expert in db does not match medical expert updated with ban")
	}
	// update FCM token
	newFCMToken := "asdgjhouwbet"
	fCMRequestObject := models.FCMToken{
		FCMToken: &newFCMToken,
	}

	_, ok = UpdateFCMToken(ctx, fCMRequestObject, testUserID)

	// get expert
	expert := Expert{}
	err := db.First(&expert, "id = ?", testUserID).Error
	if err == gorm.ErrRecordNotFound {
		t.Error("cannont find medical expert in getting after fcm token updatedS")
	} else if err != nil {
		t.Error("some other error with getting medical expert")
	}

	newName := "Jim Bob"
	// newEmail := "jimbob@gmail.com"
	newExpertise := "Otolaryngology"
	newDegree := "MD"

	// newExpert
	newExpert := models.Expert{
		Name:       &newName,
		Email:      &email,
		Expertise:  &newExpertise,
		Degree:     &newDegree,
		DeviceType: &deviceType,
		FCMToken:   &FCMToken,
	}
	// update expert info
	_, ok = UpdateMedicalExpert(ctx, newExpert, testUserID)
	// get expert
	got, ok = GetMedicalExpert(ctx, testUserID)

	want = models.GetExpertResponse{
		ID:             &testUserID,
		Name:           &newName,
		Email:          &email,
		Expertise:      &newExpertise,
		Degree:         &newDegree,
		Verified:       &booleanFalse,
		EmailConfirmed: &booleanFalse,
		Banned:         &booleanFalse,
	}

	ok = reflect.DeepEqual(*got, want)
	if !ok {
		t.Error("medical expert in db does not match medical expert updated with UpdateMedicalExpert")
	}
	// delete expert

	_, ok = DeleteMedicalExpert(ctx, testUserID)
	if !ok {
		t.Error("delete medical expert failed")
	}
	_, ok = GetMedicalExpert(ctx, testUserID)
	if ok {
		t.Error("get medical expert should have failed after deleting medical expert in db")

	}

}

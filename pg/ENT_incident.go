package pg

import (
	"context"

	"github.com/gircapp/api/models"
	"gorm.io/gorm"
)

type ENTIncident struct {
	ID string `json:"ID"`

	Country string `json:"country"`

	Year string `json:"year"`

	AgeYears string `json:"ageYears"`

	AgeMonths string `json:"ageMonths"`

	Gender string `json:"gender"`

	IncidentDescription string `json:"incidentDescription"`

	DaysUntilRemoval float32 `json:"daysUntilRemoval"`

	HoursUntilRemoval float32 `json:"hoursUntilRemoval"`

	MinutesUntilRemoval float32 `json:"minutesUntilRemoval"`

	RemovalStrategy float32 `json:"removalStrategy"`

	OpenSurgery string `json:"openSurgery"`

	EaseOfRemoval string `json:"easeOfRemoval"`

	WasIncidentLifeThreatening string `json:"wasIncidentLifeThreatening"`

	Symptoms []string `json:"symptoms"`

	CustomSymptoms []string `json:"customSymptoms"`

	SymptomSeverity string `json:"symptomSeverity"`

	Complications []string `json:"complications"`

	CustomComplications []string `json:"customComplications"`

	Anesthesia string `json:"anesthesia"`

	Prognosis string `json:"prognosis"`

	HospitalStay string `json:"hospitalStay"`

	DeviceType string `json:"deviceType"`

	Submitted bool `json:"submitted,omitempty"`

	SwallowedObjects []ENTSwallowedObject `json:"swallowedObjects"`
}

type ENTSwallowedObject struct {
	ID string `json:"ID"`

	RadioOpacity string `json:"radioOpacity"`

	Imaging string `json:"imaging"`

	AnteriorPhoto string `json:"anteriorPhoto"`

	PosteriorPhoto string `json:"posteriorPhoto"`

	LateralPhoto string `json:"lateralPhoto"`

	AnteriorLongestLength string `json:"anteriorLongestLength"`

	PosteriorLongestLength string `json:"posteriorLongestLength"`

	LateralLongestLength string `json:"lateralLongestLength"`

	ObjectLocation string `json:"objectLocation"`

	NumberOfThisObject string `json:"numberOfThisObject"`

	ObjectIntact string `json:"objectIntact"`

	NumberOfPieces string `json:"numberOfPieces"`

	ObjectDescription string `json:"objectDescription"`

	ObjectShape string `json:"objectShape"`

	ObjectCustomShape string `json:"objectCustomShape"`

	ObjectDimensionality string `json:"objectDimensionality"`

	OtherCharacteristics []string `json:"otherCharacteristics"`

	Material string `json:"material"`

	CustomMaterial string `json:"customMaterial"`

	IsBatteryOrMagnet string `json:"isBatteryOrMagnet"`

	BatteryType string `json:"batteryType"`

	CustomBatteryType string `json:"customBatteryType"`

	BatteryImprintCode string `json:"batteryImprintCode"`

	MitigatingFeatures []string `json:"mitigatingFeatures"`

	CustomMitigatingFeatures []string `json:"customMitigatingFeatures"`

	NegativePoleDirection string `json:"negativePoleDirection"`

	Honey string `json:"honey"`

	Sucralfate string `json:"sucralfate"`

	AceticAcid string `json:"aceticAcid"`

	MagnetType string `json:"magnetType"`

	CustomMagnetType string `json:"customMagnetType"`

	DeviceType string `json:"deviceType"`

	Submitted bool `json:"submitted,omitempty"`
}

func CreateENTIncident(ctx context.Context, incidentRequest *models.ENTIncident, userID string) (*models.GoodResponse, bool) {
	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	bytea := getByteaFromString(encryptedUserID)

	// var swallowedObjects []ENTSwallowedObject
	// for _, object := range incidentRequest.SwallowedObjects {

	// }
	incidentRequest.EncryptedExpertID = bytea

	err = db.Create(&incidentRequest).Error
	if err != nil {
		return nil, false
	}
	return &models.GoodResponse{
		Message: "ENT Incident Created",
	}, true
}

func DeleteENTIncident(ctx context.Context, incidentID string) (*models.GoodResponse, bool) {
	err := db.Delete(&models.ENTIncident{}, "id = ?", incidentID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	return &models.GoodResponse{
		Message: "ENT Incident Created",
	}, true
}

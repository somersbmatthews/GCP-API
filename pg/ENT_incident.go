package pg

import (
	"context"
	"fmt"

	"github.com/gircapp/api/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ENTIncident struct {
	//CreatedAt int64

	ID string `gorm:"unique;not null;primary_key"`

	EncryptedExpertID string `gorm:"type:bytea"`

	Country string

	Year string

	AgeYears string

	AgeMonths string

	Gender string

	IncidentDescription string

	DaysUntilRemoval int64

	HoursUntilRemoval int64

	MinutesUntilRemoval int64

	RemovalStrategy string

	RemovalSetting string

	OpenSurgery string

	EaseOfRemoval string

	WasIncidentLifeThreatening string

	Symptoms pq.StringArray `gorm:"type:text[]"`

	CustomSymptoms pq.StringArray `gorm:"type:text[]"`

	SymptomSeverity string

	Complications pq.StringArray `gorm:"type:text[]"`

	CustomComplications pq.StringArray `gorm:"type:text[]"`

	Anesthesia string

	Prognosis string

	HospitalStay string

	DeviceType string

	Submitted bool

	Validated bool

	SwallowedObjects []SwallowedObject `gorm:"foreignKey:IncidentID;references:ID"`
}

// `gorm:"foreignkey:ID"`

type SwallowedObject struct {
	ID string `gorm:"unique; not null"`

	IncidentID string `gorm:"not null"`

	RadioOpacity string

	Imaging string

	AnteriorPhoto string

	PosteriorPhoto string

	LateralPhoto string

	AnteriorLongestLength float64

	PosteriorLongestLength float64

	LateralLongestLength float64

	ObjectLocation string

	NumberOfThisObject string

	ObjectIntact string

	NumberOfPieces string

	ObjectDescription string

	ObjectShape string

	ObjectCustomShape string

	ObjectDimensionality string

	OtherCharacteristics pq.StringArray `gorm:"type:text[]"`

	Material string

	CustomMaterial string

	IsBatteryOrMagnet string

	BatteryType string

	CustomBatteryType string

	BatteryImprintCode string

	MitigatingFeatures pq.StringArray `gorm:"type:text[]"`

	CustomMitigatingFeatures pq.StringArray `gorm:"type:text[]"`

	NegativePoleDirection string

	Honey string

	Sucralfate string

	AceticAcid string

	MagnetType string

	CustomMagnetType string

	DeviceType string

	Validated bool

	Submitted bool
}

func CreateENTIncident(ctx context.Context, incidentRequestObject *models.ENTIncident, userID string) (*models.GoodResponse, bool) {
	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	bytea := getByteaFromString(encryptedUserID)

	incidentRequest := incidentRequestObject

	var swallowedObjects []SwallowedObject
	for _, objectValue := range incidentRequest.SwallowedObjects {
		object := objectValue
		newObject := SwallowedObject{
			ID:                       *object.ID,
			IncidentID:               *incidentRequest.ID,
			RadioOpacity:             *object.RadioOpacity,
			Imaging:                  *object.Imaging,
			AnteriorPhoto:            *object.AnteriorPhoto,
			PosteriorPhoto:           *object.PosteriorPhoto,
			LateralPhoto:             *object.LateralPhoto,
			AnteriorLongestLength:    *object.AnteriorLongestLength,
			PosteriorLongestLength:   *object.PosteriorLongestLength,
			LateralLongestLength:     *object.LateralLongestLength,
			ObjectLocation:           *object.ObjectLocation,
			NumberOfThisObject:       *object.NumberOfThisObject,
			ObjectIntact:             *object.ObjectIntact,
			NumberOfPieces:           *object.NumberOfPieces,
			ObjectDescription:        *object.ObjectDescription,
			ObjectShape:              *object.ObjectShape,
			ObjectCustomShape:        *object.ObjectCustomShape,
			ObjectDimensionality:     *object.ObjectDimensionality,
			OtherCharacteristics:     pq.StringArray(object.OtherCharacteristics),
			Material:                 *object.Material,
			CustomMaterial:           *object.CustomMaterial,
			IsBatteryOrMagnet:        *object.IsBatteryOrMagnet,
			BatteryType:              *object.BatteryType,
			CustomBatteryType:        *object.CustomBatteryType,
			BatteryImprintCode:       *object.BatteryImprintCode,
			MitigatingFeatures:       pq.StringArray(object.MitigatingFeatures),
			CustomMitigatingFeatures: pq.StringArray(object.CustomMitigatingFeatures),
			NegativePoleDirection:    *object.NegativePoleDirection,
			Honey:                    *object.Honey,
			Sucralfate:               *object.Sucralfate,
			AceticAcid:               *object.AceticAcid,
			MagnetType:               *object.MagnetType,
			CustomMagnetType:         *object.CustomMagnetType,
			DeviceType:               *object.DeviceType,
			Submitted:                *object.Submitted,
			Validated:                *object.Validated,
		}

		swallowedObjects = append(swallowedObjects, newObject)
	}

	newIncident := &ENTIncident{
		///	CreatedAt:                  time.Now().Unix(),
		ID:                         *incidentRequest.ID,
		EncryptedExpertID:          bytea,
		Country:                    *incidentRequest.Country,
		Year:                       *incidentRequest.Year,
		AgeYears:                   *incidentRequest.AgeYears,
		AgeMonths:                  *incidentRequest.AgeMonths,
		Gender:                     *incidentRequest.Gender,
		IncidentDescription:        *incidentRequest.IncidentDescription,
		DaysUntilRemoval:           *incidentRequest.DaysUntilRemoval,
		HoursUntilRemoval:          *incidentRequest.HoursUntilRemoval,
		MinutesUntilRemoval:        *incidentRequest.MinutesUntilRemoval,
		RemovalStrategy:            *incidentRequest.RemovalStrategy,
		RemovalSetting:             *incidentRequest.RemovalSetting,
		OpenSurgery:                *incidentRequest.OpenSurgery,
		EaseOfRemoval:              *incidentRequest.EaseOfRemoval,
		WasIncidentLifeThreatening: *incidentRequest.WasIncidentLifeThreatening,
		Symptoms:                   pq.StringArray(incidentRequest.Symptoms),
		CustomSymptoms:             pq.StringArray(incidentRequest.CustomSymptoms),
		SymptomSeverity:            *incidentRequest.SymptomSeverity,
		Complications:              pq.StringArray(incidentRequest.Complications),
		CustomComplications:        pq.StringArray(incidentRequest.CustomComplications),
		Anesthesia:                 *incidentRequest.Anesthesia,
		Prognosis:                  *incidentRequest.Prognosis,
		HospitalStay:               *incidentRequest.HospitalStay,
		DeviceType:                 *incidentRequest.DeviceType,
		Validated:                  *incidentRequest.Validated,
		Submitted:                  *incidentRequest.Submitted,
		SwallowedObjects:           swallowedObjects,
	}

	err = db.Create(newIncident).Error
	if err != nil {
		return nil, false
	}
	return &models.GoodResponse{
		Message: "ENT Incident Created",
	}, true
}

func GetENTIncidents(ctx context.Context, userID string) (*models.GetENTIncidentsGoodResponse, bool) {
	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	ENTIncidents := []ENTIncident{}

	// sql := "SELECT * FROM ent_incidents WHERE encrypted_expert_id = ? ORDER BY created_at DESC"
	sql := "SELECT * FROM ent_incidents WHERE encrypted_expert_id = ?"

	bytea := getByteaFromString(encryptedUserID)

	err = db.Raw(sql, bytea).Scan(&ENTIncidents).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	incidentResponse := []*models.ENTIncident{}

	for _, incidentValue := range ENTIncidents {
		incident := incidentValue
		var swallowedObjects []*models.SwallowedObject
		var swallowedObjectsFromDB = []SwallowedObject{}
		err = db.Find(&swallowedObjectsFromDB, "incident_id", incident.ID).Error
		if err == gorm.ErrRecordNotFound {
			return nil, false
		} else if err != nil {
			panic(err)
		}

		for _, objectValue := range swallowedObjectsFromDB {

			object := objectValue
			newSwallowedObject := models.SwallowedObject{
				ID:                       &object.ID,
				IncidentID:               incident.ID,
				AceticAcid:               &object.AceticAcid,
				AnteriorLongestLength:    &object.AnteriorLongestLength,
				AnteriorPhoto:            &object.AnteriorPhoto,
				BatteryImprintCode:       &object.BatteryImprintCode,
				BatteryType:              &object.BatteryType,
				CustomBatteryType:        &object.CustomBatteryType,
				CustomMagnetType:         &object.CustomMagnetType,
				CustomMaterial:           &object.CustomMaterial,
				CustomMitigatingFeatures: convertPQToStringArray(object.CustomMitigatingFeatures),
				DeviceType:               &object.DeviceType,
				Honey:                    &object.Honey,
				Imaging:                  &object.Imaging,
				IsBatteryOrMagnet:        &object.IsBatteryOrMagnet,
				LateralLongestLength:     &object.LateralLongestLength,
				LateralPhoto:             &object.LateralPhoto,
				MagnetType:               &object.MagnetType,
				Material:                 &object.Material,
				MitigatingFeatures:       convertPQToStringArray(object.MitigatingFeatures),
				NegativePoleDirection:    &object.NegativePoleDirection,
				NumberOfPieces:           &object.NumberOfPieces,
				NumberOfThisObject:       &object.NumberOfThisObject,
				ObjectCustomShape:        &object.ObjectCustomShape,
				ObjectDescription:        &object.ObjectDescription,
				ObjectDimensionality:     &object.ObjectDimensionality,
				ObjectIntact:             &object.ObjectIntact,
				ObjectLocation:           &object.ObjectLocation,
				ObjectShape:              &object.ObjectShape,
				OtherCharacteristics:     convertPQToStringArray(object.OtherCharacteristics),
				PosteriorLongestLength:   &object.PosteriorLongestLength,
				PosteriorPhoto:           &object.PosteriorPhoto,
				RadioOpacity:             &object.RadioOpacity,
				Validated:                &object.Validated,
				Submitted:                &object.Submitted,
				Sucralfate:               &object.Sucralfate,
			}
			swallowedObjects = append(swallowedObjects, &newSwallowedObject)
		}

		newIncident := models.ENTIncident{
			ID:                         &incident.ID,
			AgeMonths:                  &incident.AgeMonths,
			AgeYears:                   &incident.AgeYears,
			Anesthesia:                 &incident.Anesthesia,
			Complications:              convertPQToStringArray(incident.Complications),
			Country:                    &incident.Country,
			CustomComplications:        convertPQToStringArray(incident.CustomComplications),
			CustomSymptoms:             convertPQToStringArray(incident.CustomSymptoms),
			DaysUntilRemoval:           &incident.DaysUntilRemoval,
			DeviceType:                 &incident.DeviceType,
			EaseOfRemoval:              &incident.EaseOfRemoval,
			Gender:                     &incident.Gender,
			HospitalStay:               &incident.HospitalStay,
			HoursUntilRemoval:          &incident.HoursUntilRemoval,
			IncidentDescription:        &incident.IncidentDescription,
			MinutesUntilRemoval:        &incident.MinutesUntilRemoval,
			OpenSurgery:                &incident.OpenSurgery,
			Prognosis:                  &incident.Prognosis,
			RemovalStrategy:            &incident.RemovalStrategy,
			RemovalSetting:             &incident.RemovalSetting,
			Validated:                  &incident.Validated,
			Submitted:                  &incident.Submitted,
			SwallowedObjects:           swallowedObjects,
			SymptomSeverity:            &incident.SymptomSeverity,
			Symptoms:                   convertPQToStringArray(incident.Symptoms),
			WasIncidentLifeThreatening: &incident.WasIncidentLifeThreatening,
			Year:                       &incident.Year,
		}

		incidentResponse = append(incidentResponse, &newIncident)

	}

	return &models.GetENTIncidentsGoodResponse{
		Incidents: incidentResponse,
	}, true
}

func UpdateENTIncident(ctx context.Context, ENTIncidentRequest models.ENTIncident, userID string) (*models.GoodResponse, bool) {
	encryptedUserID, err := encryptUserID(userID)
	if err != nil {
		panic(err)
	}

	bytea := getByteaFromString(encryptedUserID)

	var swallowedObjects []SwallowedObject
	for _, objectValue := range ENTIncidentRequest.SwallowedObjects {
		object := objectValue
		newObject := SwallowedObject{
			ID:                       *object.ID,
			IncidentID:               *ENTIncidentRequest.ID,
			RadioOpacity:             *object.RadioOpacity,
			Imaging:                  *object.Imaging,
			AnteriorPhoto:            *object.AnteriorPhoto,
			PosteriorPhoto:           *object.PosteriorPhoto,
			LateralPhoto:             *object.LateralPhoto,
			AnteriorLongestLength:    *object.AnteriorLongestLength,
			PosteriorLongestLength:   *object.PosteriorLongestLength,
			LateralLongestLength:     *object.LateralLongestLength,
			ObjectLocation:           *object.ObjectLocation,
			NumberOfThisObject:       *object.NumberOfThisObject,
			ObjectIntact:             *object.ObjectIntact,
			NumberOfPieces:           *object.NumberOfPieces,
			ObjectDescription:        *object.ObjectDescription,
			ObjectShape:              *object.ObjectShape,
			ObjectCustomShape:        *object.ObjectCustomShape,
			ObjectDimensionality:     *object.ObjectDimensionality,
			OtherCharacteristics:     pq.StringArray(object.OtherCharacteristics),
			Material:                 *object.Material,
			CustomMaterial:           *object.CustomMaterial,
			IsBatteryOrMagnet:        *object.IsBatteryOrMagnet,
			BatteryType:              *object.BatteryType,
			CustomBatteryType:        *object.CustomBatteryType,
			BatteryImprintCode:       *object.BatteryImprintCode,
			MitigatingFeatures:       pq.StringArray(object.MitigatingFeatures),
			CustomMitigatingFeatures: pq.StringArray(object.CustomMitigatingFeatures),
			NegativePoleDirection:    *object.NegativePoleDirection,
			Honey:                    *object.Honey,
			Sucralfate:               *object.Sucralfate,
			AceticAcid:               *object.AceticAcid,
			MagnetType:               *object.MagnetType,
			CustomMagnetType:         *object.CustomMagnetType,
			DeviceType:               *object.DeviceType,
			Validated:                *object.Validated,
			Submitted:                *object.Submitted,
		}

		swallowedObjects = append(swallowedObjects, newObject)
	}

	incident := &ENTIncident{
		ID:                         *ENTIncidentRequest.ID,
		EncryptedExpertID:          bytea,
		Country:                    *ENTIncidentRequest.Country,
		Year:                       *ENTIncidentRequest.Year,
		AgeYears:                   *ENTIncidentRequest.AgeYears,
		AgeMonths:                  *ENTIncidentRequest.AgeMonths,
		Gender:                     *ENTIncidentRequest.Gender,
		IncidentDescription:        *ENTIncidentRequest.IncidentDescription,
		DaysUntilRemoval:           *ENTIncidentRequest.DaysUntilRemoval,
		HoursUntilRemoval:          *ENTIncidentRequest.HoursUntilRemoval,
		MinutesUntilRemoval:        *ENTIncidentRequest.MinutesUntilRemoval,
		RemovalStrategy:            *ENTIncidentRequest.RemovalStrategy,
		RemovalSetting:             *ENTIncidentRequest.RemovalSetting,
		OpenSurgery:                *ENTIncidentRequest.OpenSurgery,
		EaseOfRemoval:              *ENTIncidentRequest.EaseOfRemoval,
		WasIncidentLifeThreatening: *ENTIncidentRequest.WasIncidentLifeThreatening,
		Symptoms:                   pq.StringArray(ENTIncidentRequest.Symptoms),
		CustomSymptoms:             pq.StringArray(ENTIncidentRequest.CustomSymptoms),
		SymptomSeverity:            *ENTIncidentRequest.SymptomSeverity,
		Complications:              pq.StringArray(ENTIncidentRequest.Complications),
		CustomComplications:        pq.StringArray(ENTIncidentRequest.CustomComplications),
		Anesthesia:                 *ENTIncidentRequest.Anesthesia,
		Prognosis:                  *ENTIncidentRequest.Prognosis,
		HospitalStay:               *ENTIncidentRequest.HospitalStay,
		DeviceType:                 *ENTIncidentRequest.DeviceType,
		Validated:                  *ENTIncidentRequest.Validated,
		Submitted:                  *ENTIncidentRequest.Submitted,
	}

	err = db.Find(&ENTIncident{}, "id = ?", ENTIncidentRequest.ID).Updates(incident).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}

	for _, object := range swallowedObjects {
		err = db.Find(&SwallowedObject{}, "id = ?", &object.ID).Updates(object).Error
		if err == gorm.ErrRecordNotFound {
			return nil, false
		} else if err != nil {
			panic(err)
		}
	}

	message := fmt.Sprintf("incident updated with id: %s", incident.ID)
	return &models.GoodResponse{
		Message: message,
	}, true
}

func DeleteENTIncident(ctx context.Context, incidentID string) (*models.GoodResponse, bool) {
	err := db.Delete(&ENTIncident{}, "id = ?", incidentID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	err = db.Delete(&SwallowedObject{}, "incident_id = ?", incidentID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, false
	} else if err != nil {
		panic(err)
	}
	return &models.GoodResponse{
		Message: "ENT Incident Created",
	}, true
}

func convertPQToStringArray(strArray pq.StringArray) []string {
	sliceOfStrings := []string(strArray)
	fmt.Println(sliceOfStrings)
	fmt.Printf("%v \n", len(sliceOfStrings))
	// fmt.Printf("THIS IS STRING: %v \n", str)
	// sliceOfStrings := strings.Split(str, ",")
	// fmt.Printf("THIS IS SLICE OF STRINGS: %v \n", str)
	return sliceOfStrings
}

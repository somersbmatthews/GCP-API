package pg

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/gdexlab/go-render/render"
	"github.com/gircapp/api/models"
)

var objectID3 = "1234567890"
var objectID4 = "1234567790"
var objectID5 = "1234567770"
var objectID6 = "1234566789"
var aceticAcid3 = "yes"
var aceticAcid4 = "yes"
var anteriorLongestLength3 = 32.6
var anteriorLongestLength4 = 33.6
var anteriorPhoto3 = "photo/asdf.jpg"
var anteriorPhoto4 = "photo/poiu.jpg"
var batteryImprintCode3 = "40"
var batteryImprintCode4 = "50"
var batteryType3 = "Button Battery"
var batteryType4 = "AA"
var customBatteryType3 = "big"
var customBatteryType4 = "small"
var customMagnetType3 = "fridge"
var customMagnetType4 = "big"
var customMateral3 = "paper"
var customMateral4 = "plastic"
var customMitigatingFeatures3 = []string{"feature1", "feature2"}
var customMitigatingFeatures4 = []string{"feature3", "feature4"}
var honey3 = "yes"
var honey4 = "no"
var imaging3 = "MRI"
var imaging4 = "XRay"
var isBatteryOrMagnet3 = "Neither"
var isBatteryOrMagnet4 = "Magnet"
var lateralLongestLength3 = 22.1
var lateralLongestLength4 = 56.7
var lateralPhoto3 = "photo2.jpg"
var lateralPhoto4 = "photo2.jpg"
var magnetType3 = "Button Battery"
var magnetType4 = "AA"
var material3 = "paper"
var material4 = "wood"
var mitigatingFeatures3 = []string{"feature5", "feature6"}
var mitigatingFeatures4 = []string{"feature7", "feature8"}
var negativePoleDirection3 = "positive"
var negativePoleDirection4 = "negative"
var numberOfPieces3 = "5"
var numberOfPieces4 = "8"
var numberOfThisObject3 = "4"
var numberOfThisObject4 = "3"
var objectCustomShape3 = "square"
var objectCustomShape4 = "round"
var objectDescription3 = "sphere"
var objectDescription4 = "square"
var objectDimensionality3 = "3D"
var objectDimensionality4 = "2D"
var objectIntact3 = "No"
var objectIntact4 = "Yes"
var objectLocation3 = "esophagus"
var objectLocation4 = "stomach"
var objectShape3 = "square"
var objectShape4 = "round"
var otherCharacteristics3 = []string{"char1", "char2"}
var otherCharacteristics4 = []string{"char3", "char4"}
var posteriorLongestLength3 = 43.6
var posteriorLongestLength4 = 22.3
var posteriorPhoto3 = "photo4.jpg"
var posteriorPhoto4 = "photo5.jpg"
var radioOpacity3 = "No"
var radioOpacity4 = "Yes"
var sucrafate3 = "No"
var sucrafate4 = "Yes"
var objectValidated3 = true
var objectValidated4 = false
var objectSubmitted3 = true
var objectSubmitted4 = false

var ENTIncidentID3 = "1234567890"
var ENTIncidentID4 = "1234567790"
var ageMonths3 = "8"
var ageMonths4 = "9"
var ageYears3 = "2"
var ageYears4 = "4"
var anesthesia3 = "No"
var anesthesia4 = "No"
var complications3 = []string{"vomiting", "choking"}
var complications4 = []string{"choking", "vomiting"}
var country3 = "USA"
var country4 = "Mexico"
var customComplications3 = []string{"coughing", "choking"}
var customComplications4 = []string{"vomiting", "coughing"}
var customSymptoms3 = []string{"vomiting", "choking"}
var customSymptoms4 = []string{"vomiting", "choking"}
var daysUntilRemoval3 int64 = 3
var daysUntilRemoval4 int64 = 4
var deviceType3 = "iPhone 6"
var deviceType4 = "iPhone 8"
var easeOfRemoval3 = "No"
var easeOfRemoval4 = "Yes"
var gender3 = "Male"
var gender4 = "Female"
var hospitalStay3 = "Overnight"
var hospitalStay4 = "One Week"
var hoursUntilRemoval3 int64 = 3
var hoursUntilRemoval4 int64 = 6
var incidentDescription3 = "Example Discription"
var incidentDescription4 = "Example Discription2"
var minutesUntilRemoval3 int64 = 7
var minutesUntilRemoval4 int64 = 8
var openSurgery3 = "Yes"
var openSurgery4 = "No"
var prognosis3 = "Death"
var prognosis4 = "Alive"
var removalStrategy3 = "Forceps"
var removalStrategy4 = "Surgery"
var removalSetting3 = "OR"
var removalSetting4 = "Doctor's Office"
var symptomSeverity3 = "severe"
var symptomSeverity4 = "not severe"
var symptoms3 = []string{"sneezing", "choking"}
var symptoms4 = []string{"coughing", "sneezing"}
var wasIncidentLifeThreatening3 = "no"
var wasIncidentLifeThreatening4 = "yes"
var year3 = "2009"
var year4 = "2015"
var incidentValidated3 = false
var incidentValidated4 = true
var incidentSubmitted3 = false
var incidentSubmitted4 = true

var swallowedObject1 = models.SwallowedObject{
	ID:                       &objectID3,
	IncidentID:               incident1.ID,
	AceticAcid:               &aceticAcid3,
	AnteriorLongestLength:    &anteriorLongestLength3,
	AnteriorPhoto:            &anteriorPhoto3,
	BatteryImprintCode:       &batteryImprintCode3,
	BatteryType:              &batteryType3,
	CustomBatteryType:        &customBatteryType3,
	CustomMagnetType:         &customMagnetType3,
	CustomMaterial:           &customMateral3,
	CustomMitigatingFeatures: customMitigatingFeatures3,
	DeviceType:               &deviceType3,
	Honey:                    &honey3,
	Imaging:                  &imaging3,
	IsBatteryOrMagnet:        &isBatteryOrMagnet3,
	LateralLongestLength:     &lateralLongestLength3,
	LateralPhoto:             &lateralPhoto3,
	MagnetType:               &magnetType3,
	Material:                 &material3,
	MitigatingFeatures:       mitigatingFeatures3,
	NegativePoleDirection:    &negativePoleDirection3,
	NumberOfPieces:           &numberOfPieces3,
	NumberOfThisObject:       &numberOfThisObject3,
	ObjectCustomShape:        &objectCustomShape3,
	ObjectDescription:        &objectDescription3,
	ObjectDimensionality:     &objectDimensionality3,
	ObjectIntact:             &objectIntact3,
	ObjectLocation:           &objectLocation3,
	ObjectShape:              &objectShape3,
	OtherCharacteristics:     otherCharacteristics3,
	PosteriorLongestLength:   &posteriorLongestLength3,
	PosteriorPhoto:           &posteriorPhoto3,
	RadioOpacity:             &radioOpacity3,
	Sucralfate:               &sucrafate3,
	Validated:                &objectValidated3,
	Submitted:                &objectSubmitted3,
}

var swallowedObject2 = models.SwallowedObject{
	ID:                       &objectID4,
	IncidentID:               incident2.ID,
	AceticAcid:               &aceticAcid4,
	AnteriorLongestLength:    &anteriorLongestLength4,
	AnteriorPhoto:            &anteriorPhoto4,
	BatteryImprintCode:       &batteryImprintCode4,
	BatteryType:              &batteryType4,
	CustomBatteryType:        &customBatteryType4,
	CustomMagnetType:         &customMagnetType4,
	CustomMaterial:           &customMateral4,
	CustomMitigatingFeatures: customMitigatingFeatures4,
	DeviceType:               &deviceType4,
	Honey:                    &honey4,
	Imaging:                  &imaging4,
	IsBatteryOrMagnet:        &isBatteryOrMagnet4,
	LateralLongestLength:     &lateralLongestLength4,
	LateralPhoto:             &lateralPhoto4,
	MagnetType:               &magnetType4,
	Material:                 &material4,
	MitigatingFeatures:       mitigatingFeatures4,
	NegativePoleDirection:    &negativePoleDirection4,
	NumberOfPieces:           &numberOfPieces4,
	NumberOfThisObject:       &numberOfThisObject4,
	ObjectCustomShape:        &objectCustomShape4,
	ObjectDescription:        &objectDescription4,
	ObjectDimensionality:     &objectDimensionality4,
	ObjectIntact:             &objectIntact4,
	ObjectLocation:           &objectLocation4,
	ObjectShape:              &objectShape4,
	OtherCharacteristics:     otherCharacteristics4,
	PosteriorLongestLength:   &posteriorLongestLength4,
	PosteriorPhoto:           &posteriorPhoto4,
	RadioOpacity:             &radioOpacity4,
	Sucralfate:               &sucrafate4,
	Validated:                &objectValidated4,
	Submitted:                &objectSubmitted4,
}

var updatedSwallowedObject2 = models.SwallowedObject{
	ID:                       &objectID4,
	IncidentID:               incident2.ID,
	AceticAcid:               &aceticAcid3,
	AnteriorLongestLength:    &anteriorLongestLength3,
	AnteriorPhoto:            &anteriorPhoto3,
	BatteryImprintCode:       &batteryImprintCode3,
	BatteryType:              &batteryType3,
	CustomBatteryType:        &customBatteryType3,
	CustomMagnetType:         &customMagnetType3,
	CustomMaterial:           &customMateral3,
	CustomMitigatingFeatures: customMitigatingFeatures3,
	DeviceType:               &deviceType3,
	Honey:                    &honey3,
	Imaging:                  &imaging3,
	IsBatteryOrMagnet:        &isBatteryOrMagnet3,
	LateralLongestLength:     &lateralLongestLength3,
	LateralPhoto:             &lateralPhoto3,
	MagnetType:               &magnetType3,
	Material:                 &material3,
	MitigatingFeatures:       mitigatingFeatures3,
	NegativePoleDirection:    &negativePoleDirection3,
	NumberOfPieces:           &numberOfPieces3,
	NumberOfThisObject:       &numberOfThisObject3,
	ObjectCustomShape:        &objectCustomShape3,
	ObjectDescription:        &objectDescription3,
	ObjectDimensionality:     &objectDimensionality3,
	ObjectIntact:             &objectIntact3,
	ObjectLocation:           &objectLocation3,
	ObjectShape:              &objectShape3,
	OtherCharacteristics:     otherCharacteristics3,
	PosteriorLongestLength:   &posteriorLongestLength3,
	PosteriorPhoto:           &posteriorPhoto3,
	RadioOpacity:             &radioOpacity3,
	Sucralfate:               &sucrafate3,
	Validated:                &objectValidated3,
	Submitted:                &objectSubmitted4,
}

var swallowedObject3 = models.SwallowedObject{
	ID:                       &objectID5,
	IncidentID:               incident2.ID,
	AceticAcid:               &aceticAcid4,
	AnteriorLongestLength:    &anteriorLongestLength4,
	AnteriorPhoto:            &anteriorPhoto4,
	BatteryImprintCode:       &batteryImprintCode4,
	BatteryType:              &batteryType4,
	CustomBatteryType:        &customBatteryType4,
	CustomMagnetType:         &customMagnetType4,
	CustomMaterial:           &customMateral4,
	CustomMitigatingFeatures: customMitigatingFeatures4,
	DeviceType:               &deviceType4,
	Honey:                    &honey4,
	Imaging:                  &imaging4,
	IsBatteryOrMagnet:        &isBatteryOrMagnet4,
	LateralLongestLength:     &lateralLongestLength4,
	LateralPhoto:             &lateralPhoto4,
	MagnetType:               &magnetType4,
	Material:                 &material4,
	MitigatingFeatures:       mitigatingFeatures4,
	NegativePoleDirection:    &negativePoleDirection4,
	NumberOfPieces:           &numberOfPieces4,
	NumberOfThisObject:       &numberOfThisObject4,
	ObjectCustomShape:        &objectCustomShape4,
	ObjectDescription:        &objectDescription4,
	ObjectDimensionality:     &objectDimensionality4,
	ObjectIntact:             &objectIntact4,
	ObjectLocation:           &objectLocation4,
	ObjectShape:              &objectShape4,
	OtherCharacteristics:     otherCharacteristics4,
	PosteriorLongestLength:   &posteriorLongestLength4,
	PosteriorPhoto:           &posteriorPhoto4,
	RadioOpacity:             &radioOpacity4,
	Sucralfate:               &sucrafate4,
	Validated:                &objectValidated4,
	Submitted:                &objectSubmitted4,
}

var swallowedObject4 = models.SwallowedObject{
	ID:                       &objectID6,
	IncidentID:               incident1.ID,
	AceticAcid:               &aceticAcid3,
	AnteriorLongestLength:    &anteriorLongestLength3,
	AnteriorPhoto:            &anteriorPhoto3,
	BatteryImprintCode:       &batteryImprintCode3,
	BatteryType:              &batteryType3,
	CustomBatteryType:        &customBatteryType3,
	CustomMagnetType:         &customMagnetType3,
	CustomMaterial:           &customMateral3,
	CustomMitigatingFeatures: customMitigatingFeatures3,
	DeviceType:               &deviceType3,
	Honey:                    &honey3,
	Imaging:                  &imaging3,
	IsBatteryOrMagnet:        &isBatteryOrMagnet3,
	LateralLongestLength:     &lateralLongestLength3,
	LateralPhoto:             &lateralPhoto3,
	MagnetType:               &magnetType3,
	Material:                 &material3,
	MitigatingFeatures:       mitigatingFeatures3,
	NegativePoleDirection:    &negativePoleDirection3,
	NumberOfPieces:           &numberOfPieces3,
	NumberOfThisObject:       &numberOfThisObject3,
	ObjectCustomShape:        &objectCustomShape3,
	ObjectDescription:        &objectDescription3,
	ObjectDimensionality:     &objectDimensionality3,
	ObjectIntact:             &objectIntact3,
	ObjectLocation:           &objectLocation3,
	ObjectShape:              &objectShape3,
	OtherCharacteristics:     otherCharacteristics3,
	PosteriorLongestLength:   &posteriorLongestLength3,
	PosteriorPhoto:           &posteriorPhoto3,
	RadioOpacity:             &radioOpacity3,
	Sucralfate:               &sucrafate3,
	Validated:                &objectValidated3,
	Submitted:                &objectSubmitted3,
}

var ENTIncident1 = models.ENTIncident{
	ID:                         &ENTIncidentID3,
	AgeMonths:                  &ageMonths3,
	AgeYears:                   &ageYears3,
	Anesthesia:                 &anesthesia3,
	Complications:              complications3,
	Country:                    &country3,
	CustomComplications:        customComplications3,
	CustomSymptoms:             customSymptoms3,
	DaysUntilRemoval:           &daysUntilRemoval3,
	DeviceType:                 &deviceType3,
	EaseOfRemoval:              &easeOfRemoval3,
	Gender:                     &gender3,
	HospitalStay:               &hospitalStay3,
	HoursUntilRemoval:          &hoursUntilRemoval3,
	IncidentDescription:        &incidentDescription3,
	MinutesUntilRemoval:        &minutesUntilRemoval3,
	OpenSurgery:                &openSurgery3,
	Prognosis:                  &prognosis3,
	RemovalStrategy:            &removalStrategy3,
	RemovalSetting:             &removalSetting3,
	SwallowedObjects:           []*models.SwallowedObject{&swallowedObject1, &swallowedObject4},
	SymptomSeverity:            &symptomSeverity3,
	Symptoms:                   symptoms3,
	WasIncidentLifeThreatening: &wasIncidentLifeThreatening3,
	Year:                       &year3,
	Validated:                  &incidentValidated3,
	Submitted:                  &incidentSubmitted3,
}

var ENTIncident2 = models.ENTIncident{
	ID:                         &ENTIncidentID4,
	AgeMonths:                  &ageMonths4,
	AgeYears:                   &ageYears4,
	Anesthesia:                 &anesthesia4,
	Complications:              complications4,
	Country:                    &country4,
	CustomComplications:        customComplications4,
	CustomSymptoms:             customSymptoms4,
	DaysUntilRemoval:           &daysUntilRemoval4,
	DeviceType:                 &deviceType4,
	EaseOfRemoval:              &easeOfRemoval4,
	Gender:                     &gender4,
	HospitalStay:               &hospitalStay4,
	HoursUntilRemoval:          &hoursUntilRemoval4,
	IncidentDescription:        &incidentDescription4,
	MinutesUntilRemoval:        &minutesUntilRemoval4,
	OpenSurgery:                &openSurgery4,
	Prognosis:                  &prognosis4,
	RemovalStrategy:            &removalStrategy4,
	RemovalSetting:             &removalSetting4,
	SwallowedObjects:           []*models.SwallowedObject{&swallowedObject2, &swallowedObject3},
	SymptomSeverity:            &symptomSeverity4,
	Symptoms:                   symptoms4,
	WasIncidentLifeThreatening: &wasIncidentLifeThreatening4,
	Year:                       &year4,
	Validated:                  &incidentValidated4,
	Submitted:                  &incidentSubmitted4,
}

var ENTIncident2Update = models.ENTIncident{
	ID:                         &ENTIncidentID4,
	AgeMonths:                  &ageMonths3,
	AgeYears:                   &ageYears3,
	Anesthesia:                 &anesthesia3,
	Complications:              complications3,
	Country:                    &country3,
	CustomComplications:        customComplications3,
	CustomSymptoms:             customSymptoms3,
	DaysUntilRemoval:           &daysUntilRemoval3,
	DeviceType:                 &deviceType3,
	EaseOfRemoval:              &easeOfRemoval3,
	Gender:                     &gender3,
	HospitalStay:               &hospitalStay3,
	HoursUntilRemoval:          &hoursUntilRemoval3,
	IncidentDescription:        &incidentDescription3,
	MinutesUntilRemoval:        &minutesUntilRemoval3,
	OpenSurgery:                &openSurgery3,
	Prognosis:                  &prognosis3,
	RemovalStrategy:            &removalStrategy3,
	RemovalSetting:             &removalSetting3,
	SwallowedObjects:           []*models.SwallowedObject{&updatedSwallowedObject2, &swallowedObject3},
	SymptomSeverity:            &symptomSeverity4,
	Symptoms:                   symptoms4,
	WasIncidentLifeThreatening: &wasIncidentLifeThreatening4,
	Year:                       &year4,
	Validated:                  &incidentValidated4,
	Submitted:                  &incidentSubmitted4,
}

func TestCreateENTIncidents(t *testing.T) {
	ctx := context.Background()
	incidentRequest := ENTIncident1
	_, ok := CreateENTIncident(ctx, &incidentRequest, testUserID)
	if !ok {
		t.Error("createENTincident failed")
	}

	incidentRequest = ENTIncident2
	_, ok = CreateENTIncident(ctx, &incidentRequest, testUserID)
	if !ok {
		t.Error("createENTincident failed")
	}
}

func TestGetENTIncidents(t *testing.T) {
	ctx := context.Background()
	got, ok := GetENTIncidents(ctx, testUserID)
	if !ok {
		t.Error("getENTincidents failed")
	}
	want := models.GetENTIncidentsGoodResponse{
		Incidents: []*models.ENTIncident{&ENTIncident1, &ENTIncident2},
	}
	ok = reflect.DeepEqual(*got, want)
	if !ok {
		//	fmt.Println(render.Render(*got))
		//	fmt.Println(render.Render(want))
		// PrettyPrint(*got)
		// PrettyPrint(want)
		fmt.Println(render.Render(got.Incidents[0]))
		fmt.Println(render.Render(want.Incidents[0]))
		t.Error("ENTIncident got is not what we want")

	}

}

func TestUpdateENTIncident(t *testing.T) {
	ctx := context.Background()
	_, ok := UpdateENTIncident(ctx, ENTIncident2Update, testUserID)
	if !ok {
		t.Error("update ENTincident failed")
	}
	got, ok := GetENTIncidents(ctx, testUserID)
	if !ok {
		t.Error("getENTincidents failed")
	}
	want := models.GetENTIncidentsGoodResponse{
		Incidents: []*models.ENTIncident{&ENTIncident1, &ENTIncident2Update},
	}
	ok = reflect.DeepEqual(*got, want)
	if !ok {
		//	fmt.Println(render.Render(*got))
		//	fmt.Println(render.Render(want))
		// PrettyPrint(*got)
		// PrettyPrint(want)
		fmt.Println(render.Render(got.Incidents[1].SwallowedObjects))
		fmt.Println(render.Render(want.Incidents[1].SwallowedObjects))
		t.Error("ENTIncident got is not what we want")

	}
}

func TestDeleteENTIncident(t *testing.T) {
	ctx := context.Background()
	_, ok := DeleteENTIncident(ctx, *ENTIncident1.ID)
	if !ok {
		t.Error("delete ENTincident1 failed")
	}
	_, ok = DeleteENTIncident(ctx, *ENTIncident2.ID)
	if !ok {
		t.Error("delete ENTincident2 failed")
	}
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

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
var aceticAcid3 = "yes"
var aceticAcid4 = "yes"
var anteriorLongestLength3 = "40"
var anteriorLongestLength4 = "30"
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
var lateralLongestLength3 = ""
var lateralLongestLength4 = ""
var lateralPhoto3 = ""
var lateralPhoto4 = ""
var magnetType3 = ""
var magnetType4 = ""
var material3 = ""
var material4 = ""
var mitigatingFeatures3 = []string{"feature5", "feature6"}
var mitigatingFeatures4 = []string{"feature7", "feature8"}
var negativePoleDirection3 = ""
var negativePoleDirection4 = ""
var numberOfPieces3 = ""
var numberOfPieces4 = ""
var numberOfThisObject3 = ""
var numberOfThisObject4 = ""
var objectCustomShape3 = ""
var objectCustomShape4 = ""
var objectDescription3 = ""
var objectDescription4 = ""
var objectDimensionality3 = ""
var objectDimensionality4 = ""
var objectIntact3 = ""
var objectIntact4 = ""
var objectLocation3 = ""
var objectLocation4 = ""
var objectShape3 = ""
var objectShape4 = ""
var otherCharacteristics3 = []string{"char1", "char2"}
var otherCharacteristics4 = []string{"char3", "char4"}
var posteriorLongestLength3 = ""
var posteriorLongestLength4 = ""
var posteriorPhoto3 = ""
var posteriorPhoto4 = ""
var radioOpacity3 = ""
var radioOpacity4 = ""
var sucrafate3 = ""
var sucrafate4 = ""
var objectSubmitted3 = true
var objectSubmitted4 = false

var ENTIncidentID3 = "1234567890"
var ENTIncidentID4 = "1234567790"
var ageMonths3 = ""
var ageMonths4 = ""
var ageYears3 = ""
var ageYears4 = ""
var anesthesia3 = ""
var anesthesia4 = ""
var complications3 = []string{"vomiting", "choking"}
var complications4 = []string{"choking", "vomiting"}
var country3 = ""
var country4 = ""
var customComplications3 = []string{"coughing", "choking"}
var customComplications4 = []string{"vomiting", "coughing"}
var customSymptoms3 = []string{"vomiting", "choking"}
var customSymptoms4 = []string{"vomiting", "choking"}
var daysUntilRemoval3 int64 = 3
var daysUntilRemoval4 int64 = 4
var deviceType3 = ""
var deviceType4 = ""
var easeOfRemoval3 = ""
var easeOfRemoval4 = ""
var gender3 = ""
var gender4 = ""
var hospitalStay3 = ""
var hospitalStay4 = ""
var hoursUntilRemoval3 int64 = 3
var hoursUntilRemoval4 int64 = 6
var incidentDescription3 = ""
var incidentDescription4 = ""
var minutesUntilRemoval3 int64 = 7
var minutesUntilRemoval4 int64 = 8
var openSurgery3 = ""
var openSurgery4 = ""
var prognosis3 = ""
var prognosis4 = ""
var removalStrategy3 = ""
var removalStrategy4 = ""
var swallowedObjects3 = ""
var swallowedObjects4 = ""
var symptomSeverity3 = ""
var symptomSeverity4 = ""
var symptoms3 = []string{"sneezing", "choking"}
var symptoms4 = []string{"coughing", "sneezing"}
var wasIncidentLifeThreatening3 = ""
var wasIncidentLifeThreatening4 = ""
var year3 = ""
var year4 = ""
var incidentSubmitted3 = false
var incidentSubmitted4 = false

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
	Submitted:                objectSubmitted3,
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
	Submitted:                objectSubmitted4,
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
	SwallowedObjects:           []*models.SwallowedObject{&swallowedObject1},
	SymptomSeverity:            &symptomSeverity3,
	Symptoms:                   symptoms3,
	WasIncidentLifeThreatening: &wasIncidentLifeThreatening3,
	Year:                       &year3,
	Submitted:                  incidentSubmitted3,
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
	SwallowedObjects:           []*models.SwallowedObject{&swallowedObject2},
	SymptomSeverity:            &symptomSeverity4,
	Symptoms:                   symptoms4,
	WasIncidentLifeThreatening: &wasIncidentLifeThreatening4,
	Year:                       &year4,
	Submitted:                  incidentSubmitted4,
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

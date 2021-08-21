package pg

import (

)

func TestCreateMedicalExpertNormally(t *testing.T) {
	ctx := context.Background{}

	expert := Expert{

	}
	payload, ok := CreateExpertNormally(ctx context.Context, expertRequestObject *models.Expert, userID string)
	
}
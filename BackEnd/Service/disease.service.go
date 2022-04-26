package Service

import (
	"errors"

	"KobokDNA.com/Models"
	"KobokDNA.com/Utils/GlobalVar"
	"KobokDNA.com/Utils/ValidatorInput"
	"go.mongodb.org/mongo-driver/bson"
)

func AddDisease(Disease *Models.Disease) error {

	if _, err := getDisease(&Disease.DiseaseName); err == nil {
		return errors.New("Disease already exist")
	}

	diseaseDNA := Disease.SequenceDNA
	if err := ValidatorInput.DNAValidator(diseaseDNA); err != nil {
		return err
	}

	_, err := GlobalVar.DiseaseCollection.InsertOne(GlobalVar.Ctx, Disease)
	return err
}

func getDisease(DiseaseName *string) (string, error) {
	var Disease *Models.Disease
	var DiseaseDNA string

	query := bson.D{bson.E{Key: "name_disease", Value: DiseaseName}}
	err := GlobalVar.DiseaseCollection.FindOne(GlobalVar.Ctx, query).Decode(&Disease)

	if err != nil {
		return DiseaseDNA, err
	}

	DiseaseDNA = Disease.SequenceDNA
	return DiseaseDNA, err
}

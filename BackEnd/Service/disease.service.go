package Service

import (
	"errors"

	"KobokDNA.com/Helper/GlobalVar"
	"KobokDNA.com/Helper/ValidatorInput"
	"KobokDNA.com/Modules"
	"go.mongodb.org/mongo-driver/bson"
)

func AddDisease(Disease *Modules.Disease) error {

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
	var Disease *Modules.Disease
	var DiseaseDNA string

	query := bson.D{bson.E{Key: "name_disease", Value: DiseaseName}}
	err := GlobalVar.DiseaseCollection.FindOne(GlobalVar.Ctx, query).Decode(&Disease)

	if err != nil {
		return DiseaseDNA, err
	}

	DiseaseDNA = Disease.SequenceDNA
	return DiseaseDNA, err
}

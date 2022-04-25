package Service

import (
	"KobokDNA.com/Helper/GlobalVar"
	"KobokDNA.com/Modules"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func AllPrediction(diseaseName *string, date *string) ([]Modules.Result, error) {
	var results []Modules.Result
	filter := bson.D{bson.E{Key: "disease_name", Value: diseaseName}, bson.E{Key: "date", Value: date}}

	cursor, err := GlobalVar.ResultCollection.Find(GlobalVar.Ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(GlobalVar.Ctx) {
		var result Modules.Result
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(GlobalVar.Ctx)

	if len(results) == 0 {
		return nil, errors.New("There isn't prediction result data")
	}
	return results, nil
}

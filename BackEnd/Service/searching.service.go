package Service

import (
	"fmt"
	"strconv"
	"strings"

	"KobokDNA.com/Models"
	"KobokDNA.com/Utils/GlobalVar"
	"KobokDNA.com/Utils/ValidatorInput"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AllPrediction(query *string) ([]Models.Result, error) {
	var results []Models.Result
	var filter primitive.D

	date, diseaseName, err := parserQuery(*query)

	if diseaseName == "" {
		filter = bson.D{bson.E{Key: "date", Value: date}}
	} else if date == "" {
		filter = bson.D{bson.E{Key: "disease_name", Value: primitive.Regex{
			Pattern: diseaseName,
			Options: "i",
		}}}
	} else if date == "" && diseaseName == "" {
		filter = bson.D{{}}
	} else {
		filter = bson.D{bson.E{Key: "disease_name", Value: primitive.Regex{
			Pattern: diseaseName,
			Options: "i",
		}}, bson.E{Key: "date", Value: date}}
	}

	cursor, err := GlobalVar.ResultCollection.Find(GlobalVar.Ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(GlobalVar.Ctx) {
		var result Models.Result
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

/*
   Fungai parser untuk menghandle setiap hasil inputan pengguna

*/
func parserQuery(query string) (string, string, error) {
	types, err := ValidatorInput.SearchValidator(query)
	date := ""
	diseaseName := ""
	if err != nil {
		return date, diseaseName, err
	} else {
		switch types {
		case 1:
			str_date := strings.Split(query, "/")
			month, _ := strconv.Atoi(str_date[1])
			date = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
		case 2:
			splitQuery := strings.Split(query, " ")
			str_date := strings.Split(splitQuery[0], "/")
			month, _ := strconv.Atoi(str_date[1])
			date = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
			for i := 1; i < len(splitQuery); i++ {
				diseaseName += splitQuery[i]
				if i != len(splitQuery)-1 {
					diseaseName += " "
				}
			}

		case 3:
			str_date := strings.Split(query, "-")
			month, _ := strconv.Atoi(str_date[1])
			date = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
		case 4:
			splitQuery := strings.Split(query, " ")
			str_date := strings.Split(splitQuery[0], "-")
			month, _ := strconv.Atoi(str_date[1])
			date = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
			for i := 1; i < len(splitQuery); i++ {
				diseaseName += splitQuery[i]
				if i != len(splitQuery)-1 {
					diseaseName += " "
				}
			}

		case 5:
			str_date := strings.Split(query, " ")
			month, _ := strconv.Atoi(str_date[1])
			date = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
		case 6:
			splitQuery := strings.Split(query, " ")
			month, _ := strconv.Atoi(splitQuery[1])
			date = fmt.Sprintf("%s %s %s", splitQuery[0], GlobalVar.Bulan[month], splitQuery[2])
			for i := 3; i < len(splitQuery); i++ {
				diseaseName += splitQuery[i]
				if i != len(splitQuery)-1 {
					diseaseName += " "
				}
			}

		case 7:
			str_date := strings.Split(query, "/")
			date = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
		case 8:
			splitQuery := strings.Split(query, " ")
			str_date := strings.Split(splitQuery[0], "/")
			date = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
			for i := 1; i < len(splitQuery); i++ {
				diseaseName += splitQuery[i]
				if i != len(splitQuery)-1 {
					diseaseName += " "
				}
			}

		case 9:
			str_date := strings.Split(query, "-")
			date = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
		case 10:
			splitQuery := strings.Split(query, " ")
			str_date := strings.Split(splitQuery[0], "-")
			date = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
			for i := 1; i < len(splitQuery); i++ {
				diseaseName += splitQuery[i]
				if i != len(splitQuery)-1 {
					diseaseName += " "
				}
			}

		case 11:
			str_date := strings.Split(query, " ")
			date = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
		case 12:
			splitQuery := strings.Split(query, " ")
			date = fmt.Sprintf("%s %s %s", splitQuery[0], splitQuery[1], splitQuery[2])
			for i := 3; i < len(splitQuery); i++ {
				diseaseName += splitQuery[i]
				if i != len(splitQuery)-1 {
					diseaseName += " "
				}
			}
		case 13:
			diseaseName = query
		}
		return date, diseaseName, nil
	}
}

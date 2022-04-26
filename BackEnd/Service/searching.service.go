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

func AllPrediction(diseaseName *string, date *string) ([]Models.Result, error) {
	var results []Models.Result
	var filter primitive.D
	var date_in_database string

	valid_date := "DD/MM/YYYY or DD-MM-YYYY or DD MM YYYY"

	if *date != "" {
		str_date := strings.Split(*date, "-")
		if len(str_date) != 3 {
			str_date = strings.Split(*date, "/")
			if len(str_date) != 3 {
				str_date = strings.Split(*date, " ")
				if len(str_date) != 3 {
					return results, errors.New(fmt.Sprintf("Date format is not valid (use %s)", valid_date))
				}
			}
		}

		types, err := ValidatorInput.DateValidator(*date)
		if err != nil {
			return results, errors.New(fmt.Sprintf("%s (use %s)", err.Error(), valid_date))
		} else {
			switch types {
			// case "dd/mm/yyyy"
			case 1:
				month, _ := strconv.Atoi(str_date[1])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
			case 2:
				month, _ := strconv.Atoi(str_date[0])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], GlobalVar.Bulan[month], str_date[2])
			case 3:
				month, _ := strconv.Atoi(str_date[2])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], GlobalVar.Bulan[month], str_date[0])
			case 4:
				month, _ := strconv.Atoi(str_date[1])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[2], GlobalVar.Bulan[month], str_date[0])

			// case "dd-mm-yyyy"
			case 5:
				month, _ := strconv.Atoi(str_date[1])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
			case 6:
				month, _ := strconv.Atoi(str_date[0])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], GlobalVar.Bulan[month], str_date[2])
			case 7:
				month, _ := strconv.Atoi(str_date[2])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], GlobalVar.Bulan[month], str_date[0])
			case 8:
				month, _ := strconv.Atoi(str_date[1])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[2], GlobalVar.Bulan[month], str_date[0])

			// case "dd/mm/yyyy"
			case 9:
				month, _ := strconv.Atoi(str_date[1])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[0], GlobalVar.Bulan[month], str_date[2])
			case 10:
				month, _ := strconv.Atoi(str_date[0])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], GlobalVar.Bulan[month], str_date[2])
			case 11:
				month, _ := strconv.Atoi(str_date[2])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], GlobalVar.Bulan[month], str_date[0])
			case 12:
				month, _ := strconv.Atoi(str_date[1])
				date_in_database = fmt.Sprintf("%s %s %s", str_date[2], GlobalVar.Bulan[month], str_date[0])

			// case "dd/bulan/yyyy"
			case 13:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
			case 14:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], str_date[0], str_date[2])
			case 15:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], str_date[2], str_date[0])
			case 16:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[2], str_date[1], str_date[0])

			// case "dd-bulan-yyyy"
			case 17:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
			case 18:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], str_date[0], str_date[2])
			case 19:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], str_date[2], str_date[0])
			case 20:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[2], str_date[1], str_date[0])

			// case "dd bulan yyyy"
			case 21:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[0], str_date[1], str_date[2])
			case 22:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], str_date[0], str_date[2])
			case 23:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[1], str_date[2], str_date[0])
			case 24:
				date_in_database = fmt.Sprintf("%s %s %s", str_date[2], str_date[1], str_date[0])
			}
		}
	}

	if *diseaseName == "" {
		filter = bson.D{bson.E{Key: "date", Value: date_in_database}}
	} else if *date == "" {
		filter = bson.D{bson.E{Key: "disease_name", Value: diseaseName}}
	} else {
		filter = bson.D{bson.E{Key: "disease_name", Value: diseaseName}, bson.E{Key: "date", Value: date_in_database}}
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

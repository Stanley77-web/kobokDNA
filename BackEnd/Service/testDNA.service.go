package Service

import (
	// "fmt"
	// "time"

	"fmt"
	"time"

	"KobokDNA.com/Lib/Similarity"
	"KobokDNA.com/Lib/StringMatching"
	"KobokDNA.com/Models"
	"KobokDNA.com/Utils/GlobalVar"
	"KobokDNA.com/Utils/ValidatorInput"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDNA(TestData Models.TestData) (int, error) {

	// cari di database
	userName := TestData.UserName
	userDNA := TestData.DNASequence
	diseaseName := TestData.DiseaseName
	method := TestData.Method

	if err := ValidatorInput.DNAValidator(userDNA); err != nil {
		return 0, err
	}

	now := time.Now()
	dateNow := fmt.Sprintf("%d %s %d", now.Day(), GlobalVar.Bulan[int(now.Month())], now.Year())

	diagnosis := false
	// var closedMatch string
	similarityScore := 0
	DiseaseDNA, err := getDisease(&diseaseName)

	if err != nil {
		return 0, err
	}

	if method == "1" {
		diagnosis = StringMatching.KnuthMorrisPratt(userDNA, DiseaseDNA)
	} else {
		diagnosis = StringMatching.BoyerMoore(userDNA, DiseaseDNA)
	}

	if diagnosis {
		similarityScore = 100
	} else {
		similarityScore = Similarity.SimilarityScore(userDNA, DiseaseDNA)
		if similarityScore > 80 {
			diagnosis = true
		}
	}

	id := GlobalVar.Len + 1
	result := Models.Result{
		ID:              int(id),
		Date:            dateNow,
		UserName:        userName,
		DiseaseName:     diseaseName,
		SimilarityScore: similarityScore,
		Status:          diagnosis,
	}

	if err := postResult(&result); err != nil {
		return 0, err
	}

	GlobalVar.Len++

	return int(id), nil
}

func postResult(result *Models.Result) error {
	_, err := GlobalVar.ResultCollection.InsertOne(GlobalVar.Ctx, result)
	return err
}

func TestResult(ID *int) (*Models.Result, error) {
	var result *Models.Result

	filter := bson.D{bson.E{Key: "id", Value: ID}}
	err := GlobalVar.ResultCollection.FindOne(GlobalVar.Ctx, filter).Decode(&result)
	return result, err
}

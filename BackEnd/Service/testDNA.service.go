package Service

import (
	// "fmt"
	// "time"

	"fmt"
	"time"

	"KobokDNA.com/Helper/GlobalVar"
	"KobokDNA.com/Helper/ValidatorInput"
	"KobokDNA.com/Lib/Similarity"
	"KobokDNA.com/Lib/StringMatching"
	"KobokDNA.com/Modules"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDNA(TestData Modules.TestData) (int, error) {

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
		// distance := Similarity.LCSDistance(closedMatch, DNApenyakit, len(closedMatch), len(DNApenyakit))
		similarityScore = Similarity.SimilarityScore(userDNA, DiseaseDNA)
		fmt.Println("Similarity Score : ", similarityScore)
		if similarityScore > 80 {
			diagnosis = true
		}
	}

	id := GlobalVar.Len + 1
	result := Modules.Result{
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

func postResult(result *Modules.Result) error {
	_, err := GlobalVar.ResultCollection.InsertOne(GlobalVar.Ctx, result)
	return err
}

func TestResult(ID *int) (*Modules.Result, error) {
	var result *Modules.Result

	filter := bson.D{bson.E{Key: "id", Value: ID}}
	err := GlobalVar.ResultCollection.FindOne(GlobalVar.Ctx, filter).Decode(&result)

	return result, err
}

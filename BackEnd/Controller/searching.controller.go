package Controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"KobokDNA.com/Helper/GlobalVar"
	"KobokDNA.com/Helper/ValidatorInput"
	"KobokDNA.com/Service"
	"github.com/gin-gonic/gin"
)

func GetAllPrediction(ctx *gin.Context) {
	date := ctx.Query("date")
	diseaseName := ctx.Query("disease_name")
	valid_date := "DD/MM/YYYY or DD-MM-YYYY or DD MM YYYY"

	str_date := strings.Split(date, "-")
	if len(str_date) != 3 {
		str_date = strings.Split(date, "/")
		if len(str_date) != 3 {
			str_date = strings.Split(date, " ")
			if len(str_date) != 3 {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Date format is not valid (use %s)", valid_date)})
				return
			}
		}
	}

	date_in_database := ""

	types, errors := ValidatorInput.DateValidator(date)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s (use %s)", errors.Error(), valid_date)})
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

	fmt.Println(date_in_database)
	fmt.Println("Type : ", types)

	results, err := Service.AllPrediction(&diseaseName, &date_in_database)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"predictions": results,
	})
}

// func searchingDNARoutes(rg *gin.RouterGroup) {
// 	disease_route := rg.Group("/searching")
// 	disease_route.GET("/predictionResult", GetTestResult)
// }

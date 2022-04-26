package Controller

import (
	"fmt"
	"net/http"
	"strconv"

	"KobokDNA.com/Models"
	"KobokDNA.com/Service"
	"github.com/gin-gonic/gin"
)

func TestDNAController(ctx *gin.Context) {
	var TestData Models.TestData

	if err := ctx.ShouldBindJSON(&TestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ID, err := Service.TestDNA(TestData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "DNA sequence for this disease is not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success adding a result",
		"id":      ID,
	})
}

func GetTestResult(ctx *gin.Context) {
	ID, errors := strconv.Atoi(ctx.Query("id"))

	fmt.Println(ID)

	if errors != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"message": errors.Error()})
		return
	}

	result, err := Service.TestResult(&ID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"date":             result.Date,
		"user_name":        result.UserName,
		"disease_name":     result.DiseaseName,
		"similarity_score": result.SimilarityScore,
		"status":           result.Status,
	})
}

// func testDNARoutes(rg *gin.RouterGroup) {
// 	disease_route := rg.Group("/testDNA")
// 	disease_route.POST("/test", TestDNAController)
// 	disease_route.GET("/result/:name/:date", GetTestResult)
// }

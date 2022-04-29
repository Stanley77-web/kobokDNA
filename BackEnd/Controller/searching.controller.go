package Controller

import (
	"net/http"

	"KobokDNA.com/Service"
	"github.com/gin-gonic/gin"
)

func GetAllPrediction(ctx *gin.Context) {
	query := ctx.Query("query")

	results, err := Service.AllPrediction(&query)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"predictions": results,
	})
}

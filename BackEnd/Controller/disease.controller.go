package Controller

import (
	"net/http"

	"KobokDNA.com/Models"
	"KobokDNA.com/Service"
	"github.com/gin-gonic/gin"
)

func AddDiseaseController(ctx *gin.Context) {
	var disease Models.Disease
	err := ctx.ShouldBindJSON(&disease)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Semua field harus diisi"})
		return
	}

	err = Service.AddDisease(&disease)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success adding a disease"})
}

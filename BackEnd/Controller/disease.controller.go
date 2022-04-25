package Controller

import (
	"net/http"

	"KobokDNA.com/Modules"
	"KobokDNA.com/Service"
	"github.com/gin-gonic/gin"
)

func AddDiseaseController(ctx *gin.Context) {
	var disease Modules.Disease
	err := ctx.ShouldBindJSON(&disease)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = Service.AddDisease(&disease)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success adding a disease"})
}

// func DiseaseRoutes(rg *gin.RouterGroup) {
// 	disease_route := rg.Group("/disease")
// 	disease_route.POST("/add", AddDiseaseController)
// }

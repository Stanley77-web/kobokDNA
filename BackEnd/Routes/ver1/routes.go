package ver1

import (
	"KobokDNA.com/Controller"
	"KobokDNA.com/Utils/GlobalVar"
)

func Routes() {

	GlobalVar.Server.GET("/", Controller.RootPages)

	rg := GlobalVar.Server.Group("/v1")

	disease_route := rg.Group("/disease")
	disease_route.POST("/add", Controller.AddDiseaseController)

	testDNA_route := rg.Group("/testDNA")
	testDNA_route.POST("/test", Controller.TestDNAController)
	testDNA_route.GET("/result", Controller.GetTestResult)

	searching_route := rg.Group("/searching")
	searching_route.GET("/predictionResult", Controller.GetAllPrediction)
}

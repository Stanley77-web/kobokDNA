package main

import (
	// "KobokDNA.com/Models"
	// "KobokDNA.com/Handlers"

	"log"

	"KobokDNA.com/Routes/ver1"
	"KobokDNA.com/Utils/GlobalVar"
)

// "net/http"
// "KobokDNA.com/BackEnd/Handlers"
// "github.com/gin-contrib/cors"
// "github.com/gin-gonic/gin"

// var Bulan map[int]string = make(map[int]string { 1 : "Januari",
// 	2 : "Februari",
// 	3 : "Maret",
// 	4 : "April",
// 	5 : "Mei",
// 	6 : "Juni",
// 	7 : "Juli",
// 	8 : "Agustus",
// 	9 : "September",
// 	10 : "Oktober",
// 	11 : "November",
// 	12 : "Desember",
// }

func main() {
	err := GlobalVar.Init()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer GlobalVar.MongoDB.Disconnect(GlobalVar.Ctx)
	// var TestDNA DNA.TestDNAInput
	// router := gin.Default()
	// router.Use(cors.Default())
	ver1.Routes()
	// router.POST("/Penyakit", Handlers.PostPenyakit)
	// router.POST("/TestDNA", Handlers.PostTestDNA)
	// // fmt.Println("Test DNA :", TestDNA.NamaPengguna, TestDNA.NamaPenyakit, TestDNA.SequenceDNA, TestDNA.Method)
	// // router.GET("/GetTestDNA", Handlers.GetTestDNA)
	// router.POST("/Searching", Handlers.PostSearching)
	GlobalVar.Server.Run()
	// foundKMP, firstPosKMP, closedMatchKMP := Script.KnuthMorrisPratt("abacaabaccabacabaabb", "abacabb")
	// if foundKMP {
	// 	fmt.Printf("Found match at first index: %d\n", firstPosKMP)
	// } else {
	// 	fmt.Printf("No match found. ")
	// 	fmt.Printf("Closed match: %s. ", closedMatchKMP)
	// 	distance := Script.LCSDistance(closedMatchKMP, "abacadb", len(closedMatchKMP), len("abacabb"))
	// 	fmt.Printf("Similarity Value : %d\n", Script.Similarity(len("abacabb"), distance))
	// }
	// foundBM, firstPosBM, closedMatchBM := Script.BoyerMoore("abacaabaccabacabaabb", "abacabb")
	// if foundBM {
	// 	fmt.Printf("Found match at first index: %d\n", firstPosBM)
	// } else {
	// 	fmt.Printf("No match found. ")
	// 	fmt.Printf("Closed match: %s. ", closedMatchBM)
	// 	distance := Script.LCSDistance(closedMatchBM, "abacadb", len(closedMatchBM), len("abacabb"))
	// 	fmt.Printf("Similarity Value : %d\n", Script.Similarity(len("abacabb"), distance))
	// }
}

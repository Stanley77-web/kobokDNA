package main

import (
	"log"

	"KobokDNA.com/Routes/ver1"
	"KobokDNA.com/Utils/GlobalVar"
)

func main() {
	err := GlobalVar.Init()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer GlobalVar.MongoDB.Disconnect(GlobalVar.Ctx)
	ver1.Routes()
	GlobalVar.Server.Run()
}

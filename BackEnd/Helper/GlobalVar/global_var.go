package GlobalVar

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Server            *gin.Engine
	MongoDB           *mongo.Client
	DiseaseCollection *mongo.Collection
	ResultCollection  *mongo.Collection
	Ctx               context.Context
	Bulan             map[int]string
	err               error
	Len               int64
)

func Init() error {
	Bulan = map[int]string{
		1:  "Januari",
		2:  "Februari",
		3:  "Maret",
		4:  "April",
		5:  "Mei",
		6:  "Juni",
		7:  "Juli",
		8:  "Agustus",
		9:  "September",
		10: "Oktober",
		11: "November",
		12: "Desember",
	}

	Ctx = context.TODO()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	conn := options.Client().
		ApplyURI("mongodb+srv://Client:v4HqSWcUYRpnRJKl@cluster0.ckucv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)

	MongoDB, err = mongo.Connect(Ctx, conn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = MongoDB.Ping(Ctx, readpref.Primary())

	DiseaseCollection = MongoDB.Database("myFirstDatabase").Collection("disease")
	ResultCollection = MongoDB.Database("myFirstDatabase").Collection("result")

	Len, err = ResultCollection.CountDocuments(Ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	Server = gin.Default()
	Server.Use(cors.Default())
	return nil
}

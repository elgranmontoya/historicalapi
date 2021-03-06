package main

import (
	"github.com/adamhei/historicalapi/handlers"
	"github.com/adamhei/historicalapi/routes"
	"github.com/adamhei/historicaldata/models"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"time"
)

func main() {
	mgoDialInfo := &mgo.DialInfo{
		Addrs:    []string{models.DbUrl},
		Timeout:  1 * time.Hour,
		Database: models.AUTHDB,
		Username: models.USERNAME,
		Password: models.PASSWORD,
	}
	sesh, err := mgo.DialWithInfo(mgoDialInfo)
	defer sesh.Close()

	if err != nil {
		log.Println("Could not connect to DB")
		panic(err)
	}

	db := sesh.DB(models.DbName)

	appContext := &handlers.AppContext{Db: db}
	router := routes.NewRouter(appContext)

	log.Fatal(http.ListenAndServe(":8080", router))
}

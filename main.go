package main

import (
	"encoding/json"
	"log"
	"os"
	"restProject/restlayer"
	"restProject/serverlayer/dbtools"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

func main() {
	file, err := os.Open("serverlayer/configuration/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	conf := new(Configuration)
	json.NewDecoder(file).Decode(conf)
	dbtools.DbInit(conf.DriverName, conf.DataSourceName)
	restlayer.RestStart("127.0.0.1:8081")
}

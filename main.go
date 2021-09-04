package main

import (
	"connectionDB/Settings"
	"connectionDB/database"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"io/ioutil"
	"log"
)

func main() {

	bytes, err := ioutil.ReadFile("Settings/settings.json")

	if err != nil {
		log.Fatalf("Cannot read file %e", err)
	}
	var Database settings.SettingsDB
	err = json.Unmarshal(bytes, &Database)
	if err != nil {
		log.Fatalf("Error can't parse bytes %e", err)
	}

	fmt.Println(Database)

	dsn := pgx.ConnConfig{
		Host:                 "localhost",
		Port:                 5432,
		Database:             "BankingExamination",
		User:                 "postgres",
		Password:             "umedjana1",
		TLSConfig:            nil,
		UseFallbackTLS:       false,
		FallbackTLSConfig:    nil,
		Logger:               nil,
		LogLevel:             0,
		Dial:                 nil,
		RuntimeParams:        nil,
		OnNotice:             nil,
		CustomConnInfo:       nil,
		CustomCancel:         nil,
		PreferSimpleProtocol: false,
		TargetSessionAttrs:   "",
	}

	//urlDatabase := fmt.Sprintf("postgres://postgres:umedjana1@localhost:5432/BankingExamination")
	connect, err := pgx.Connect(dsn)
	if err != nil {
		fmt.Printf("Не удалось соединиться с БД %e", err)
	}

	database.DBinitialisation(connect)
}


	/*clients, err := models.ListOfClients(connect)
	if err != nil {
		fmt.Println("BAD")
	} else {
		fmt.Println(clients)
	}
	models.Payer(connect, 1,5 ,4)
	if err != nil {
		fmt.Println("cant do %a", err)
	}else{
		fmt.Println("shopiing status success")
	}*/
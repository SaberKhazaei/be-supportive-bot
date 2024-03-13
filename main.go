package main

import (
	"BeSupporterBot/controller"
	"BeSupporterBot/model"
	"fmt"
	baleApi "github.com/ghiac/bale-bot-api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const botToken = "1771038996:yy4kRET57XKJdmEsQH5At68RMpZ7t6DCxEH6OAmh"

func main() {
	// connect to postgres:
	dsn := "host=localhost user=balebot password=balebot dbname=balebot port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	database := model.NewDatabase(db)

	// Bale bot:
	bot, err := baleApi.NewBaleBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := baleApi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	connector := controller.NewConnector(database, bot)
	for update := range updates {
		err = connector.Handler(&update)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

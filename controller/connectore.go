package controller

import (
	"BeSupporterBot/model"
	baleAPI "github.com/ghiac/bale-bot-api"
)

type Connector struct {
	db  model.Database
	bot *baleAPI.BotAPI
}

type State string

var (
	EnterPhoneNumberState  State = "enterPhoneNumber"
	EnterNationalCodeState State = "enterNational"
	EnterBirthDayState     State = "enterBirthDay"
	LoginStat              State = "login"
)

func NewConnector(db model.Database, bot *baleAPI.BotAPI) *Connector {
	return &Connector{
		db:  db,
		bot: bot,
	}
}

func (con *Connector) Handler(updates *baleAPI.Update) error {
	// start:
	if updates.Message != nil {
		if updates.Message.Text == "/start" {
			return con.Start(updates.Message.Chat.ID)
		}
	}
	// login callback:
	if updates.CallbackQuery != nil {
		if updates.CallbackQuery.Data == "login" {
			return con.CheckLogin(updates.CallbackQuery.Message.Chat.ID)
		}
	}
	// get stat:
	return con.CheckStat(updates.Message.Chat.ID, updates.Message.Text)
}

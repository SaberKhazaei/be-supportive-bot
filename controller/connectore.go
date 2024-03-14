package controller

import (
	"BeSupporterBot/model"
	"BeSupporterBot/view"
	baleAPI "github.com/ghiac/bale-bot-api"
)

type Connector struct {
	db  model.Database
	bot *baleAPI.BotAPI
}

type State string

var (
	EnterPhoneNumberState  State = "enterPhoneNumber"
	EnterVerificationCode  State = "enterVerificationCode"
	EnterFirstName         State = "enterFirstName"
	EnterLastName          State = "enterLastName"
	EnterNationalCodeState State = "enterNational"
	EnterBirthDayState     State = "enterBirthDay"
	EnterJobTitle          State = "enterJobTitle"
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
			check, err := con.CheckLogin(updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			if !check {
				return nil
			}
			err = view.EnterNumber(con.bot, updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}

			err = con.db.UpdateStat(updates.CallbackQuery.Message.Chat.ID, "enterPhoneNumber")
			if err != nil {
				return err
			}
			return nil
		} else {
			updates.CallbackQuery.Message.Text = updates.CallbackQuery.Data
		}
	}

	var id int64
	var text string
	// get stat:
	if updates.Message != nil {
		id = updates.Message.Chat.ID
		text = updates.Message.Text
	} else if updates.CallbackQuery != nil {
		id = updates.CallbackQuery.Message.Chat.ID
		text = updates.CallbackQuery.Message.Text
	}
	err := con.CheckStat(id, text)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) CheckStat(id int64, message string) error {
	stat, err := con.db.GetUserState(id)
	if err != nil {
		return err
	}
	switch stat {
	case string(EnterPhoneNumberState):
		return con.ManageEnterPhoneNumberState(message, id)
	case string(EnterVerificationCode):
		return con.ManageEnterVerificationCode(message, id)
	case string(EnterFirstName):
		return con.ManageEnterFirstName(message, id)
	case string(EnterLastName):
		return con.ManageEnterLastName(message, id)
	case string(EnterNationalCodeState):
		return con.ManageEnterNationalCodeState(message, id)
	case string(EnterBirthDayState):
		return con.ManageEnterBirthDayState(message, id)
	case string(EnterJobTitle):
		return con.ManageEnterJobTitle(message, id)
	}
	return nil
}

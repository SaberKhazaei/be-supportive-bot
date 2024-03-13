package controller

import (
	"BeSupporterBot/view"
	"fmt"
	baleAPI "github.com/ghiac/bale-bot-api"
	"strconv"
	"time"
)

func (con *Connector) CheckStat(id int64, message string) error {
	stat, err := con.db.GetUserState(id)
	if err != nil {
		return err
	}
	switch stat {
	case string(EnterPhoneNumberState):
		userPhoneNumber, err := strconv.Atoi(message)
		if err != nil {
			return fmt.Errorf("error in converting the user number from string to int64, error: %v", err)
		}
		err = con.SendEnterNativeCode(int64(userPhoneNumber), id)
		if err != nil {
			return err
		}
	case string(EnterNationalCodeState):
		nationalCode, err := strconv.Atoi(message)
		if err != nil {
			return fmt.Errorf("error in converting the user NationalCode from string to int64, error: %v", err)
		}
		err = con.SendEnterBirthDate(int64(nationalCode), id)
		if err != nil {
			return err
		}
	case string(EnterBirthDayState):
		date, err := time.Parse("2006-01-02", message)
		if err != nil {
			con.sendNewMessage(id, "لطفا تاریخ را به درستی وارد نمایید")
			return nil
		}
		err = con.db.SetUserBirthDate(id, date)
		if err != nil {
			return err
		}
		// send message
		con.sendNewMessage(id, "ثبت نام شما با موفقیت انجام شد")
		err = con.Start(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (con *Connector) CheckLogin(id int64) error {
	check, err := con.db.CheckUser(id)
	if err != nil {
		return err
	}
	if !check {
		// add user.
		err := con.CreateUser(id)
		if err != nil {
			return err
		}
		return nil
	} else {
		// send message
		con.sendNewMessage(id, "شما قبلا ثبت نام کرده اید")
		err := con.Start(id)
		if err != nil {
			return err
		}
		return nil
	}
}

func (con *Connector) Start(id int64) error {
	err := view.ListOfServices(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) CreateUser(id int64) error {
	err := view.EnterNumber(con.bot, id)
	if err != nil {
		return err
	}
	err = con.db.AddUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) SendEnterNativeCode(userPhoneNumber int64, id int64) error {
	err := con.db.SetUserPhoneNumber(id, userPhoneNumber)
	if err != nil {
		return err
	}
	err = view.EnterNativeCode(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) SendEnterBirthDate(nationalCode, id int64) error {
	err := con.db.SetUserNationalCode(id, nationalCode)
	if err != nil {
		return err
	}
	err = view.EnterBirthDate(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) sendNewMessage(id int64, message string) {
	msg := baleAPI.NewMessage(id, message)
	con.bot.Send(msg)
}

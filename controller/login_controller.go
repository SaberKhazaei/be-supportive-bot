package controller

import (
	"BeSupporterBot/view"
	"fmt"
	baleAPI "github.com/ghiac/bale-bot-api"
	"strconv"
	"time"
)

func (con *Connector) ManageEnterJobTitle(message string, id int64) error {
	err := con.db.SetUserJobTitle(id, message)
	if err != nil {
		return err
	}
	userInfo, err := con.db.GetUserInfo(id)
	if err != nil {
		return err
	}

	resMessage, password, err := SendUserInformation(userInfo.FirstName, userInfo.LastName, userInfo.NationalCode, userInfo.BirthDate, userInfo.PhoneNumber, userInfo.VerificationCode, userInfo.VerificationToken)
	if resMessage != "" {
		con.sendNewMessage(id, resMessage)
		err = con.Start(id)
		if err != nil {
			return err
		}
		return nil
	}
	con.sendNewMessage(id, fmt.Sprintf("پسورد شما : %s \n  ثبت نام با موفقیت ثبت شد", password))
	err = con.Start(id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterBirthDayState(message string, id int64) error {
	date, err := time.Parse("2006-01-02", message)
	if err != nil {
		con.sendNewMessage(id, "لطفا تاریخ را به درستی وارد نمایید")
		return nil
	}
	err = con.db.SetUserBirthDate(id, date)
	if err != nil {
		return err
	}
	err = con.SendEnterJobTitle(id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterNationalCodeState(message string, id int64) error {
	nationalCode, err := strconv.Atoi(message)
	if err != nil {
		return fmt.Errorf("error in converting the user NationalCode from string to int64, error: %v", err)
	}
	if len(message) < 10 {
		con.sendNewMessage(id, "لطفا کد ملی خود رو به درستی وارد نمایید.")
		return nil
	}
	err = con.SendEnterBirthDate(int64(nationalCode), id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterLastName(message string, id int64) error {
	err := con.db.SetUserLastName(id, message)
	if err != nil {
		return err
	}

	err = con.SendEnterNativeCode(id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterFirstName(message string, id int64) error {
	err := con.db.SetUserFirstName(id, message)
	if err != nil {
		return err
	}

	err = con.SendEnterLastName(id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterPhoneNumberState(message string, id int64) error {
	userPhoneNumber, err := strconv.Atoi(message)
	if err != nil {
		return fmt.Errorf("error in converting the user phone number from string to int64, error: %v", err)
	}
	if len(message) < 11 {
		con.sendNewMessage(id, "لطفا شماره همراه خود رو به درستی وارد نمایید.")
		return nil
	}

	check, requestVerificationToken, err := Registry()
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("we can't find the request verification token, error: %v", err)
	}

	err = con.db.SetUserVerificationToken(requestVerificationToken, id)
	if err != nil {
		return err
	}

	check, resMessage, err := CheckUserLoginKomiteEmdad(int64(userPhoneNumber), requestVerificationToken)
	if err != nil {
		return err
	}

	if check {
		con.sendNewMessage(id, fmt.Sprintf("%v\n لطفا کد ارسال شده رو وارد نمایید", resMessage))
		err := con.db.SetUserPhoneNumber(id, int64(userPhoneNumber))
		if err != nil {
			return err
		}

		err = con.db.UpdateStat(id, "enterVerificationCode")
		if err != nil {
			return err
		}
		return nil
	} else {
		con.sendNewMessage(id, resMessage)
		err := con.Start(id)
		if err != nil {
			return err
		}
		return nil
	}
}

func (con *Connector) ManageEnterVerificationCode(message string, id int64) error {
	verificationToken, err := con.db.GetUserVerificationToken(id)
	if err != nil {
		return err
	}

	enteredCode, err := strconv.Atoi(message)
	if err != nil {
		return err
	}
	fmt.Printf("\n entered code by user: %v \n ", enteredCode)
	check, resMessage, err := CheckVerificationCode(int64(enteredCode), verificationToken)
	if err != nil {
		return err
	}
	if !check {
		con.sendNewMessage(id, resMessage)
		err = con.Start(id)
		if err != nil {
			return err
		}
		return nil
	}
	err = con.db.SetUserEnteredCodeByPhoneMessage(id, int64(enteredCode))
	if err != nil {
		return err
	}

	err = con.SendEnterFirstName(id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) CheckLogin(id int64) (bool, error) {
	check, err := con.db.CheckUser(id)
	if err != nil {
		return false, err
	}
	if !check {
		// add user.
		err := con.CreateUser(id)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
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

func (con *Connector) SendEnterFirstName(id int64) error {
	err := con.db.UpdateStat(id, "enterFirstName")
	if err != nil {
		return err
	}
	err = view.EnterFirstName(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) SendEnterLastName(id int64) error {
	err := con.db.UpdateStat(id, "enterLastName")
	if err != nil {
		return err
	}
	err = view.EnterLastName(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) SendEnterNativeCode(id int64) error {
	err := con.db.UpdateStat(id, "enterNational")
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

func (con *Connector) SendEnterJobTitle(id int64) error {
	err := view.EnterJobTitle(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) sendNewMessage(id int64, message string) {
	msg := baleAPI.NewMessage(id, message)
	con.bot.Send(msg)
}

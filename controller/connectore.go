package controller

import (
	"BeSupporterBot/model"
	"BeSupporterBot/view"
	"encoding/json"
	"fmt"
	baleAPI "github.com/ghiac/bale-bot-api"
	"strings"
)

type Connector struct {
	db  model.Database
	bot *baleAPI.BotAPI
}

type State string

var (
	EnterPhoneNumberState       State = "enterPhoneNumber"
	EnterVerificationCode       State = "enterVerificationCode"
	EnterFirstName              State = "enterFirstName"
	EnterLastName               State = "enterLastName"
	EnterNationalCodeState      State = "enterNational"
	EnterNationalCodeResetState State = "enterNationalCode(reset)"
	EnterPhoneNumberResetState  State = "enterPhoneNumber(reset)"
	EnterCaptchaResetState      State = "enterCaptcha(reset)"
	EnterNationalCodeLoginState State = "enterNationalCode(login)"
	EnterPasswordLoginState     State = "enterLoginPassword"
	EnterCaptchaLoginState      State = "enterCaptcha"
	EnterBirthDayState          State = "enterBirthDay"
	EnterJobTitle               State = "enterJobTitle"
	LoginStat                   State = "login"
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
		if updates.CallbackQuery.Data == "register" {
			err := con.CheckRegister(updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			err = view.EnterNativeCode(con.bot, updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}

			err = con.db.UpdateStat(updates.CallbackQuery.Message.Chat.ID, "enterPhoneNumber")
			if err != nil {
				return err
			}
			return nil
		} else if updates.CallbackQuery.Data == "login" {
			err := con.CheckRegister(updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			//if check == false {
			//	con.sendNewMessage(updates.CallbackQuery.Message.Chat.ID, fmt.Sprintf("شما ثبت نام نکرده اید \n لطفا از طریق کلید ثبت نام مراحل ثبت نام را طی نمایید."))
			//	return nil
			//}

			err = view.EnterNativeCode(con.bot, updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			err = con.db.UpdateStat(updates.CallbackQuery.Message.Chat.ID, "enterNationalCode(login)")
			if err != nil {
				return err
			}
			return nil
		} else if updates.CallbackQuery.Data == "reset" {
			err := con.CheckRegister(updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			jobId, err := GetJobId()
			if err != nil {
				return err
			}
			err = con.db.SetJobIdLoginCode(updates.CallbackQuery.Message.Chat.ID, jobId)
			if err != nil {
				return err
			}

			err = view.EnterNativeCode(con.bot, updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			err = con.db.UpdateStat(updates.CallbackQuery.Message.Chat.ID, "enterNationalCode(reset)")
			if err != nil {
				return err
			}
			return nil
		} else if updates.CallbackQuery.Data == "newChild" {
			cookie, verificationToken, err := con.db.GetUserSiteCookieAndVerificationToken(updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				con.sendNewMessage(updates.CallbackQuery.Message.Chat.ID, "خطا در پردازش")
				return err
			}
			fullNamesMap, err := GetChildren(cookie, verificationToken)
			if err != nil {
				con.sendNewMessage(updates.CallbackQuery.Message.Chat.ID, "خطا در پردازش")
				return err
			}
			var fullNames []string
			for k, _ := range fullNamesMap {
				fullNames = append(fullNames, k)
			}
			fullNamesMapMarshalled, err := json.Marshal(fullNamesMap)
			if err != nil {
				return err
			}
			err = con.db.SetRepresentedChildForUser(updates.CallbackQuery.Message.Chat.ID, fullNamesMapMarshalled)
			if err != nil {
				return err
			}
			err = view.SendChoseChildButton(fullNames, con.bot, updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
		} else if updates.CallbackQuery.Data == "history" {
			cookie, verificationToken, err := con.db.GetUserSiteCookieAndVerificationToken(updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				con.sendNewMessage(updates.CallbackQuery.Message.Chat.ID, "خطا در پردازش")
				return err
			}
			err = GetListOfMyChildren(cookie, verificationToken)
			if err != nil {
				return err
			}
			return nil
		} else if strings.Contains(updates.CallbackQuery.Data, "ChoosingChild") {
			childName := strings.Replace(updates.CallbackQuery.Data, "ChoosingChild", " ", -1)
			childName = strings.TrimSpace(childName)
			fullNameMap, err := con.db.GetRepresentedChildForUser(updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
			var fullNameInfo map[string]string
			var childId string
			var ok bool
			err = json.Unmarshal(fullNameMap, &fullNameInfo)
			if err != nil {
				return err
			}
			if childId, ok = fullNameInfo[childName]; !ok {
				con.sendNewMessage(updates.CallbackQuery.Message.Chat.ID, "خطا در پردازش")
				return fmt.Errorf("error in get the child id by the child name")
			}
			fmt.Printf("name: %v \n", childName)
			fmt.Printf("id: %v \n", childId)
			currentChildren := make(map[string]string)
			currentChildren[childName] = childId
			currentChildrenMarshalled, err := json.Marshal(currentChildren)

			err = con.db.SetCurrentChildForUser(updates.CallbackQuery.Message.Chat.ID, currentChildrenMarshalled)
			if err != nil {
				return err
			}
			err = view.SendChoseChildPayment(con.bot, updates.CallbackQuery.Message.Chat.ID)
			if err != nil {
				return err
			}
		} else if strings.Contains(updates.CallbackQuery.Data, "EnteredPriceForChildPay") {
			childName := strings.Replace(updates.CallbackQuery.Data, "EnteredPriceForChildPay", " ", -1)
			childName = strings.TrimSpace(childName)

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
	fmt.Printf("state: %v\n", stat)
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
	case string(EnterNationalCodeLoginState):
		return con.ManageEnterNationalCodeLogin(message, id)
	case string(EnterPasswordLoginState):
		return con.ManageEnterLoginPassword(message, id)
	case string(EnterCaptchaLoginState):
		return con.ManageEnterCaptcha(message, id)
	case string(EnterNationalCodeResetState):
		return con.ManageEnterNationalCodeReset(message, id)
	case string(EnterPhoneNumberResetState):
		return con.ManageEnterPhoneNumberReset(message, id)
	case string(EnterCaptchaResetState):
		return con.ManageEnterCaptchaResetState(message, id)
	}
	return nil
}

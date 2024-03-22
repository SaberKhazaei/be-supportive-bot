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
	EnterPhoneNumberState         State = "enterPhoneNumber"
	EnterVerificationCode         State = "enterVerificationCode"
	EnterFirstName                State = "enterFirstName"
	EnterLastName                 State = "enterLastName"
	EnterNationalCodeState        State = "enterNational"
	EnterNationalCodeResetState   State = "enterNationalCode(reset)"
	EnterPhoneNumberResetState    State = "enterPhoneNumber(reset)"
	EnterCaptchaResetState        State = "enterCaptcha(reset)"
	ResetState                    State = "reset"
	EnterNationalCodeLoginState   State = "enterNationalCode(login)"
	EnterPasswordLoginState       State = "enterLoginPassword"
	EnterCaptchaLoginState        State = "enterCaptcha"
	EnterBirthDayState            State = "enterBirthDay"
	EnterJobTitle                 State = "enterJobTitle"
	EnterNumberMonthForSupporting State = "enterNumberOfMonthForSupporting"
	LoginState                    State = "login"
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

	// callbacks:
	if updates.CallbackQuery != nil {
		CheckContinue, err := con.ManageCallBacks(updates.CallbackQuery.Data, updates.CallbackQuery.Message.Chat.ID)
		if err != nil {
			return err
		}
		if !CheckContinue {
			return nil
		}
	}

	//manage states:
	var id int64
	var text string
	if updates.Message != nil {
		id = updates.Message.Chat.ID
		text = updates.Message.Text
	} else if updates.CallbackQuery != nil {
		id = updates.CallbackQuery.Message.Chat.ID
		text = updates.CallbackQuery.Message.Text
	}
	err := con.CheckState(id, text)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) CheckState(id int64, message string) error {
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
	case string(EnterNumberMonthForSupporting):
		err := con.CheckLogin(id)
		if err != nil {
			return err
		}
		return con.ManageEnterNumberOfMonthForSupporting(message, id)
	}
	return nil
}

func (con Connector) ManageCallBacks(data string, id int64) (bool, error) {
	if data == "register" {

		err := con.CheckRegister(id)
		if err != nil {
			return false, err
		}
		err = view.EnterNativeCode(con.bot, id)
		if err != nil {
			return false, err
		}

		err = con.db.UpdateStat(id, "enterPhoneNumber")
		if err != nil {
			return false, err
		}

		return false, nil
	} else if data == "login/register" {
		err := view.ListOfTheLoginAndRegisterService(con.bot, id)
		if err != nil {
			return false, err
		}
	} else if data == "login" {
		err := con.CheckRegister(id)
		if err != nil {
			return false, err
		}

		err = view.EnterNativeCode(con.bot, id)
		if err != nil {
			return false, err
		}

		err = con.db.UpdateStat(id, "enterNationalCode(login)")
		if err != nil {
			return false, err
		}

		return false, nil
	} else if data == "reset" {
		err := con.CheckRegister(id)
		if err != nil {
			return false, err
		}

		jobId, err := GetJobId()
		if err != nil {
			return false, err
		}
		err = con.db.SetJobIdLoginCode(id, jobId)
		if err != nil {
			return false, err
		}

		err = view.EnterNativeCode(con.bot, id)
		if err != nil {
			return false, err
		}
		err = con.db.UpdateStat(id, "enterNationalCode(reset)")
		if err != nil {
			return false, err
		}
		return false, nil
	} else if data == "newChild" {
		err := view.AskForChoseChildFiltering(con.bot, id)
		if err != nil {
			return false, err
		}
		return false, nil
	} else if strings.Contains(data, "state") {
		stateSplit := strings.Split(data, " ")
		stateId := strings.TrimSpace(stateSplit[1])
		siteCookie, _, err := con.db.GetUserSiteCookieAndVerificationToken(id)
		if err != nil {
			return false, err
		}

		citiesOfState, err := GetCityOfState(stateId, siteCookie)
		if err != nil {
			return false, err
		}

		err = view.ChoseCityForChildButton(citiesOfState, con.bot, id)
		if err != nil {
			return false, err
		}
		return false, nil
	} else if strings.Contains(data, "city") {
		citySplit := strings.Split(data, " ")
		cityId := strings.TrimSpace(citySplit[1])
		providerId := strings.TrimSpace(citySplit[3])
		cookie, verificationToken, err := con.db.GetUserSiteCookieAndVerificationToken(id)
		if err != nil {
			con.sendNewMessage(id, "خطا در پردازش")
			return false, err
		}
		fullNamesMap, err := GetChildren(cookie, verificationToken, providerId, cityId)
		if err != nil {
			con.sendNewMessage(id, "خطا در پردازش")
			return false, err
		}
		var fullNames []string
		for k, _ := range fullNamesMap {
			fullNames = append(fullNames, k)
		}

		fullNamesMapMarshalled, err := json.Marshal(fullNamesMap)
		if err != nil {
			return false, err
		}
		err = con.db.SetRepresentedChildForUser(id, fullNamesMapMarshalled)
		if err != nil {
			return false, err
		}
		err = view.SendChoseChildButton(fullNames, con.bot, id)
		if err != nil {
			return false, err
		}
	} else if data == "choseChildByFilter" {
		err := view.ChoseStateForChoseChild(con.bot, id)
		if err != nil {
			return false, err
		}
		return false, nil
	} else if data == "choseChildByEmdade" {
		err := con.CheckLogin(id)
		if err != nil {
			return false, err
		}

		cookie, verificationToken, err := con.db.GetUserSiteCookieAndVerificationToken(id)
		if err != nil {
			con.sendNewMessage(id, "خطا در پردازش")
			return false, err
		}

		fullNamesMap, err := GetChildren(cookie, verificationToken, "", "")
		if err != nil {
			con.sendNewMessage(id, "خطا در پردازش")
			return false, err
		}
		var fullNames []string
		for k, _ := range fullNamesMap {
			fullNames = append(fullNames, k)
		}

		fullNamesMapMarshalled, err := json.Marshal(fullNamesMap)
		if err != nil {
			return false, err
		}
		err = con.db.SetRepresentedChildForUser(id, fullNamesMapMarshalled)
		if err != nil {
			return false, err
		}
		err = view.SendChoseChildButton(fullNames, con.bot, id)
		if err != nil {
			return false, err
		}

		return false, nil
	} else if data == "childrenList" {
		err := con.CheckLogin(id)
		if err != nil {
			return false, err
		}

		cookie, verificationToken, err := con.db.GetUserSiteCookieAndVerificationToken(id)
		if err != nil {
			con.sendNewMessage(id, "خطا در پردازش")
			return false, err
		}
		err = GetListOfMyChildren(cookie, verificationToken)
		if err != nil {
			return false, err
		}

		return false, nil
	} else if strings.Contains(data, "ChoosingChild") {
		var fullNameInfo map[string]map[string]string
		var childInfo map[string]string
		var ok bool
		err := con.CheckLogin(id)
		if err != nil {
			return false, err
		}

		childName := strings.Replace(data, "ChoosingChild", " ", -1)
		childName = strings.TrimSpace(childName)
		fullNameMap, err := con.db.GetRepresentedChildForUser(id)
		if err != nil {
			return false, err
		}

		err = json.Unmarshal(fullNameMap, &fullNameInfo)
		if err != nil {
			return false, err
		}

		if childInfo, ok = fullNameInfo[childName]; !ok {
			con.sendNewMessage(id, "خطا در پردازش")
			return false, fmt.Errorf("error in get the child id by the child name")
		}

		//fmt.Printf("name: %v \n", childName)
		//fmt.Printf("id: %v \n", childInfo["id"])
		currentChildren := make(map[string]string)
		currentChildren["name"] = childName
		currentChildren["OrphanId"] = childInfo["OrphanId"]
		currentChildren["OrphanCodeMelli"] = childInfo["OrphanCodeMelli"]
		currentChildrenMarshalled, err := json.Marshal(currentChildren)

		err = con.db.SetCurrentChildForUser(id, currentChildrenMarshalled)
		if err != nil {
			return false, err
		}
		err = view.SendChoseChildPayment(con.bot, id)
		if err != nil {
			return false, err
		}

		return false, nil
	} else if strings.Contains(data, "EnteredPriceForChildPay") {
		err := con.CheckLogin(id)
		if err != nil {
			return false, err
		}

		PriceForOrphan := strings.Replace(data, "EnteredPriceForChildPay", " ", -1)
		PriceForOrphan = strings.Replace(data, "﷼", " ", -1)
		PriceForOrphan = strings.TrimSpace(PriceForOrphan)
		currentChild, err := con.db.GetCurrentChildForUser(id)
		if err != nil {
			return false, err
		}

		var currentChildInfo map[string]string
		err = json.Unmarshal(currentChild, &currentChildInfo)
		if err != nil {
			return false, err
		}
		currentChildInfo["AllowancesForOrphans"] = PriceForOrphan
		currentChildInfoMarshalled, err := json.Marshal(currentChildInfo)
		if err != nil {
			return false, err
		}
		err = con.db.SetCurrentChildForUser(id, currentChildInfoMarshalled)
		if err != nil {
			return false, err
		}

		con.sendNewMessage(id, "لطفا تعداد ماه برای جامی شدن فرزند انتخاب شده را ارسال نمایید \n از یک تا ۱۲ ماه را می توانید انتخاب نمایید")
		err = con.db.UpdateStat(id, "enterNumberOfMonthForSupporting")
		if err != nil {
			return false, err
		}

		return false, nil
	}
	return true, nil
}

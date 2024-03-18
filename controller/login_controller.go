package controller

import (
	"BeSupporterBot/view"
	"fmt"
	baleAPI "github.com/ghiac/bale-bot-api"
	"strings"
)

var JobInfo = map[string]string{
	"آتشنشان":       "7000",
	"آزاد":          "5032",
	"استاد دانشگاه": "5044",
	"بازنشسته":      "5044",
	"بسیج (حامی ایتام و محسنین)": "5048",
	"بنیاد غدیر":                 "5054",
	"پرستار":                     "5053",
	"پزشک":                       "5026",
	"حوضه قضایی":                 "8000",
	"خانه دار":                   "2008",
	"خبرنگار":                    "5040",
	"خلبان":                      "6000",
	"روانشناس":                   "5076",
	"روحانی":                     "5028",
	"صنوف":                       "5025",
	"کارگر":                      "5031",
	"کارمند":                     "5000",
	"کشاورز":                     "5075",
	"محصل":                       "5034",
	"مرکزنیکوکاری":               "5055",
	"معلم":                       "5038",
	"مهندس":                      "5027",
	"نظامی":                      "5041",
	"ورزشکار":                    "5030",
	"هنرمند":                     "5029",
	"سایر":                       "5037",
}

func (con *Connector) ManageEnterJobTitle(message string, id int64) error {
	var jobId string
	jobId, ok := JobInfo[message]
	if !ok {
		jobId = "5037"
	}

	err := con.db.SetUserJobId(id, jobId)
	if err != nil {
		return err
	}

	userInfo, err := con.db.GetUserInfo(id)
	if err != nil {
		return err
	}

	resMessage, password, err := SendUserInformation(userInfo.FirstName, userInfo.LastName, userInfo.NationalCode, userInfo.BirthDate, userInfo.PhoneNumber, userInfo.VerificationCode, userInfo.VerificationToken, jobId)
	if resMessage == " عملیات با موفقیت انجام شد" {
		con.sendNewMessage(id, fmt.Sprintf("پسورد شما : %s \n  ثبت نام با موفقیت ثبت شد", password))
		err = con.Start(id)
		if err != nil {
			return err
		}
	} else {
		if strings.Contains(resMessage, "کد ملی فوق قبلا با شماره") {
			err = con.db.DeleteUser(id)
			if err != nil {
				return err
			}
			con.sendNewMessage(id, fmt.Sprintf("%v \n لطفا با شماره ای که قبلا ثبت نام کرده اید وارد شوید", resMessage))

			err = con.Start(id)
			if err != nil {
				return err
			}
			return nil
		}
		con.sendNewMessage(id, resMessage)
		err = con.Start(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (con *Connector) ManageEnterBirthDayState(message string, id int64) error {
	err := con.db.SetUserBirthDate(id, message)
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
	if len(message) != 10 {
		con.sendNewMessage(id, "لطفا کد ملی خود رو به درستی وارد نمایید.")
		return nil
	}
	err := con.SendEnterNationalCode(message, id)
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
	var verificationToken string
	var err error

	if len(message) != 11 {
		con.sendNewMessage(id, "لطفا شماره همراه خود رو به درستی وارد نمایید.")
		return nil
	}

	if message != "" {
		check, requestVerificationToken, cookie, err := Registry()
		if err != nil {
			return err
		}
		if !check {
			return fmt.Errorf("we can't find the request verification token, error: %v", err)
		}

		err = con.db.SetUserVerificationToken(requestVerificationToken, id, cookie)
		if err != nil {
			return err
		}
	}

	check, resMessage, err := CheckUserLoginKomiteEmdad(message, verificationToken)
	if err != nil {
		return err
	}

	if check {
		con.sendNewMessage(id, fmt.Sprintf("%v\n لطفا کد ارسال شده رو وارد نمایید", resMessage))
		err := con.db.SetUserPhoneNumber(id, message, "enterVerificationCode")
		if err != nil {
			return err
		}

		err = con.db.UpdateStat(id, "enterVerificationCode")
		if err != nil {
			return err
		}
		return nil
	} else {
		if resMessage == "این شماره تلفن برای کاربر دیگری ثبت شده است، لطفا با شماره تلفن دیگری ثبت نام نمایید." {
			err := con.db.SetUserPhoneNumber(id, message, "enterPhoneNumber")
			if err != nil {
				return err
			}
		}
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

	fmt.Printf("user with id: %v entered code by user: %v \n ", id, message)
	check, resMessage, err := CheckVerificationCode(message, verificationToken)
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
	err = con.db.SetUserEnteredCodeByPhoneMessage(id, message)
	if err != nil {
		return err
	}

	err = con.SendEnterFirstName(id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterNationalCodeLogin(message string, id int64) error {
	fmt.Printf("\n national code: %v \n", message)
	if len(message) != 10 {
		con.sendNewMessage(id, "لطفا کد ملی خود رو به درستی وارد نمایید.")
		return nil
	}

	if message != "" {
		check, requestVerificationToken, cookie, err := Identity()
		if err != nil {
			return err
		}
		if !check {
			return fmt.Errorf("we can't find the request verification token, error: %v", err)
		}
		err = con.db.SetUserVerificationToken(requestVerificationToken, id, cookie)
		if err != nil {
			return err
		}
	}

	err := con.db.SetUserNationalCode(id, message, "enterLoginPassword")
	if err != nil {
		return err
	}
	err = con.db.UpdateStat(id, "enterLoginPassword")
	if err != nil {
		return err
	}
	err = view.EnterPassword(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterLoginPassword(message string, id int64) error {
	fmt.Printf("send captcha")
	err := con.db.SetUserPassword(id, message, "enterCaptcha")
	if err != nil {
		return err
	}
	usr, err := con.db.GetUserInfo(id)
	if err != nil {
		return err
	}
	imageCap, cookie, err := GetCaptcha(usr.SiteCookie)
	if err != nil {
		return err
	}
	fmt.Printf("cookie: %v\n", cookie)
	err = con.db.SetUserSiteCookie(id, cookie)
	if err != nil {
		return err
	}

	con.sendNewMessage(id, "کد امنیتی زیر رو وارد نمایید")
	_, err = con.bot.Send(baleAPI.NewPhotoUpload(id, baleAPI.FileBytes{Name: "image.png", Bytes: imageCap}))
	if err != nil {
		return err
	}

	return nil
}

func (con *Connector) ManageEnterCaptcha(message string, id int64) error {
	if len(message) != 6 {
		con.sendNewMessage(id, "لطفا کد امنیتی را به درستی وارد نمایید")
		return nil
	}
	err := con.db.SetUserCaptcha(id, message, "login")
	if err != nil {
		return err
	}
	userInfo, err := con.db.GetUserInfo(id)
	if err != nil {
		return err
	}
	_, cookie, err := GetIdentity(userInfo.NationalCode, userInfo.Password, userInfo.Captcha, userInfo.VerificationToken, userInfo.SiteCookie)
	if err != nil {
		con.sendNewMessage(id, "خطایی رخ داده است \n لطفا دوباره تلاش کنید")
		return err
	}
	err = con.db.SetUserSiteCookie(id, cookie)
	if err != nil {
		return err
	}
	con.sendNewMessage(id, "شما با موفقیت وارد شدید.")
	err = con.Start(id)
	if err != nil {
		return err
	}
	return nil
}
func (con *Connector) ManageEnterNationalCodeReset(message string, id int64) error {
	fmt.Printf("\n national code: %v \n", message)
	if len(message) != 10 {
		con.sendNewMessage(id, "لطفا کد ملی خود رو به درستی وارد نمایید.")
		return nil
	}

	err := con.db.SetUserNationalCode(id, message, "enterPhoneNumber(reset)")
	if err != nil {
		return err
	}
	err = view.EnterNumber(con.bot, id)
	if err != nil {
		return err
	}
	return nil
}

func (con *Connector) ManageEnterPhoneNumberReset(message string, id int64) error {
	if len(message) != 11 {
		con.sendNewMessage(id, "لطفا شماره همراه خود رو به درستی وارد نمایید.")
		return nil
	}
	err := con.db.SetUserPhoneNumber(id, message, "enterCaptcha(reset)")
	if err != nil {
		return err
	}

	check, requestVerificationToken, cookie, err := Identity()
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("we can't find the request verification token, error: %v", err)
	}
	err = con.db.SetUserVerificationToken(requestVerificationToken, id, cookie)
	if err != nil {
		return err
	}

	usr, err := con.db.GetUserInfo(id)
	if err != nil {
		return err
	}
	imageCap, cookie, err := GetCaptcha(usr.SiteCookie)
	if err != nil {
		return err
	}
	err = con.db.SetUserSiteCookie(id, cookie)
	if err != nil {
		return err
	}

	con.sendNewMessage(id, "کد امنیتی زیر رو وارد نمایید")
	_, err = con.bot.Send(baleAPI.NewPhotoUpload(id, baleAPI.FileBytes{Name: "image.png", Bytes: imageCap}))
	if err != nil {
		return err
	}
	return nil
}
func (con *Connector) ManageEnterCaptchaResetState(message string, id int64) error {
	if len(message) != 6 {
		con.sendNewMessage(id, "لطفا کد امنیتی را به درستی وارد نمایید")
		return nil
	}
	err := con.db.SetUserCaptcha(id, message, "login")
	if err != nil {
		return err
	}

	userInformation, err := con.db.GetUserInfo(id)
	if err != nil {
		return err
	}
	fmt.Println("TEST4")

	err = SendResetPasswordRequest(userInformation.SiteCookie, userInformation.VerificationToken, userInformation.NationalCode, userInformation.PhoneNumber, userInformation.Captcha, userInformation.JobIdLoginCode)
	if err != nil {
		return err
	} else {
		con.sendNewMessage(id, "پیامک برای شما ارسال شد.")
		err = con.Start(id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (con *Connector) CheckRegister(id int64) error {
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
	}
	return nil
}

func (con *Connector) CheckLogin(id int64) (bool, error) {
	check, err := con.db.CheckUser(id)
	if err != nil {
		return false, err
	}
	if !check {
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
	err := con.db.AddUser(id)
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

func (con *Connector) SendEnterNationalCode(nationalCode string, id int64) error {
	err := con.db.SetUserNationalCode(id, nationalCode, "enterBirthDay")
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

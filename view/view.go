package view

import (
	"fmt"
	baleApi "github.com/ghiac/bale-bot-api"
)

func ListOfServices(bot *baleApi.BotAPI, id int64) error {
	buttons := baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("ثبت نام/ورود", "login/register"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("سرپرستی فرزند جدید", "newChild"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("لیست فرزندان من", "childrenList"),
		),
	)

	message := baleApi.NewMessage(id, "لیست خدمات شامل موارد زیر است :")
	message.ReplyMarkup = buttons

	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list of services button with id: %v, error: %v", id, err)
	}
	return nil
}

func ListOfTheLoginAndRegisterService(bot *baleApi.BotAPI, id int64) error {
	buttons := baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("ثبت نام", "register"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("ورود", "login"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("بازیابی رمز عبور", "reset"),
		),
	)
	message := baleApi.NewMessage(id, fmt.Sprintf("شما می توانید با استفاده از 'ثبت نام' در سامانه اکرام یک حساب شخصی بسازید. \n اگر قبلا حساب شخصی خود را ایجاد کرده اید می توانید با استفاده از 'ورود' وارد حساب کاربری خود بشوید.\n در صورتی که رمز حساب کاربری خود را فراموش کرده اید می توانید با استفاده از 'بازیابی رمز عبور' رمز شخصی خود را تعویض کنید."))
	message.ReplyMarkup = buttons
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list of login services button with id: %v, error: %v", id, err)
	}
	return nil
}

func SendChoseChildButton(buttons []string, bot *baleApi.BotAPI, id int64) error {
	message := baleApi.NewMessage(id, "گزینش فرزندان برای حامی شدن:")
	//rows := baleApi.NewInlineKeyboardMarkup()
	//for i := 0; i < len(buttons); i++ {
	//	newRow := baleApi.NewInlineKeyboardRow(
	//		baleApi.NewInlineKeyboardButtonData(buttons[i], fmt.Sprintf("ChoosingChild %v", buttons[0])),
	//	)
	//	rows.InlineKeyboard[i] = newRow
	//}
	message.ReplyMarkup = baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(buttons[0], fmt.Sprintf("ChoosingChild %v", buttons[0])),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(buttons[1], fmt.Sprintf("ChoosingChild %v", buttons[1])),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(buttons[2], fmt.Sprintf("ChoosingChild %v", buttons[2])),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(buttons[3], fmt.Sprintf("ChoosingChild %v", buttons[3])),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(buttons[4], fmt.Sprintf("ChoosingChild %v", buttons[4])),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(buttons[5], fmt.Sprintf("ChoosingChild %v", buttons[5])),
		),
		//baleApi.NewInlineKeyboardRow(
		//	baleApi.NewInlineKeyboardButtonData(buttons[6], fmt.Sprintf("ChoosingChild %v", buttons[6])),
		//),
		//baleApi.NewInlineKeyboardRow(
		//	baleApi.NewInlineKeyboardButtonData(buttons[7], fmt.Sprintf("ChoosingChild %v", buttons[7])),
		//),
	)

	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list of services button with id: %v, error: %v", id, err)
	}
	return nil
}

func SendChoseChildPayment(bot *baleApi.BotAPI, id int64) error {
	message := baleApi.NewMessage(id, "مبلغ پرداختی برای هر فرزند :")

	message.ReplyMarkup = baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("۱۰۰۰۰۰ ﷼", fmt.Sprintf("EnteredPriceForChildPay %v", "100000")),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("۲۰۰۰۰۰ ﷼", fmt.Sprintf("EnteredPriceForChildPay %v", "200000")),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("۵۰۰۰۰۰ ﷼", fmt.Sprintf("EnteredPriceForChildPay %v", "500000")),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("۱۰۰۰۰۰۰ ﷼", fmt.Sprintf("EnteredPriceForChildPay %v", "1000000")),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("۲۰۰۰۰۰۰ ﷼", fmt.Sprintf("EnteredPriceForChildPay %v", "2000000")),
		),
	)

	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list of services button with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterNumber(bot *baleApi.BotAPI, id int64) error {
	message := baleApi.NewMessage(id, "شماره همراه خود را وارد نمایید")
	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your number message with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterFirstName(bot *baleApi.BotAPI, id int64) error {
	// Send the message.
	message := baleApi.NewMessage(id, "نام خود را وارد نمایید")
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your first name with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterLastName(bot *baleApi.BotAPI, id int64) error {
	// Send the message.
	message := baleApi.NewMessage(id, "نام خانوادگی خود را وارد نمایید")
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your last name message with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterNativeCode(bot *baleApi.BotAPI, id int64) error {
	// Send the message.
	message := baleApi.NewMessage(id, "کد ملی خود را وارد نمایید")
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your National Code message with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterPassword(bot *baleApi.BotAPI, id int64) error {
	// Send the message.
	message := baleApi.NewMessage(id, "پسورد خود را وارد نمایید")
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your password message with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterBirthDate(bot *baleApi.BotAPI, id int64) error {
	// Send the message.
	message := baleApi.NewMessage(id, fmt.Sprintf("تاریخ تولدتان را وارد نمایید: \n مانند : 24-03-1383"))
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your Birth Date message with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterJobTitle(bot *baleApi.BotAPI, id int64) error {
	buttons := baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("آتشنشان", "آتشنشان"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("آزاد", "آزاد"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("استاد دانشگاه", "استاد دانشگاه"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("بازنشسته", "بازنشسته"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("بسیج (حامی ایتام و محسنین)", "بسیج (حامی ایتام و محسنین)"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("بنیاد غدیر", "بنیاد غدیر"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("پرستار", "پرستار"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("پزشک", "پزشک"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("حوضه قضایی", "حوضه قضایی"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("خانه دار", "خانه دار"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("خبرنگار", "خبرنگار"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("خلبان", "خلبان"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("روانشناس", "روانشناس"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("روحانی", "روحانی"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("صنوف", "صنوف"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("کارگر", "کارگر"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("کارمند", "کارمند"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("کشاورز", "کشاورز"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("محصل", "محصل"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("مرکزنیکوکاری", "مرکزنیکوکاری"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("معلم", "معلم"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("مهندس", "مهندس"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("نظامی", "نظامی"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("ورزشکار", "ورزشکار"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("هنرمند", "هنرمند"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("سایر", "سایر"),
		),
	)
	// Send the message.
	message := baleApi.NewMessage(id, fmt.Sprintf("شفل خود را از گزینه های زیر انتخاب نمایید: "))
	message.ReplyMarkup = buttons
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your job title message with id: %v, error: %v", id, err)
	}
	return nil
}

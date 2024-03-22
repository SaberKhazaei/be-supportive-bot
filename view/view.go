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

func AskForChoseChildFiltering(bot *baleApi.BotAPI, id int64) error {
	text := "نیکوکار گرامی، این سامانه بر‌اساس اصل رعایت عدالت در توزیع مساعدت‌های حامیان بین فرزندان نیازمند طراحی شده‌است. \n لذا در نمایش فرزندان همواره اسامی بالای فهرست، محرومیت بیشتری نسبت به سطرهای بعدی خواهند داشت.\n چنانچه شما فیلتر را به انتخاب امداد تعیین نموده ، محرومترین و نیازمندترین فرزندان در سراسر کشور به شما نمایش داده خواهد شد.\n \n در صورتی که تمایل دارید فیلتر مورد نظر را اعمال کنید گزینه فیلتر را انتخاب نمایید در غیر اینصورت گزینه به انتخاب امداد را انتخاب نمایید."
	message := baleApi.NewMessage(id, text)
	message.ReplyMarkup = baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("فیلتر", "choseChildByFilter"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("به انتخاب امداد", "choseChildByEmdade"),
		),
	)
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send ask for chose child filtering button with id: %v, error: %v", id, err)
	}
	return nil
}

func ChoseStateForChoseChild(bot *baleApi.BotAPI, id int64) error {
	message := baleApi.NewMessage(id, "استان مورد نظر را انتخاب نمایید:")
	message.ReplyMarkup = baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("آ‌ذربایجان شرقی", "state 10"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("آ‌ذربایجان غربی", "state 11"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("اردبیل", "state 34"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("اصفهان", "state 12"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("البرز", "state 36"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("ایلام", "state 13"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("بوشهر", "state 15"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("تهران", "state 16"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("چهارمحال وبختیاری", "state 19"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("خراسان جنوبی", "state 39"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("خراسان رضوی", "state 17"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("خراسان شمالی", "state 40"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("خوزستان", "state 18"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("زنجان", "state 20"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("سمنان", "state 21"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("سیستان وبلوچستان", "state 22"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("فارس", "state 23"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("قم", "state 37"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("کردستان", "state 24"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("کرمان", "state 25"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("کرمانشاه", "state 14"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("کهگیلویه بویراحمد", "state 26"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("گلستان", "state 38"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("گیلان", "state 27"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("لرستان", "state 28"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("مازندران", "state 29"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("مرکزی", "state 30"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("هرمزگان", "state 31"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("همدان", "state 32"),
		),
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("یزد", "state 33"),
		),
	)
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list of the states button with id: %v, error: %v", id, err)
	}
	return nil
}

func ChoseCityForChildButton(cities []map[string]interface{}, bot *baleApi.BotAPI, id int64) error {
	message := baleApi.NewMessage(id, "شهر مورد نظر را انتخاب نمایید:")

	var rows [][]baleApi.InlineKeyboardButton
	for _, city := range cities {
		row := baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(city["CityName"].(string), fmt.Sprintf("city %v providerId: %v ", city["CityId"].(float64), city["ProvinceId"].(float64))),
		)
		rows = append(rows, row)
	}

	message.ReplyMarkup = baleApi.NewInlineKeyboardMarkup(rows...)
	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list of cities button with id: %v, error: %v", id, err)
	}
	return nil
}

func SendChoseChildButton(fullNames []string, bot *baleApi.BotAPI, id int64) error {
	message := baleApi.NewMessage(id, "گزینش فرزندان برای حامی شدن:")

	var rows [][]baleApi.InlineKeyboardButton
	for _, fullName := range fullNames {
		row := baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData(fullName, fmt.Sprintf("ChoosingChild %v", fullName)),
		)
		rows = append(rows, row)
	}

	message.ReplyMarkup = baleApi.NewInlineKeyboardMarkup(rows...)
	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list childrens for chosing button with id: %v, error: %v", id, err)
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

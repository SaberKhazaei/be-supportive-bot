package view

import (
	"fmt"
	baleApi "github.com/ghiac/bale-bot-api"
)

func ListOfServices(bot *baleApi.BotAPI, id int64) error {
	buttons := baleApi.NewInlineKeyboardMarkup(
		baleApi.NewInlineKeyboardRow(
			baleApi.NewInlineKeyboardButtonData("ثبت نام/ورود", "login"),
			baleApi.NewInlineKeyboardButtonData("سرپرستی فرزند جدید", "newChild"),
			baleApi.NewInlineKeyboardButtonData("پرداخت سرانه فرزندان", "pay"),
			baleApi.NewInlineKeyboardButtonData("مشاهده خیرات من", "history"),
		),
	)

	message := baleApi.NewMessage(id, "لیست خدمات")
	message.ReplyMarkup = buttons

	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send list of services button with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterNumber(bot *baleApi.BotAPI, id int64) error {
	message := baleApi.NewMessage(id, "شماره همراه خود را وارد نمایید :")
	// Send the message
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your number message with id: %v, error: %v", id, err)
	}
	return nil
}

func EnterNativeCode(bot *baleApi.BotAPI, id int64) error {
	// Send the message.
	message := baleApi.NewMessage(id, "کد ملی خود را وارد نمایید :")
	if _, err := bot.Send(message); err != nil {
		return fmt.Errorf("error in send enter your National Code message with id: %v, error: %v", id, err)
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

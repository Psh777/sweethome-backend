package telegram

import (
	"../../class/fito"
	"../../class/security"
	"../../db/postgres"
	"../../types"
	"../../webserver/handlers"
	"../config"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"strings"
)

var bot *tgbotapi.BotAPI
var chatIds []int64

func RunBot(myconfig types.MyConfig, ver string) {

	var err error
	bot, err = tgbotapi.NewBotAPI(myconfig.Env.TelegramBot)
	if err != nil {
		fmt.Println("Telegram connection refused", err)
		return
	}

	chatIds, err = postgres.GetChatID()
	if err != nil {
		return
	}

	SendMsgBot("Backend restart " + ver)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch strings.ToLower(update.Message.Text) {

		case "hi":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
			_, _ = bot.Send(msg)

		case "admin":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "fuck off")
			_, _ = bot.Send(msg)

		case "list":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", chatIds))
			_, _ = bot.Send(msg)

		case "help":
			var mess string
			mess = mess + "hi - check connect bot\n"
			mess = mess + "start {PASSWORD}\n"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mess)
			_, _ = bot.Send(msg)

		}

		fraza := strings.Split(update.Message.Text, " ")

		switch strings.ToLower(fraza[0]) {
		case "start":
			if len(fraza) > 1 {
				if fraza[1] == config.GetMyConfig().MainConfig.AdminPassword {
					//chatIds = append(chatIds, update.Message.Chat.ID)
					//chatIds = lib.RemoveDuplicatesInt(chatIds)
					fmt.Printf("ADD Telega Chat: %+v\n", update.Message.Chat.ID)
					_ = postgres.SetChatID(update.Message.Chat.ID)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Done! Your id: %v", update.Message.Chat.ID))
					_, _ = bot.Send(msg)
					chatIds, err = postgres.GetChatID()
					if err != nil {
						return
					}
				}
			}
		}

		if !CheckStart(update.Message.Chat.ID) {
			continue
		}

		switch strings.ToLower(fraza[0]) {

		case "sequrity":

			if fraza[1] == "on" {
				result, _ := security.SetOn()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				_, _ = bot.Send(msg)
			}
			if fraza[1] == "off" {
				result, _ := security.SetOff()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				_, _ = bot.Send(msg)
			}

		case "fito":

			if fraza[1] == "on" {
				result, _ := fito.SetOn()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				_, _ = bot.Send(msg)
			}
			if fraza[1] == "off" {
				result, _ := fito.SetOff()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
				_, _ = bot.Send(msg)
			}

		}

	}

}

func SendMsgBot(text string) {
	for i := 0; i < len(chatIds); i++ {
		msg := tgbotapi.NewMessage(chatIds[i], text)
		_, _ = bot.Send(msg)
	}
}

func CheckStart(id int64) bool {
	for i := 0; i < len(chatIds); i++ {
		if id == chatIds[i] {
			return true
		}
	}
	return false
}

func SendTest(w http.ResponseWriter) {
	SendMsgBot("Test")
	handlers.HandlerSuccess(w, "ok")
}

func ChangeStatus(status bool) string {
	if status == true {
		return "OK"
	}
	return "ER"
}

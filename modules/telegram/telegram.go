package telegram

import (
	"../../class/kassa"
	"../../db/postgres"
	"../../db/redis"
	"../../types"
	"../../webserver/handlers"
	"../config"
	"../lib"
	"../singleton"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var bot *tgbotapi.BotAPI
var chat_ids []int64
var temp_status map[string]bool

func RunBot(myconfig types.MyConfig) {

	temp_status = make(map[string]bool)
	temp_status["new"] = true
	temp_status["unpaid"] = true
	temp_status["pending"] = true
	temp_status["paid"] = true
	temp_status["work"] = true
	temp_status["done"] = true

	var err error
	bot, err = tgbotapi.NewBotAPI(myconfig.Env.TelegramBot)
	if err != nil {
		fmt.Println("Telegram connection refused")
		return
	}

	chat_ids = redis.GetChatID()

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
		case "status":

			var m map[string]types.Pong
			m = make(map[string]types.Pong)

			pings := singleton.GetPings()
			for i := 0; i < len(pings); i++ {
				ping := pings[i]
				m[ping.Source+"_"+ping.Target] = ping
			}

			var msgs string

			keys := make([]string, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}

			for i := 0; i < len(keys); i++ {
				b := m[keys[i]]
				d := time.Now().Unix() - b.Time.Unix()
				if b.Success == true {
					msgs = msgs + fmt.Sprintf("%v  /  %v  /  %v sec ago  /  %v\n", ChangeStatus(b.Success), b.Source+"-"+b.Target, d, b.Ping)
				} else {
					msgs = msgs + fmt.Sprintf("%v  /  %v  /  %v sec ago  /  %v\n", ChangeStatus(b.Success), b.Source+"-"+b.Target, d, b.Error.Error())
				}
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgs)
			_, _ = bot.Send(msg)

		case "hi":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
			_, _ = bot.Send(msg)

		case "admin":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "fuck off")
			_, _ = bot.Send(msg)

		case "list":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", chat_ids))
			_, _ = bot.Send(msg)

		case "help":
			var mess string
			mess = mess + "hi - check connect bot\n"
			mess = mess + "start {PASSWORD}\n"
			mess = mess + "orders {LIMIT} {PAGE}\n"
			mess = mess + "order {ORDER-ID}\n"
			mess = mess + "getfile {ORDER-ID} {FILE-ID} \n"
			mess = mess + "setstatus {ORDER-ID} {STATUS} \n"
			mess = mess + "setsprice {ORDER ID} {AMOUNT}\n"
			mess = mess + "setpayment {ORDER ID}\n"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mess)
			_, _ = bot.Send(msg)
		}

		fraza := strings.Split(update.Message.Text, " ")

		switch strings.ToLower(fraza[0]) {
		case "start":
			if len(fraza) > 1 {
				if fraza[1] == config.GetMyConfig().MainConfig.AdminPassword {
					chat_ids = append(chat_ids, update.Message.Chat.ID)
					chat_ids = lib.RemoveDuplicatesInt(chat_ids)
					fmt.Printf("Telega Chats: %+v\n", chat_ids)
					redis.SetChatID(chat_ids)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Done! Your id: %v", update.Message.Chat.ID))
					_, _ = bot.Send(msg)
				}
			}
		}

		if !CheckStart(update.Message.Chat.ID) {
			continue
		}

		switch strings.ToLower(fraza[0]) {
		case "file":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "eee "+fraza[1])
			_, _ = bot.Send(msg)

		case "getfile":
			if len(fraza) == 3 {
				filename, file_type, err := postgres.GetFileName(fraza[1], fraza[2])
				//f, err := os.Open("./files/" + filename)
				if err != nil {
					fmt.Println(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "File: "+filename+"")
				_, _ = bot.Send(msg)
				msg1 := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, config.GetMyConfig().Env.UploadFolder+"/"+fraza[1]+"_"+fraza[2]+"."+file_type)
				_, _ = bot.Send(msg1)
			} else {
				fmt.Println("Incorrect command", len(fraza))
				break
			}

		case "order":
			if len(fraza) > 1 {
				if fraza[1] == "" {
					break
				}
			} else {
				break
			}
			order, err := postgres.GetOrder(strings.ToUpper(fraza[1]))
			if err != nil {
				break
			}
			mess := PatternOrder(order)
			mess = mess + "\n"
			files, _ := postgres.ListFiles(strings.ToUpper(fraza[1]))
			for i := 0; i < len(files); i++ {
				ms := PatternFile(files[i])
				mess = mess + ms
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mess)
			_, _ = bot.Send(msg)

		case "orders":

			var out string
			limit := 10
			page := 1

			if len(fraza) == 3 {
				limit, err = strconv.Atoi(fraza[1])
				page, err = strconv.Atoi(fraza[2])

				if err != nil {
					break
				}
			}

			list, _ := postgres.ListOrders(limit, (page-1)*limit)
			for i := 0; i < len(list); i++ {
				mess := PatternOrder(list[i])
				out = out + mess
				out = out + "==========================================\n"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, out)
			_, _ = bot.Send(msg)

		case "setprice":

			if len(fraza) > 1 {
				err := postgres.SetPrice(fraza[1], fraza[2])
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
					_, _ = bot.Send(msg)
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Done")
				_, _ = bot.Send(msg)
			}

		case "setstatus":

			if len(fraza) > 1 {
				err := postgres.SetStatus(fraza[1], fraza[2])
				if !temp_status[fraza[2]] {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: status incorrect (new, unpaid, pending, paid, work, done)")
					_, _ = bot.Send(msg)
					continue
				}
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
					_, _ = bot.Send(msg)
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Done")
				_, _ = bot.Send(msg)
			}

		case "setpayment":

			order, err := postgres.GetOrder(fraza[1])

			if err != nil {
				fmt.Println("setpayment Error", err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: "+err.Error())
				_, _ = bot.Send(msg)
			}

			answer, err := kassa.CreatePayment(order.Price, order.ID, "Заказ N: "+order.ID, order.Email)
			if err != nil {
				fmt.Println("setpayment Error", err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: "+err.Error())
				_, _ = bot.Send(msg)
				continue
			}

			err = postgres.OrderUpdateAfterCreatePayment(order.ID, answer.ID, answer.Confirmation.ConfirmationUrl)
			if err != nil {
				fmt.Println("setpayment Error", err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error: "+err.Error())
				_, _ = bot.Send(msg)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ok, Payment ID: "+answer.ID)
			_, _ = bot.Send(msg)

		}
	}

}

func SendMsgBot(text string) {
	for i := 0; i < len(chat_ids); i++ {
		msg := tgbotapi.NewMessage(chat_ids[i], text)
		_, _ = bot.Send(msg)
	}
}

func CheckStart(id int64) bool {
	for i := 0; i < len(chat_ids); i++ {
		if id == chat_ids[i] {
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

func PatternOrder(item types.Order) string {
	var mess string
	mess = mess + "ID: " + item.ID + "\n"
	mess = mess + "Name: " + item.Name + "\n"
	mess = mess + "Email: " + item.Email + "\n"
	mess = mess + "Phone: " + item.Phone + "\n"
	mess = mess + "Lang: " + item.Lang + "\n"
	mess = mess + "Type: " + item.Type + "\n"
	mess = mess + "Payment: " + strconv.FormatBool(item.Payment) + "\n"
	mess = mess + "Payment Type: " + item.PaymentType + "\n"
	mess = mess + "Price: " + fmt.Sprintf("%v", item.Price) + " Rub\n"
	mess = mess + "Status: " + item.Status + "\n"
	mess = mess + "Date: " + fmt.Sprintf("%v", item.CreatedAt.Format("Mon Jan _2 15:04:05 2006")) + "\n"
	return mess
}

func PatternFile(item types.File) string {
	var mess string
	mess = mess + "File ID: " + fmt.Sprintf("%v", item.FileID) + "  /  "
	mess = mess + "File Name: " + item.FileName + "  /  "
	mess = mess + "Date: " + fmt.Sprintf("%v", item.CreatedAt.Format("Mon Jan _2 15:04:05 2006")) + "\n"
	return mess
}

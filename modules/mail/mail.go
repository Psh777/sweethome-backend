package mail

import (
	"../config"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var mail_template = map[string]map[string]string{}

var mail_from string

type Mailer struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Type    string `json:"type"`
	UserId  string `json:"user_id"`
}

func Send(from, to, subject, body string) {
	if from == "" {
		c := config.GetMyConfig().Env
		from = c.MailFrom
	}
	msg := Mailer{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	}
	//nats_client.Publish("mail", []byte(b))
	SendMail(msg)
}

func ReadFileMailTemplate() error {
	var err error
	err = ReadFileMailTemplateLang("ru")
	//err = ReadFileMailTemplateLang("en")
	return err
}

func ReadFileMailTemplateLang(lang string) error {
	//temp := make(map[string]string, 0)

	file, err1 := ioutil.ReadFile("./static/lang/mail/mail." + lang + ".json")
	if err1 != nil {
		fmt.Println("error : ", file, err1)
		return err1
	}
	c := make(map[string]string)
	//var c map[string]string
	err := json.Unmarshal(file, &c)
	if err != nil {
		return err
	}

	key_value := make(map[string]string, 0)
	for key, value := range c {
		//keys = append(keys, key)
		key_value[key + "_subject"] = value
		body, err := ioutil.ReadFile("./static/lang/mail/" + key +"."+ lang)
		if err != nil {
			fmt.Println("error : ", body, err)
			return err
		}
		key_value[key + "_body"] = string(body)
	}

	mail_template[lang] = key_value

	//fmt.Printf("GGG: %v\n", mail_template)

	conf := config.GetMyConfig().Env
	mail_from = conf.MailFrom

	return nil
}

func GetMailTemplate(code string, lang string) string {
	//if lang == "en" || lang == "ru" { } else {
	//	lang = "en"
	//}
	if mail_template[lang] == nil {
		lang = "en"
	}

	return mail_template[lang][code]
}

package mail

import (
	"../config"
	"../../types"
	"fmt"
	"strings"
)

func getSubjectBody(typer, lang string) (string, string){
	subject := GetMailTemplate(typer + "_subject", lang)
	body := GetMailTemplate(typer + "_body", lang)
	return subject, body
}

func TestEmail() {
	subject, body := getSubjectBody("test", "ru")
	Send(mail_from, "me@pp.je", subject, body)
}

func NewOrderAdmin(order types.Order) {
	subject, body := getSubjectBody("neworder_admin", "ru")
	c := config.GetMyConfig()
	subject = strings.Replace(subject, "{{order_id}}", order.ID, 100)
	body = strings.Replace(body, "{{project_name}}", c.MainConfig.ProjectName, 100)
	body = strings.Replace(body, "{{order_id}}", order.ID, 100)
	body = strings.Replace(body, "{{order_name}}", order.Name, 100)
	body = strings.Replace(body, "{{order_email}}", order.Email, 100)
	body = strings.Replace(body, "{{order_phone}}", order.Phone, 100)
	body = strings.Replace(body, "{{order_comment}}", order.Comment, 100)
	body = strings.Replace(body, "{{order_type}}", order.Type, 100)
	body = strings.Replace(body, "{{order_lang}}", order.Lang, 100)
	body = strings.Replace(body, "{{order_created}}", fmt.Sprintf("%v", order.CreatedAt), 100)
	Send(mail_from, "bukva@bukva.me", subject, body)
}

func NewOrderUser(order types.Order) {
	subject, body := getSubjectBody("neworder_user", "ru")
	c := config.GetMyConfig()
	subject = strings.Replace(subject, "{{order_id}}", order.ID, 100)
	subject = strings.Replace(subject, "{{project_url}}", c.MainConfig.ProjectUrl, 100)
	subject = strings.Replace(subject, "{{project_name}}", c.MainConfig.ProjectName, 100)

	body = strings.Replace(body, "{{project_name}}", c.MainConfig.ProjectName, 100)
	body = strings.Replace(body, "{{project_url}}", c.MainConfig.ProjectUrl, 100)
	body = strings.Replace(body, "{{order_id}}", order.ID, 100)
	body = strings.Replace(body, "{{order_name}}", order.Name, 100)
	body = strings.Replace(body, "{{order_email}}", order.Email, 100)
	body = strings.Replace(body, "{{order_phone}}", order.Phone, 100)
	body = strings.Replace(body, "{{order_comment}}", order.Comment, 100)
	body = strings.Replace(body, "{{order_type}}", order.Type, 100)
	body = strings.Replace(body, "{{order_lang}}", order.Lang, 100)
	body = strings.Replace(body, "{{order_created}}", fmt.Sprintf("%v", order.CreatedAt), 100)
	Send(mail_from, order.Email, subject, body)
}

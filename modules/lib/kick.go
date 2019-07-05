package lib

import (

	"regexp"
	"strconv"
	"fmt"
	"net/http"
)

func GetFormInt(r *http.Request, form string) int64 {

	pre_form := r.FormValue(form)
	form_out, _ := KickInt(pre_form)

	return form_out
}

func GetFormFloat(r *http.Request, form string) float64 {

	pre_form := r.FormValue(form)
	form_out, _ := KickFloat(pre_form)

	return form_out
}

func GetFormString(r *http.Request, form string) string {

	pre_form := r.FormValue(form)
	form_out, _ := KickString(pre_form)

	return form_out
}

func GetFormAddress(r *http.Request, form string) string {

	pre_form := r.FormValue(form)
	form_out, _ := KickAddress(pre_form)

	return form_out
}

func GetFormText(r *http.Request, form string) string {

	pre_form := r.FormValue(form)
	form_out, _ := KickText(pre_form)

	return form_out
}

func KickText(s string) (string, error) {

	r := regexp.MustCompile("[[\\p{L}\\d+_.,\n\r?!@#$%&*()\"<>=^/ -]+")

	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}

	return out, nil
}

func KickAddress(s string) (string, error) {

	r := regexp.MustCompile("[[\\p{L}\\d_.,/ -]+")

	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}

	return out, nil
}


func KickString(s string) (string, error) {

	r := regexp.MustCompile("[[\\p{L}\\d_-]+")

	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}

	return out, nil
}

func KickInt(s string) (int64, error) {

	r := regexp.MustCompile("[0-9]+")

	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}

	if out == "" { return 0, nil}

	out_int, err := strconv.ParseInt(out, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return out_int, nil

}

func KickFloat(s string) (float64, error) {

	r := regexp.MustCompile("[0-9].+")

	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}

	if out == "" { return 0, nil}

	out_int, err := strconv.ParseFloat(out,  64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return out_int, nil

}

func KickName(s string) string {

	r := regexp.MustCompile("[a-zA-Z\\s]+")

	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}

	return out

}

func KickStringAndNum(s string) bool {
	r := regexp.MustCompile("[A-Za-z0-9_,-]+")
	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}
	if s != out {
		return false
	}
	return true
}

func KickEmail(s string) bool {
	r := regexp.MustCompile("\\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,10}\\b")
	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}
	if s != out {
		return false
	}
	return true
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func IsPhone(s string) bool {
	r := regexp.MustCompile("[0-9 +-]+")
	var out string
	for _, match := range r.FindStringSubmatch(s) {
		out = match
	}
	//fmt.Println(s, out)
	if s != out {
		return false
	}
	return true
}
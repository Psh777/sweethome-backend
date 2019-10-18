package main

import (
	"./db/postgres"
	"./modules/config"
	"./modules/lang"
	"./modules/telegram"
	"./webserver"
	"fmt"
	"time"
)

func main() {

	fmt.Println("SWEET HOME SERVER v.0.0.5")

	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())

	//config
	myConfig := config.GetConfig("./static/env_config.json", "./static/main_config.json")

	//init
	postgres.InitX(myConfig.Env)

	err := lang.ReadFileCodeError()
	if err != nil {
		fmt.Println("ReadFileCodeError ERROR")
		panic(err)
	}

	go telegram.RunBot(myConfig)

	webserver.Init(myConfig)

}

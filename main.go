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
	ver := "v.0.0.10"
	fmt.Println("SWEET HOME SERVER " + ver)

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

	go telegram.RunBot(myConfig, ver)

	webserver.Init(myConfig)

}

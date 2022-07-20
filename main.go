package main

import (
	"fmt"
	"github.com/Psh777/sweethome-backend/db/postgres"
	"github.com/Psh777/sweethome-backend/modules/config"
	"github.com/Psh777/sweethome-backend/modules/lang"
	"github.com/Psh777/sweethome-backend/modules/telegram"
	"github.com/Psh777/sweethome-backend/webserver"
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

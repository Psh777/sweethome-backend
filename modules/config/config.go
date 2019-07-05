package config

import (
	"../../types"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

var projectConfig types.MyConfig

func GetConfig(envFile string, mainFile string) types.MyConfig {

	var myconfig types.MyConfig

	var commandLineConfigUrl = flag.String("c", "", "placeholder")
	flag.Parse()

	if len(*commandLineConfigUrl) == 0 {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		myconfig.Env = UnmarshalConfigTech(ReadFileConfig("", envFile))
		fmt.Println("conf path:", usr.HomeDir +envFile)
	} else {
		myconfig.Env = UnmarshalConfigTech(ReadFileConfig("", *commandLineConfigUrl))
		fmt.Println("conf path:", *commandLineConfigUrl)
	}

	myconfig.MainConfig = UnmarshalConfigMain(ReadFileConfig("", mainFile))
	projectConfig = myconfig

	return myconfig
}

func ReadFileConfig(homedir string, file string) []byte {
	readfile, err := ioutil.ReadFile(homedir + file)
	if err != nil {
		fmt.Println("ReadFileConfig error : ", homedir, err)
		os.Exit(1)
	}
	return readfile
}

func UnmarshalConfigMain(file []byte) types.MainConfig {
	var conf types.MainConfig
	err := json.Unmarshal(file, &conf)
	if err != nil {
		fmt.Println("UnmarshalConfigMain error config file")
		os.Exit(1)
	}
	return conf
}

func UnmarshalConfigTech(file []byte) types.Env {
	var conf types.Env
	err := json.Unmarshal(file, &conf)
	if err != nil {
		fmt.Println("UnmarshalConfigTech error config file")
		os.Exit(1)
	}
	return conf
}

func GetMyConfig() types.MyConfig {
	return projectConfig
}



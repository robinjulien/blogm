package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

// initInstance creates all the directories and file that a blogm instance need
func initInstance() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		file, err := os.Create("config.json")
		check(err)

		// Default config
		cfg := Config{
			Host: "localhost",
			Port: "8080",
		}

		jsoncfg, err := json.MarshalIndent(cfg, "", "	")

		file.Write(jsoncfg)

		err = file.Close()
		check(err)
	} else {
		// config file already exists, there is a chance that a website is already init here
		fmt.Println("A blogm instance is already init here.")
	}
}

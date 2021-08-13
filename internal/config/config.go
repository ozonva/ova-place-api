package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Host string `json:"host"`
}

func Load() {
	configuration := Config{}

	updateConfig := func(path string) error {
		file, err := os.Open(path)
		if err != nil {
			return errors.New("an error occurred while opening the file")
		}

		var deferErr error

		defer func(file *os.File) {
			deferErr = file.Close()
		}(file)

		decoder := json.NewDecoder(file)

		err = decoder.Decode(&configuration)
		if err != nil {
			return errors.New("an error occurred while decoding the json")
		}

		return deferErr
	}

	for {
		err := updateConfig("config/example.json")
		if err != nil {
			panic("An error occurred while updating the config")
		}
		fmt.Println(configuration.Host)
		time.Sleep(8 * time.Second)
	}
}

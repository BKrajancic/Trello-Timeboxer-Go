package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BoardID string
	AppKey  string
	Token   string
	Members []string
	Delays  map[string]float64
}

func GetConfig() (Config, error) {
	var config Config

	data, err := ioutil.ReadFile("config.json")
	if err == nil {
		err = json.Unmarshal(data, &config)
	}

	return config, err
}

package weather

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var LookIn = filepath.Join(os.Getenv("HOME"), ".weather.json")

const (
	OpenWeatherMapHost = "api.openweathermap.org"
	OpenWeatherMapPath = "data/2.5/weather"
)

type Configuration struct {
	Host    string `json:"host"`
	Path    string `json:"path"`
	Token   string `json:"token"`
	ZipCode string `json:"zip"`
	Units   string `json:"units"`
}

func LoadConfiguration(data []byte) (*Configuration, error) {
	config := &Configuration{
		Host: OpenWeatherMapHost,
		Path: OpenWeatherMapPath,
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}

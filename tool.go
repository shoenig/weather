package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	clean "github.com/hashicorp/go-cleanhttp"

	"gophers.dev/pkgs/ignore"
)

type Tool struct {
	configuration *Configuration
	httpClient    *http.Client
}

func NewTool(c *Configuration) *Tool {
	client := clean.DefaultClient()
	client.Timeout = 3 * time.Second
	return &Tool{
		configuration: c,
		httpClient:    client,
	}
}

func (t *Tool) Run() error {
	request, err := t.request()
	if err != nil {
		return err
	}

	client := clean.DefaultClient()
	client.Timeout = 10 * time.Second
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer ignore.Drain(response.Body)

	description, temp, err := t.process(response.Body)
	if err != nil {
		return err
	}

	t.output(description, temp)

	return nil
}

func (t *Tool) output(description string, temp float64) {
	unit := "°C"
	if t.configuration.Units == "imperial" {
		unit = "°F"
	}

	fmt.Printf("%s -- %3.1f %s\n", description, temp, unit)
}

type owmData struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func (t *Tool) process(r io.Reader) (string, float64, error) {
	var response owmData
	if err := json.NewDecoder(r).Decode(&response); err != nil {
		return "", 0, err
	}

	description := response.Weather[0].Description
	temp := response.Main.Temp

	return description, temp, nil
}

func (t *Tool) request() (*http.Request, error) {
	query := url.Values{
		"zip":   []string{t.configuration.ZipCode},
		"units": []string{t.configuration.Units},
		"APPID": []string{t.configuration.Token},
	}.Encode()

	rURL := (&url.URL{
		Scheme:   "https",
		User:     nil,
		Host:     t.configuration.Host,
		Path:     t.configuration.Path,
		RawQuery: query,
	}).String()

	request, err := http.NewRequest(http.MethodGet, rURL, nil)
	return request, err
}

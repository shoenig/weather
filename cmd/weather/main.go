package main

import (
	"io/ioutil"

	"gophers.dev/cmds/weather"
)

func main() {
	bs, err := ioutil.ReadFile(weather.LookIn)
	if err != nil {
		panic(err)
	}

	wConfig, err := weather.LoadConfiguration(bs)
	if err != nil {
		panic(err)
	}

	tool := weather.NewTool(wConfig)
	if err := tool.Run(); err != nil {
		panic(err)
	}
}

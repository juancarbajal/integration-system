package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type TAdapters struct {
	Adapters []TAdapter `json:"adapters"`
}

type TAdapter struct {
	Name  string `json:"name"`
	Delay int    `json:"delay"`
	count int
}

type TConfig struct {
	filename string
}

func (u TConfig) load() TAdapters {
	jsonFile, err := os.Open(u.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var adapters TAdapters
	json.Unmarshal(byteValue, &adapters)
	return adapters
}

func (a TAdapter) incCount() {
	a.count++
}
func (a TAdapter) isReady() bool {
	return a.count == a.Delay
}
func (a TAdapter) resetCount() {
	a.count = 0
}
func (a TAdapter) run() {

}

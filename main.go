package main

import (
	"fmt"
	"time"
)

func main() {
	config := TConfig{"config.json"}.load()
	for {
		for _, v := range config.Adapters {
			if v.isReady() {
				fmt.Println(v.Name)
				v.run()
				v.resetCount()
			} else {
				v.incCount()
			}
			time.Sleep(time.Second)
		}
	}
}

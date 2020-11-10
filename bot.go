package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ktu-telegram-notification-bot/pkg/notification"
	"ktu-telegram-notification-bot/pkg/scrapper"
	"log"
	"sync"
)

type Conf struct {
	Token    string
	Interval int
	Channel  string
}

func main() {
	dat, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatal("Error reading configuration file (conf.json).")
	}
	var conf Conf
	err = json.Unmarshal(dat, &conf)
	if err != nil {
		log.Fatal("Error parsing configuration parameters.")
	}
	var wg sync.WaitGroup
	c := make(chan scrapper.Notification)
	wg.Add(1)
	go notification.ListenAndRelayNotifications(c, &wg, conf.Interval)
	for i := range c {
		fmt.Println(i)
	}
	wg.Wait()
}

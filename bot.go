package main

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"ktu-telegram-notification-bot/pkg/notification"
	"ktu-telegram-notification-bot/pkg/scrapper"
	"ktu-telegram-notification-bot/pkg/telegram"
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
	for notification := range c {
		_, err := telegram.SendNotification(notification, conf.Token, conf.Channel)
		if err != nil {
			log.Print(err)
		}
	}
	wg.Wait()
}

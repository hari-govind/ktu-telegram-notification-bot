package main

import (
	"fmt"
	"ktu-telegram-notification-bot/pkg/notification"
	"ktu-telegram-notification-bot/pkg/scrapper"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan scrapper.Notification)
	wg.Add(1)
	go notification.ListenAndRelayNotifications(c, &wg)
	for i := range c {
		fmt.Println(i)
	}
	wg.Wait()
}

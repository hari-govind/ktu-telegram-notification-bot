package notification

import (
	"fmt"
	"ktu-telegram-notification-bot/pkg/scrapper"
)

func Print() {
	//Prints 5 most recent notifications
	fmt.Println(scrapper.ScrapeNotifications(5))
}

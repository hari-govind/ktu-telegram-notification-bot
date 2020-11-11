package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"ktu-telegram-notification-bot/pkg/scrapper"
	_ "log"
	"net/http"
	"strings"
)

type Message struct {
	Text      string `json:"text"`
	ChatID    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}

//Returns telegram bot API endpoint
func getTelegramAPI(token string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/", token)
}

func sendMessage(message Message, token string) (*http.Response, error) {
	messageJSON, _ := json.Marshal(message)
	url := getTelegramAPI(token)
	return http.Post(url+"sendMessage", "application/json", bytes.NewBuffer(messageJSON))
}

func SendNotification(notification scrapper.Notification, token string, channel string) (*http.Response, error) {
	title := strings.ToUpper(html.EscapeString(notification.Title))
	date := html.EscapeString(notification.Date)
	desc := html.EscapeString(notification.Desc)
	message := fmt.Sprintf("<b><u>%s</u></b>\n\n<i>%s</i>\n\n<pre>%s</pre>\n", title, date, desc)
	for _, link := range notification.Links {
		message += fmt.Sprintf("\n<a href=\"%s\">%s</a>", link.Url, link.Title)
	}
	return sendMessage(Message{message, channel, "HTML"}, token)
}

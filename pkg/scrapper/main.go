package scrapper

import (
	_ "errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	_ "log"
	"net/http"
	_ "os"
	"strings"
)

// Grabs KTU notification page as HTML
func GetHTML() (io.ReadCloser, error) {
	ktuURL := "https://ktu.edu.in/eu/core/announcements.htm"
	response, err := http.Get(ktuURL)
	if err != nil {
		return nil, err
	}
	return response.Body, err
}

func formatDate(date string) string {
	dateArray := strings.Fields(date)
	date = fmt.Sprintf("%s, %s %s %s", dateArray[0], dateArray[2], dateArray[1], dateArray[5])
	return date
}

type link struct {
	Url   string
	Title string
}

type Notification struct {
	Date  string
	Title string
	Desc  string
	Links []link
}

//Returns most recent n notifications(defined by number) in array
func ScrapeNotifications(number int) ([]Notification, error) {
	baseURL := "https://ktu.edu.in"
	body, err := GetHTML()
	if err != nil {
		return []Notification{Notification{}}, err
	}
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return []Notification{Notification{}}, err
	}
	notifications := make([]Notification, 0)
	doc.Find("tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i == number {
			return false
		}
		date := formatDate(strings.TrimSpace(s.Find("td strong").First().Text()))
		title := s.Find("li b").First().Text()
		desc := strings.TrimSpace(s.Find("li").Contents().Eq(3).Text())
		links := make([]link, 0)
		s.Find("a").Each(func(i int, li *goquery.Selection) {
			hrefVal, _ := li.Attr("href")
			if !strings.HasPrefix(hrefVal, "http") {
				hrefVal = baseURL + hrefVal
			}
			currentLink := link{hrefVal, li.Text()}
			links = append(links, currentLink)
		})
		notifications = append(notifications, Notification{date, title, desc, links})
		return true
	})
	return notifications, nil
}

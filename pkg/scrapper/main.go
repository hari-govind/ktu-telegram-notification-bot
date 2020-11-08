package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	_ "os"
	"strings"
)

// Grabs KTU notification page as HTML
func GetHTML() io.ReadCloser {
	ktuURL := "https://ktu.edu.in/eu/core/announcements.htm"
	response, err := http.Get(ktuURL)
	if err != nil {
		log.Fatal(err)
	}
	return response.Body
}

func formatDate(date string) string {
	dateArray := strings.Fields(date)
	date = fmt.Sprintf("%s, %s %s %s", dateArray[0], dateArray[2], dateArray[1], dateArray[5])
	return date
}

type link struct {
	url   string
	title string
}

type notification struct {
	date  string
	title string
	desc  string
	links []link
}

//Returns most recent n notifications(defined by number) in array
func ScrapeNotifications(number int) []notification {
	baseURL := "https://ktu.edu.in"
	body := GetHTML()
	defer body.Close()
	doc, _ := goquery.NewDocumentFromReader(body)
	notifications := make([]notification, 0)
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
		notifications = append(notifications, notification{date, title, desc, links})
		return true
	})
	return notifications
}

package notification

import (
	"crypto/md5"
	_ "fmt"
	badger "github.com/dgraph-io/badger/v2"
	"ktu-telegram-notification-bot/pkg/scrapper"
	"log"
	"sync"
	"time"
)

var dbFolder string = "data"

//Check if given notification is already relayed, if not then save it to database and return true
func isNewNotification(notification *scrapper.Notification, db *badger.DB) bool {
	txn := db.NewTransaction(true)
	key := md5.Sum([]byte(notification.Date + notification.Title))
	_, err := txn.Get(key[0:16])
	if err == nil {
		return false
	} else {
		err = txn.Set((key[0:16]), []byte("1"))
		if err != nil {
			log.Fatal(err)
		}
		txn.Commit()
		return true
	}
}

// Poll for new notifications from the official KTU announcements page
func ListenAndRelayNotifications(c chan scrapper.Notification, wg *sync.WaitGroup) {
	ticker := time.NewTicker(5 * time.Minute) //Minutes before each poll
	defer ticker.Stop()
	defer wg.Done()
	db, err := badger.Open(badger.DefaultOptions(dbFolder).WithLogger(nil))
	for _ = range ticker.C {
		notifications, err := scrapper.ScrapeNotifications(5)
		if err != nil {
			log.Print(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, notification := range notifications {
			if isNewNotification(&notification, db) {
				c <- notification
				time.Sleep(1 * time.Second)
			}
		}
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

}

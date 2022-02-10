package databaseservice

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dieptv"
	password = "dieptv"
	dbname   = "wallet_crypto"
)

var (
	Database *sql.DB
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

func ConnectDatabase() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	Database = db
}

func WaitForNotification() {
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(psqlInfo, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen("event_channel")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case n := <-listener.Notify:
			fmt.Println("Received data from channel [", n.Channel, "] :")
			// Prepare notification payload for pretty print
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
			if err != nil {
				fmt.Println("Error processing JSON: ", err)
				return
			}
			fmt.Println(string(prettyJSON.Bytes()))
			break
		case <-time.After(90 * time.Second):
			fmt.Println("Received no events for 90 seconds, checking connection")
			go func() {
				listener.Ping()
			}()
			break
		}
	}
}

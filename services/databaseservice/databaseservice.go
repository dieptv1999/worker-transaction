package databaseservice

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/lib/pq"
)

const (
	host     = "117.5.33.157"
	port     = 5432
	user     = "dieptv"
	password = "dieptv"
	dbname   = "wallet_crypto"
)

type Store struct {
	*Queries
	Db *sql.DB
}

var (
	store    *Store
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	once sync.Once
)

func GetDatabase() *Store {
	once.Do(func() {
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Println(err.Error(), "err.Error() services/databaseservice/databaseservice.go:29")
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			log.Println(err.Error(), "err.Error() services/databaseservice/databaseservice.go:34")
			panic(err)
		}

		fmt.Println("Successfully connected!")

		store = &Store{
			Db:      db,
			Queries: New(db),
		}
	})
	return store
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

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

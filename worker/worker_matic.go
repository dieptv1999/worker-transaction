package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"worker-transaction/appconfig"
	"worker-transaction/consts"
	"worker-transaction/listeners"
	"worker-transaction/services/databaseservice"

	"github.com/joho/godotenv"
)

var (
	startChan = make(chan bool, 1)
)

func main() {
	appconfig.LoadInitConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	defer func() {
		if databaseservice.GetDatabase() != nil {
			databaseservice.GetDatabase().Db.Close()
		}
	}()

	go func() {
		fmt.Println("listener")
		databaseservice.WaitForNotification()
	}()

	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-startChan:
				log.Println("=== WATCH POLYGON ===")
				time.Sleep(time.Millisecond)
				listeners.WatchEventErc20(consts.POLYGON_CHAIN, func() {
					startChan <- true
				})
			}
		}
	}(&wg)

	startChan <- true
	wg.Wait()
}

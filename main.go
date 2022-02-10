package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"sync"
	"time"
	"worker-transaction/consts"
	"worker-transaction/listeners"
	"worker-transaction/services/databaseservice"
)

var (
	startChan = make(chan bool, 1)
)

func main() {
	//appconfig.LoadInitConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseservice.ConnectDatabase()

	defer func() {
		if databaseservice.Database != nil {
			databaseservice.Database.Close()
		}
	}()

	go func() {
		fmt.Println("listener")
		databaseservice.WaitForNotification()
	}()

	for true {

	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-startChan:
				log.Println("=== WATCH BSC ===")
				time.Sleep(time.Millisecond)
				listeners.WatchEventErc20(consts.BSC_CHAIN, func() {
					startChan <- true
				})
			}
		}
	}(&wg)

	startChan <- true
	wg.Wait()
}

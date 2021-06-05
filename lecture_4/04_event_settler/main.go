package main

import (
	"log"
	"net/http"
	"time"

	"04_event_settler/api"
)

const ipAddress = "http://127.0.0.1"

const betsApiPort = ":8081"
const statusPath = "/bets?status="

const status = "active"

func main() {
	log.Print("Resolving all bets with status:" + status)
	httpClient := http.Client{Timeout: 15 * time.Second}
	// all values could be changed without impact on the code
	betsUrl := ipAddress + betsApiPort + statusPath + status
	betsForStatus, err := api.GetOddsWithStatus(httpClient, betsUrl)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(betsForStatus))

	err = api.ResolveBets(httpClient, betsForStatus)
	if err != nil{
		log.Fatal(err)
	}
	log.Print("All bets with status:" + status +  " are resolved")
}

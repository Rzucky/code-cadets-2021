package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed(
	httpClient http.Client,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)
	// repeatedly:
	for {
		select {
		// exit and close updates channel if context is finished
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			// get odds from HTTP server
			update, err := a.httpClient.Get(axilisFeedURL)
			if err != nil {
				log.Println("axilis offer feed, http get", err)
				continue
			}

			bodyContent, err := ioutil.ReadAll(update.Body)
			if err != nil {
				log.Println("axilis offer feed, json decode", err)
				continue
			}

			var feed []axilisOfferOdd
			err = json.Unmarshal(bodyContent, &feed)
			if err != nil {
				log.Println("axilis offer feed, json unmarshalling", err)
				continue
			}

			for _, x := range feed {
				// write data to updatesOdd then pass to updates channel
				updateOdd := models.Odd{
					Id:          x.Id,
					Name:        x.Name,
					Match:       x.Match,
					Coefficient: x.Details.Price,
					Timestamp:   time.Now(),
				}
				// in case it closes before we can send
				select {
				case <-ctx.Done():
					return nil
				case a.updates <- updateOdd:
					// do nothing
				}
			}
		}
	}
}

func (a *AxilisOfferFeed) String() string {
	return "axilis offer feed"
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}

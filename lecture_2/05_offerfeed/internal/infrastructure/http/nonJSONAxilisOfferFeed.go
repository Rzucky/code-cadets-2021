package http

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

const axilisFeedURL2 = "http://18.193.121.232/axilis-feed-2"

type NonJSONAxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewNonJSONAxilisOfferFeed(
	httpClient http.Client,
) *NonJSONAxilisOfferFeed {
	return &NonJSONAxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *NonJSONAxilisOfferFeed) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		// exit and close updates channel if context is finished
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			// get odds from HTTP server
			update, err := a.httpClient.Get(axilisFeedURL2)
			if err != nil {
				log.Println("axilis non json offer feed, http get", err)
				continue
			}
			bodyContent, err := ioutil.ReadAll(update.Body)
			if err != nil {
				log.Println("axilis non json offer feed, json decode", err)
				continue
			}
			textBodyContent := string(bodyContent)
			splittedBodyContent := strings.Split(textBodyContent, "\n")
			for _, x := range splittedBodyContent {
				updateOdd, err := parseArguments(x)
				if err != nil {
					log.Println("axilis non json offer feed, parsing data", err)
					break
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

func parseArguments(updateString string) (models.Odd, error) {
	attributes := strings.Split(updateString, ",")
	coefficient, err := strconv.ParseFloat(attributes[3], 64)
	if err != nil {
		return models.Odd{}, err
	}
	// write parsed data to updateOdd and return into channel
	updateOdd := models.Odd{
		Id:          attributes[0],
		Name:        attributes[1],
		Match:       attributes[2],
		Coefficient: coefficient,
		Timestamp:   time.Now(),
	}
	return updateOdd, nil
}

func (a *NonJSONAxilisOfferFeed) String() string {
	return "non JSON axilis offer feed"
}

func (a *NonJSONAxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

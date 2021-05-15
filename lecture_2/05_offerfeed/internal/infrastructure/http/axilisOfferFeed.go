package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"github.com/pkg/errors"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"
const axilisFeedURL2 = "http://18.193.121.232/axilis-feed-2"

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
	// repeatedly:
	for {

		select {
		case <-time.After(time.Second):
			// - get odds from HTTP server
			update, err1 := a.httpClient.Get(axilisFeedURL)
			//if available to get
			if err1 == nil {
				bodyContent, err := ioutil.ReadAll(update.Body)
				if err != nil {
					return errors.WithMessage(err, "reading body of update response")
				}
				var feed []axilisOfferOdd
				err = json.Unmarshal(bodyContent, &feed)
				if err != nil {
					return errors.WithMessage(err, "unmarshalling the JSON body content")
				}

				for _, x := range feed {
					//write data to updates channel
					a.updates <- models.Odd{
						Id:          x.Id,
						Name:        x.Name,
						Match:       x.Match,
						Coefficient: x.Details.Price,
						Timestamp:   time.Now(),
					}
				}
			} else {
				log.Print(err1, "URL not reachable")
			}

			//second URL updates
			update2, err2 := a.httpClient.Get(axilisFeedURL2)
			//if available to get
			if err2 == nil {
				bodyContent, err := ioutil.ReadAll(update2.Body)
				if err != nil {
					return errors.WithMessage(err, "reading body of update response")
				}

				textBodyContent := string(bodyContent)
				splittedBodyContent := strings.Split(textBodyContent, "\n")

				for _, x := range splittedBodyContent {
					attributes := strings.Split(x, ",")
					koef, _ := strconv.ParseFloat(attributes[3], 64)
					//write parsed data to updates channel
					a.updates <- models.Odd{
						Id:          attributes[0],
						Name:        attributes[1],
						Match:       attributes[2],
						Coefficient: koef,
						Timestamp:   time.Now(),
					}
				}
			} else {
				log.Print(err2, "URL not reachable")
			}
			if err1 != nil && err2 != nil{
				return errors.New("both URLs unreachable when getting offer feed")
			}

		//exit and close updates channel if context is finished
		case <-ctx.Done():
			close(a.updates)
			return nil
		}
	}
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

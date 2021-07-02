package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"04_event_settler/models"
)

const ipAddress = "http://127.0.0.1"
const eventApiPort = ":8080"
const eventPath = "/event/update"

// Arrays can't be constants https://stackoverflow.com/questions/13137463/declare-a-constant-array
var outcomes = [...]string{"won", "lost"}

func GetOddsWithStatus(client http.Client, url string) ([]models.BetDto, error) {
	response, err := client.Get(url)
	defer response.Body.Close()
	if err != nil {
		return []models.BetDto{}, errors.Wrap(err, "error getting data from bets API")
	}

	bodyContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []models.BetDto{}, errors.Wrap(err, "error reading response body")
	}

	var decodedBets []models.BetDto
	err = json.Unmarshal(bodyContent, &decodedBets)
	if err != nil {
		return []models.BetDto{}, errors.Wrap(err, "error unmarshalling response body")
	}

	return decodedBets, nil
}

func ResolveBets(client http.Client, bets []models.BetDto) error {
	// UnixNano() is recommended by https://golang.org/pkg/math/rand/
	// example: https://play.golang.org/p/8qOg760mLNi
	rand.Seed(time.Now().UnixNano())

	// set is used for storing all events from all bets, some bets have multiple events
	setForEvents := map[string]bool{}
	for _, bet := range bets {
		setForEvents[bet.SelectionId] = true
	}

	// resolve bets and send to event updates API
	for eventId := range setForEvents {
		update := models.EventUpdateDto{
				Id:      eventId,
				Outcome: outcomes[rand.Intn(len(outcomes))],
		}

		// printOutput is []byte
		marshalledUpdate, err := json.Marshal(update)
		if err != nil {
			return errors.WithMessage(err, "marshalling event updates data into JSON")
		}

		eventUpdatesUrl := ipAddress + eventApiPort + eventPath
		marshalledUpdateReader := bytes.NewReader(marshalledUpdate)
		_, err = client.Post(eventUpdatesUrl, "application/json", marshalledUpdateReader)
		if err != nil {
			return errors.WithMessage(err, "error sending POST request to update event.")
		}
	}

	return nil
}

package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"github.com/superbet-group/code-cadets-2021/lecture_4/03_bet_acceptance_api/internal/infrastructure/rabbitmq/models"
)

const contentTypeTextPlain = "text/plain"

// BetPublisher handles bet received queue publishing.
type BetPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewBetPublisher create a new instance of BetPublisher.
func NewBetPublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetPublisher {
	return &BetPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

// Publish publishes an received bet to the queue.
func (p *BetPublisher) Publish(bet models.BetDto) error {

	betJson, err := json.Marshal(bet)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        betJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", betJson)
	return nil
}

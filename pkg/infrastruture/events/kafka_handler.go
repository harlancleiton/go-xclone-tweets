package events

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
	domainEvents "github.com/harlancleiton/go-tweets/pkg/domain/events"
	"github.com/harlancleiton/go-tweets/pkg/kafka"
)

type eventJSON struct {
	ID         string      `json:"id"`
	EntityID   string      `json:"entity_id"`
	EntityName string      `json:"entity_name"`
	Name       string      `json:"name"`
	OcurredAt  time.Time   `json:"ocurred_at"`
	Payload    interface{} `json:"payload"`
}

type KafkaEventHandler struct {
	producer sarama.SyncProducer
}

func NewKafkaEventHandler(producer sarama.SyncProducer) *KafkaEventHandler {
	return &KafkaEventHandler{producer: producer}
}

func (h *KafkaEventHandler) Handle(event domainEvents.Event, wg *sync.WaitGroup) {
	log.Println("Handling event", event.Name())
	defer wg.Done()

	j := eventJSON{
		ID:         event.ID().String(),
		EntityID:   event.EntityID().String(),
		EntityName: event.EntityName(),
		Name:       event.Name(),
		OcurredAt:  event.OcurredAt(),
		Payload:    event.Payload(),
	}
	msg, err := json.Marshal(j)

	if err != nil {
		log.Println("Error marshalling event", err)
		return
	}

	_, _, err = kafka.SendMessage(h.producer, event.Name(), event.EntityName(), msg)

	if err != nil {
		log.Println("Error sending message to Kafka", err)
	}
}

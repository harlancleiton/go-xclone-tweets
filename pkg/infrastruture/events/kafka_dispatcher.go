package events

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/harlancleiton/go-tweets/pkg/domain/events"
	"github.com/harlancleiton/go-tweets/pkg/kafka"
)

type EventJSON struct {
	ID         string      `json:"id"`
	EntityID   string      `json:"entity_id"`
	EntityName string      `json:"entity_name"`
	Name       string      `json:"name"`
	OcurredAt  time.Time   `json:"ocurred_at"`
	Payload    interface{} `json:"payload"`
}

type KafkaEventDispatcher struct {
	producer  sarama.SyncProducer
	decoratee events.EventDispatcher
}

func (k *KafkaEventDispatcher) Dispatch(event events.Event) error {
	jsonAdapter := EventJSON{
		ID:         event.ID().String(),
		EntityID:   event.EntityID().String(),
		EntityName: event.EntityName(),
		Name:       event.Name(),
		OcurredAt:  event.OcurredAt(),
		Payload:    event.Payload(),
	}
	eventBytes, err := json.Marshal(jsonAdapter)

	if err != nil {
		return err
	}

	kafka.SendMessage(k.producer, event.Name(), event.EntityName(), eventBytes)

	return k.decoratee.Dispatch(event)
}

func (k *KafkaEventDispatcher) RegisterHandler(eventName string, handler events.EventHandler) error {
	return k.decoratee.RegisterHandler(eventName, handler)
}

func (k *KafkaEventDispatcher) UnregisterHandler(eventName string, handler events.EventHandler) error {
	return k.decoratee.UnregisterHandler(eventName, handler)
}

func (k *KafkaEventDispatcher) Has(eventName string, handler events.EventHandler) bool {
	return k.decoratee.Has(eventName, handler)
}

func (k *KafkaEventDispatcher) Clear() {
	k.decoratee.Clear()
}

func NewKafkaEventDispatcher(producer sarama.SyncProducer, decoratee events.EventDispatcher) *KafkaEventDispatcher {
	return &KafkaEventDispatcher{producer: producer, decoratee: decoratee}
}

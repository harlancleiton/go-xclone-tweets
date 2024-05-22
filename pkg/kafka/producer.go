package kafka

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

func NewSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "[Kafka] ", log.Ltime)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		return nil, err
	}

	return producer, nil
}

func SendMessage(producer sarama.SyncProducer, topic, key string, message []byte) (partition int32, offset int64, err error) {
	log.Printf("Sending message to Kafka topic %s, key %s, message %s", topic, key, message)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(message),
	}

	partition, offset, err = producer.SendMessage(msg)

	if err != nil {
		return 0, 0, err
	}

	return partition, offset, nil
}

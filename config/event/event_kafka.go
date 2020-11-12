package event

import "github.com/segmentio/kafka-go"

type kafkaReceiver struct {
	reader *kafka.Reader
}

func newKafkaEvent() *kafkaReceiver {
	return &kafkaReceiver{}
}

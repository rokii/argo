package event

import (
	"context"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var w = kafka.NewWriter(kafka.WriterConfig{
	Brokers:  []string{"rhs-pvrvkiaa-kfk-slc-1.rheos-streaming-qa.svc.33.tess.io:9092", "rhs-pvrvkiaa-kfk-slc-2.rheos-streaming-qa.svc.33.tess.io:9092"},
	Topic:    "senioric.reserved.test.demo",
	Balancer: &kafka.LeastBytes{},
})

// Send publishes event through kafka
func Send(key, value string) {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
	if err != nil {
		log.Errorf("error occurs while publishing msg to kafka, %v", err)
	}

}

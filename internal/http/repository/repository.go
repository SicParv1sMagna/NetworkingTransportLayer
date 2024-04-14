package repository

import (
	"github.com/IBM/sarama"
)

type Repository struct {
	producer *sarama.AsyncProducer
	consumer *sarama.Consumer
}

func New() (*Repository, error) {
	// config := sarama.NewConfig()
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.Retry.Max = 5
	// config.Producer.Return.Successes = true

	// producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	// if err != nil {
	// 	log.Fatalf("Error creating producer: %s", err)
	// }
	// defer producer.Close()

	// // config := sarama.NewConfig()

	return &Repository{
		producer: nil,
		consumer: nil,
	}, nil
}

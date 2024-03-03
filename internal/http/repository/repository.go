package repository

import "github.com/IBM/sarama"

type Repository struct {
	producer *sarama.AsyncProducer
	consumer *sarama.Consumer
}

func New() (*Repository, error) {
	return &Repository{
		producer: nil,
		consumer: nil,
	}, nil
}

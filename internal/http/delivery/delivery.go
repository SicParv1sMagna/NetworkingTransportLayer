package delivery

import (
	"log"

	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/http/usecase"
	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/kafka"
)

type Handler struct {
	UseCase  *usecase.UseCase
	Producer *kafka.Producer
}

func New(uc *usecase.UseCase) *Handler {
	producer, err := kafka.NewProducer()
	if err != nil {
		log.Fatal("Error occured while creating producer: ", err)
	}

	return &Handler{
		UseCase:  uc,
		Producer: producer,
	}
}

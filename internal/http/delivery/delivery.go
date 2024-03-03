package delivery

import "github.com/SicParv1sMagna/NetworkingTransportLayer/internal/http/usecase"

type Handler struct {
	UseCase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}

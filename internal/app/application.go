package app

import (
	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/http/delivery"
	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/http/repository"
	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/http/usecase"
)

type Application struct {
	repository *repository.Repository
	usecase    *usecase.UseCase
	handler    *delivery.Handler
}

func New() (*Application, error) {
	repo, err := repository.New()
	uc := usecase.New(repo)
	h := delivery.New(uc)

	if err != nil {
		return &Application{}, err
	}

	return &Application{
		repository: repo,
		usecase:    uc,
		handler:    h,
	}, nil
}

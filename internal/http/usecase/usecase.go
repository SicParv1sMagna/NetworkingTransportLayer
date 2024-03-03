package usecase

import "github.com/SicParv1sMagna/NetworkingTransportLayer/internal/http/repository"

type UseCase struct {
	Repository *repository.Repository
}

func New(r *repository.Repository) *UseCase {
	return &UseCase{
		Repository: r,
	}
}

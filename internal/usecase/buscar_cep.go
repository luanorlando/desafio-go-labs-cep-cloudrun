package usecase

import "github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/entity"

type CEPRepository interface {
	Fetch(cep string) (*entity.Cep, error)
}

type FetchCEPUsecase struct {
	repository CEPRepository
}

func NewFetchCEPUsecase(r CEPRepository) *FetchCEPUsecase {
	return &FetchCEPUsecase{
		repository: r,
	}
}

func (u FetchCEPUsecase) Execute(cep string) (*entity.Cep, error) {
	if !entity.ValidarCEP(cep) {
		return nil, entity.ErrCEPInvalid
	}

	return u.repository.Fetch(cep)
}

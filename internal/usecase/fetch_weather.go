package usecase

import "github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/entity"

type CEPRepository interface {
	Fetch(cep string) (*entity.Cep, error)
}

type FetchWeatherUsecase struct {
	repository CEPRepository
}

func NewFetchWeatherUsecase(r CEPRepository) *FetchWeatherUsecase {
	return &FetchWeatherUsecase{
		repository: r,
	}
}

func (u FetchWeatherUsecase) Execute(cep string) (*entity.Cep, error) {
	if !entity.ValidarCEP(cep) {
		return nil, entity.ErrCEPInvalid
	}

	return u.repository.Fetch(cep)
}

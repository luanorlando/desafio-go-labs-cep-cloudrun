package entity

import (
	"errors"
	"regexp"
)

type Weather struct {
	Celsius    float32 `json:"temp_C"`
	Fahrenheit float32 `json:"temp_F"`
	Kelvin     float32 `json:"temp_k"`
}

type Cep struct {
	City string `json:"localidade"`
	Erro string `json:"erro,omitempty"`
}

func ValidarCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return re.MatchString(cep)
}

var ErrCEPInvalid = errors.New("invalid zipcode")
var ErrCEPNotFound = errors.New("can not find zipcode")

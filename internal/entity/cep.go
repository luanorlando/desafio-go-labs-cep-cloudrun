package entity

import "regexp"

type Cep struct {
	city string `json:"localidade"`
}

func ValidarCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return re.MatchString(cep)
}

package webserver

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(params) != 1 || params[0] == "" {
		http.Error(w, "CEP não encontrado na url", http.StatusBadRequest)
		return
	}

	cep := params[0]

	if !validarCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	fmt.Println(cep)

}

func validarCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return re.MatchString(cep)
}

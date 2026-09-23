package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/entity"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/usecase"
)

type CEPHandler struct {
	usecase *usecase.FetchCEPUsecase
}

func NewHandler(u *usecase.FetchCEPUsecase) *CEPHandler {
	return &CEPHandler{
		usecase: u,
	}
}

func (h *CEPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// params := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	// if len(params) != 2 || params[0] == "" {
	// 	http.Error(w, "CEP não encontrado na url", http.StatusBadRequest)
	// 	return
	// }

	cep := r.PathValue("cep")

	result, err := h.usecase.Execute(cep)

	if err != nil {
		if errors.Is(err, entity.ErrCEPInvalid) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if errors.Is(err, entity.ErrCEPNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound) // <-- CORRIGIDO PARA 404
			return
		}

		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(result)
}

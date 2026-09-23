package main

import (
	"log"
	"net/http"

	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/handler"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/repository"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/usecase"
)

func main() {
	cepRepo := repository.NewViaCEPRepository()
	fetchCEPUsecase := usecase.NewFetchCEPUsecase(cepRepo)
	cepHandler := handler.NewHandler(fetchCEPUsecase)

	http.Handle("GET /cep/{cep}", cepHandler)

	log.Println("Servidor rodando na porta :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

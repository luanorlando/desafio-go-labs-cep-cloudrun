package main

import (
	"log"
	"net/http"

	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/handler"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/repository"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/usecase"
)

func main() {

	client := http.Client{}

	cepRepo := repository.NewViaCEPRepository(&client)
	weatherRepo := repository.NewExternalWeatherRepository("PEGAR APIKey", &client)
	usecase := usecase.NewFetchWeatherUsecase(cepRepo, weatherRepo)
	cepHandler := handler.NewHandler(usecase)

	http.Handle("GET /weather/{cep}", cepHandler)

	log.Println("Servidor rodando na porta :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

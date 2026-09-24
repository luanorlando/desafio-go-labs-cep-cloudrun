package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/config"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/handler"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/repository"
	"github.com/luanorlando/desafio-go-labs-cep-cloudrun.git/internal/usecase"
)

func main() {
	config, _ := config.LoadConfig(".")

	client := http.Client{}

	cepRepo := repository.NewViaCEPRepository(&client)
	weatherRepo := repository.NewExternalWeatherRepository(config.WeatherAPIKey, &client)
	usecase := usecase.NewFetchWeatherUsecase(cepRepo, weatherRepo)
	cepHandler := handler.NewHandler(usecase)

	http.Handle("GET /weather/{cep}", cepHandler)

	httpPort := fmt.Sprintf(":%s", config.HTTPPort)
	log.Println("Servidor rodando na porta %s", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

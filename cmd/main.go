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
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// PROTEÇÃO ANTI-PANIC: Garante que a struct existe na memória
	if config == nil {
		log.Fatal("Erro fatal: A estrutura de configuração retornou nula (nil)")
	}

	// Agora a linha 20 está segura contra Nil Pointer
	if config.WeatherAPIKey == "" {
		log.Fatal("A variável WEATHER_API_KEY é obrigatória")
	}

	client := http.Client{}

	cepRepo := repository.NewViaCEPRepository(&client)
	weatherRepo := repository.NewExternalWeatherRepository(config.WeatherAPIKey, &client)
	usecase := usecase.NewFetchWeatherUsecase(cepRepo, weatherRepo)
	cepHandler := handler.NewHandler(usecase)

	http.Handle("GET /weather/{cep}", cepHandler)

	httpPort := fmt.Sprintf(":%s", config.HTTPPort)
	log.Printf("Servidor iniciando na porta %s...", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

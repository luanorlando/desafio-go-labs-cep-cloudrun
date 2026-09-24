# Desafio: Clima por CEP no Google Cloud Run
Sistema em Go com deploy no Cloud Run — Go Expert

# Objetivo

Desenvolver um sistema em Go que receba um CEP, identifique a cidade correspondente e retorne o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O requisito final é que este sistema esteja publicado e acessível no Google Cloud Run.

# Setup

No Projeto tem um arquivo chamado `.env.example`, será necessário criar um arquivo `.env` e adicionar a apiKey gerada ao se cadastrar em [WeatherAPI](https://www.weatherapi.com/)

### Rodar os testes localmente via Docker

#### 1. Faz o build focado apenas no estágio dos testes
```shell
docker build --target builder -t desafio-testes .
```
#### 2. Executa o contêiner dos testes (ele rodará o 'go test ./... -v')
```shell
docker run --rm desafio-testes
```

### Rodar a Aplicação Localmente no Docker 

#### 1. Faz o build da imagem completa (vai gerar o executável final leve)
```shell
docker build -t desafio-app .
```
#### 2. Roda a aplicação injetando o seu arquivo .env local para dentro do contêiner
```shell
docker run --rm -p 8080:8080 --env-file .env desafio-app
```

# Requisitos funcionais

Entrada: O sistema deve receber um CEP válido de 8 dígitos.
Identificação de localização: O sistema deve realizar a busca do CEP para encontrar o nome da localização (cidade).
Consulta de clima: A partir da localização, o sistema deve consultar a temperatura atual.
Conversão: O sistema deve retornar as temperaturas formatadas em Celsius, Fahrenheit e Kelvin.

# Especificações da API (contrato)

Cenário 1: Sucesso — 200 OK
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

## Cenários de falha

|Cenário	Condição|	Status|	Mensagem|
|-----------------|-------|---------|
|Formato inválido	CEP sem 8 dígitos ou com caracteres inválidos|	422|	invalid zipcode|
|CEP não encontrado	CEP com formato correto, mas inexistente na base de dados|	404	|can not find zipcode|

## Fórmulas de conversão
Celsius para Fahrenheit: F = C × 1.8 + 32
Celsius para Kelvin: K = C + 273
Dicas de APIs externas

Você pode utilizar as seguintes APIs (ou similares) para obter os dados:

*Localização:* [ViaCEP](https://viacep.com.br/)

*Temperatura:* [WeatherAPI](https://www.weatherapi.com/)

# Requisitos de infraestrutura e deploy

- [ ] Docker: O projeto deve possuir um Dockerfile para containerização.
- [ ] Cloud Run: A aplicação deve ser implantada no Google Cloud Run (pode utilizar o free tier).
- [x] Testes: Devem ser implementados testes automatizados que comprovem o funcionamento das conversões e das requisições.

# Entregável

- [x] Código fonte: Link do repositório no GitHub.
- [ ] URL de acesso: O endereço ativo da aplicação no Google Cloud Run (deve constar no README).
- [x] Testes: O projeto deve conter testes automatizados.

# Regras de entrega

- [x] Repositório exclusivo: O repositório deve conter apenas o projeto em questão.
- [x] Branch principal: Todo o código deve estar na branch main.

# README: O arquivo deve conter:

A URL do sistema rodando no Cloud Run.
Instruções de como rodar os testes e a aplicação localmente via Docker.

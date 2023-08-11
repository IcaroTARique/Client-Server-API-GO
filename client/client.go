package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type USDBRL struct {
	Bid string `json:"bid"`
}

type Cotacao struct {
	CotacaoUsdBrl USDBRL `json:"USDBRL"`
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Limite de tempo para a requisição excedido! Tente novamente.")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte("Dolar: " + cotacao.CotacaoUsdBrl.Bid))
	if err != nil {
		panic(err)
	}

}

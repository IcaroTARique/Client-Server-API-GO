package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var debugFlag bool

type USDBRL struct {
	Bid string `json:"bid"`
}

type Cotacao struct {
	CotacaoUsdBrl USDBRL `json:"USDBRL"`
}

func main() {
	const create string = `
	CREATE TABLE IF NOT EXISTS cotacoes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	bid VARCHAR(10)
	);`

	var err error
	db, err = sql.Open("sqlite3", "cotacoes.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(create)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/cotacao", GetCotacao)
	http.ListenAndServe(":8080", nil)

}

func GetCotacao(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
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

	cotacao.DBPersist()
	//cotacao.SelectAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)

}

func (c *Cotacao) DBPersist() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := db.Prepare("insert into cotacoes (bid) VALUES (?)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.ExecContext(ctx, c.CotacaoUsdBrl.Bid)
	if err != nil {
		log.Println("Limite de tempo para a requisição excedido! Tente novamente.")
	}
}

func (c *Cotacao) SelectAll() {
	rows, err := db.Query("SELECT bid FROM cotacoes")
	if err != nil {
		panic(err)
	}

	var cotacoes []Cotacao
	for rows.Next() {
		var c Cotacao
		err = rows.Scan(&c.CotacaoUsdBrl.Bid)
		if err != nil {
			panic(err)
		}
		cotacoes = append(cotacoes, c)
	}

	fmt.Println(cotacoes)
}

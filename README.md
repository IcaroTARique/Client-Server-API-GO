# Client-Server-API-GO

## Client.go

A proposta é que o Client.go realize uma requisição HTTP no Server.go para obter a cotação do dólar. Essa funcionalidade
deve ser implementada utilizando CONTEXTOS. O tempo máximo de resposta aceitável do deve ser de 300 milissegundos. A cotação também deve ser 

- Consulta de dados ao Server.go
- Resposta em no max 300 ms
- Recebe do Server.go
- Em caso de timeout -> entrada no log
- Salvar a cotação num arquivo "cotacao.txt"


## Server.go

O Server.go deve consumir uma API através de uma requisição ao seguinte link https://economia.awesomeapi.com.br/json/last/USD-BRL. A URL disponibilizada
fornece, em formato JSON uma cotação USD - BRL. O Server.go deve consultar a API e obter a resposta em no máximo 200 milissegundos. Essa funcionalidade
deve ser implementada utilizando CONTEXTOS. A persistência em banco de dados tem o tempo máximo de 10ms.
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

- Consulta API com timeout de 200ms
- Persistência no banco com timeout de 10ms
- Em caso de timeout -> entrada no log

O valor em JSON contém a seguinte estrutura:

```json
{
    "USDBRL": {
        "code": "USD",
        "codein": "BRL",
        "name": "Dólar Americano/Real Brasileiro",
        "high": "4.9186",
        "low": "4.846",
        "varBid": "-0.0429",
        "pctChange": "-0.87",
        "bid": "4.8738",
        "ask": "4.8743",
        "timestamp": "1691182792",
        "create_date": "2023-08-04 17:59:52"
        }
}
```
O valor a ser armazenado deve ser o bid.

recebe instruções para armazenar no banco de dados, com 10 milissegundos para resposta, os registros.

- Receber a chamada do Client.go
- Consultar API e receber resposta em no max 200 ms
- Persistir a informação

## Banco de dados

Banco de dados SQLite e a tabela deve ter apenas um campo ()
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	time.Sleep(150 * time.Second)
	connectionStr := "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	connection, err := sql.Open("postgres", connectionStr)
	if err != nil {
		// fmt.Errorf("Não foi possível se conectar no banco %w", err)
		panic(err)
	}

	rows, err := connection.Query("SELECT version();")
	if err != nil {
		// fmt.Errorf("Não foi possível executar sua query %w", err)
		panic(err)
	}

	for rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Println("Aqui está sua versão de banco de dados: \n", version)
	}

	fmt.Println("O código está em sleep")
	time.Sleep(120 * time.Second)

	rows.Close()

	connection.Close()
}

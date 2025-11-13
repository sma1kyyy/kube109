package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var gwAppName, gwBindPort, gwDbURL string

func init() {
	var ok bool
	if gwAppName, ok = os.LookupEnv("GATEWAY_APP_NAME"); !ok {
		gwAppName = "gateway"
	}

	if gwBindPort, ok = os.LookupEnv("GATEWAY_BIND_PORT"); !ok {
		gwBindPort = "8080"
	}

	if gwDbURL, ok = os.LookupEnv("DATABASE_URL"); !ok {
		gwDbURL = "postgres://user:pass@localhost:5432/demodb"
	}
}

func main() {
	fmt.Println("GATEWAY_APP_NAME:", gwAppName)
	fmt.Println("GATEWAY_BIND_PORT:", gwBindPort)
	fmt.Println("DATABASE_URL:", gwDbURL)

	// Подключение к БД
	pgCtx, pgCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pgCancel()

	pgConn, err := pgx.Connect(pgCtx, gwDbURL)
	if err != nil {
		fmt.Println("database connect error:", err)
	} else {
		defer pgConn.Close(pgCtx)
		fmt.Println("database connected successfully")
	}

	http.HandleFunc("/", rootHandler)

	fmt.Printf("\nserver started on :%s\n", gwBindPort)

	if err := http.ListenAndServe(":"+gwBindPort, nil); err != nil {
		log.Fatalln(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from GoLang server"))
}
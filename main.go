package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hebelsan/go-web-shell/handlers"
	"github.com/hebelsan/go-web-shell/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.Handle("/cmd",
		middleware.Log(
			middleware.TokenCheck(
				http.HandlerFunc(handlers.CommandHandler))))

	port := flag.String("port", "5555", "Port of server")
	token := flag.String("token", "12345", "Secret required to run /cmd")
	flag.Parse()

	err := os.Setenv("SECRET_TOKEN", *token)
	if err != nil {
		panic("unable to set env SECRET_TOKEN")
	}

	log.Println(fmt.Sprintf("Server is starting on port %s", *port))
	log.Fatal(http.ListenAndServe(":"+*port, mux))
}

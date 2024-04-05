package web

import (
	"air-e-notification-adviser/config"
	"fmt"
	"log"
	"net/http"
)

func StartServer(cfg *config.EnvVars) {
	http.HandleFunc("/", HelloHandler)

	log.Println("Listening on port", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

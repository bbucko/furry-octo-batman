package main

import (
	"os"
	"net/http"
	"log"
	"fmt"
)

func getEnv(key string, defVal string) (env string) {
	env = os.Getenv(key)

	if env == "" {
		env = defVal
	}
	return
}


func main() {
	host := getEnv("HOST", "")
	port := getEnv("PORT", "8080")
	wsPort := getEnv("WSPORT", "8080")
//	mongoURL := getEnv("MONGOHQ_URL", "localhost")

	bind := fmt.Sprintf("%s:%s", host, port)
	log.Println("Starting server on", bind, "with websocket on port", wsPort)

	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

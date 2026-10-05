package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/betacalendars/betacalendars-calendar-boundary-lab/internal/httpapi"
)

var version = "dev"

func main() {
	httpapi.Version = version
	port := 8080
	if v := os.Getenv("PORT"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n < 65536 { port = n }
	}
	addr := ":" + strconv.Itoa(port)
	log.Printf("calendar-boundary-lab listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, httpapi.New()))
}

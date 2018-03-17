package main

import (
	"net/http"

	"github.com/colinx05/comp-availability/handlers"
)

// main configures the handlers and starts up the server
func main() {
	handlers.ConfigureHandlers()
	http.ListenAndServe(":8080", handlers.Router)
}

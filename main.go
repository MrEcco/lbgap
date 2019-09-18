package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Custom listen port
	listenPort := getEnv("CUSTOM_LISTEN_PORT", ":80")

	// Custom return body
	returnBody := getEnv("CUSTOM_RESPONCE", "OK")

	// Custom return code
	var returnCode int
	if r, err := strconv.Atoi(getEnv("CUSTOM_CODE", "200")); err != nil {
		log.Printf("Fallback to 200 return code. \"%s\" is not an integer", getEnv("CUSTOM_CODE", "200"))
		returnCode = 200
	} else {
		returnCode = r
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("from %s", r.RemoteAddr)
		w.WriteHeader(returnCode)
		fmt.Fprintf(w, returnBody)
	})

	if errCustomPort := http.ListenAndServe(listenPort, nil); errCustomPort != nil {
		// Bad port. Lets try to listen just ":80"
		log.Printf("Cannot listen \"%s\" port. Fallback to \":80\"", listenPort)
		if err := http.ListenAndServe(":80", nil); err != nil {
			log.Printf("Cannot listen fallback port: %s", err.Error())
		}
	}
}

// Get env with fallback
func getEnv(key string, fallback string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		env = fallback
	}
	return env
}

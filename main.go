package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /greet", func(res http.ResponseWriter, req *http.Request) {
		response := map[string]interface{}{
			"status":  http.StatusOK,
			"message": "hello world danny",
		}

		data, err := json.Marshal(&response)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("error parsing json"))
		}

		res.Write(data)
	})

	// upgrade program
	server := &http.Server{
		Addr:         ":8002",
		Handler:      router,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
	}

	log.Printf("Started server at port %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server failed to start %#v", err)
	}
}

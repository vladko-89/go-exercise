package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := rand.Intn(6)
		w.Write([]byte(strconv.Itoa(res)))
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running on port 8081")
	server.ListenAndServe()
}

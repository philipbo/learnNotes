package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request my handler start", time.Now())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 1; i < 11; i++ {
		go work(ctx, i)
	}
	time.Sleep(3 * time.Second)
	log.Println("request my handler end", time.Now())

	w.Write([]byte("hello world!\n"))
}

func main() {
	http.HandleFunc("/cancel", myHandler)
	srv := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("start server http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}

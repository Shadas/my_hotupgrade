package main

import (
	"log"
	"my_hotupgrade/grace"
	"net/http"
	"os"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(10 * time.Second))
	w.Write([]byte("WORLD!"))
	w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
	w.Write([]byte("\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)

	err := grace.ListenAndServe("localhost:8080", mux)
	// err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")
	os.Exit(0)
}

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request user agent: %s", r.Header.Get("User-Agent"))
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	})
	if err := http.ListenAndServe(":8899", nil); err != nil {
		log.Fatal(err)
	}
}

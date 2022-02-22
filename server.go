package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	_ "net/http/pprof"

	"github.com/lucas-clemente/quic-go/http3"
)

type binds []string

func Serv() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "6121"
	}

	bs := binds{"0.0.0.0:" + port}

	var wg sync.WaitGroup

	wg.Add(len(bs))

	for _, b := range bs {
		bCap := b
		go func() {
		  var err error

		  log.Printf("Listening on %s", bCap)

			http3.ListenAndServe(bCap, "./certs/cert.pem", "./certs/priv.pem", handler())

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Done serving on", bCap)
			wg.Done()
		}()
	}

	wg.Wait()
}

func handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%#v\n", r)
		fmt.Fprintf(w, "Hello to cryptograph API")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%#v\n", r)
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/kmd", handleCertificat)

	return mux
}

func handleCertificat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		
	}
}
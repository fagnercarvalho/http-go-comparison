package main

import (
	"fmt"
	"net/http"
	"path"
	"sync"

	"github.com/lucas-clemente/quic-go/http3"
)

var port = "6124"

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cors(w)
		w.Write([]byte("Hello from QUIC!"))
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		certFile, keyFile := getCerts()
		err := http3.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), certFile, keyFile, handler)
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Printf("Started quic-go server on port %s! \n", port)
	wg.Wait()
}

func getCerts() (string, string) {
	return path.Join("/go/src/app", "cert.pem"), path.Join("/go/src/app", "key.pem")
}

func cors(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
}

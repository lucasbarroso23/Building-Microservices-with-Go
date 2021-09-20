package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Ooops"))
			// all of this can be replaced by
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}

		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello again %s", d)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}

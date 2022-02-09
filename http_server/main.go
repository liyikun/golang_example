package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleHello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	select {
	case <-time.After(10 * time.Second):
		name := req.URL.Query().Get("name")
		str := fmt.Sprintln("hello world", name)
		w.Write([]byte(str))
		w.Header().Add("Content-Type", "text/plain")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println("server: ", err)
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
		}
	}
}

func main() {

	http.HandleFunc("/", handleHello)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

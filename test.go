package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	    // http://localhost:8080/
	    	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	    		fmt.Fprint(writer, "Hello heroku")
	    	})
	    	if err := http.ListenAndServe(":8080",mux); err != nil {
	    		log.Fatalf("big problem %v ", err)
			}

}

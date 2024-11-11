package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Good morning with everyone, my name is Edgar!")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}
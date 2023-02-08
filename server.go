package main

import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    http.HandleFunc("/add/3/4", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "7")
    })


    fmt.Printf("Starting server at port 4321\n")
    if err := http.ListenAndServe(":4321", nil); err != nil {
        log.Fatal(err)
    }
}

package main

import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    fmt.Printf("Starting server at port 4321\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
}
}

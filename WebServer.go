package main
import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Friends!")
    })
    log.Println("Server listening on port 8080....")
    http.ListenAndServe(":8080", nil)
}
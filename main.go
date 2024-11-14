package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    client, err := ConnectDB()
    if err != nil {
        log.Fatalf("Could not connect to MongoDB: %v", err)
    }
    defer client.Disconnect(nil)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        fmt.Fprintln(w, "Exam Application Project")
    })

    http.HandleFunc("POST /person/create", CreatePerson)
    http.HandleFunc("GET /person/getone/", GetPerson)
    http.HandleFunc("PUT /person/update/", UpdatePerson)
    http.HandleFunc("DELETE /person/delete/", DeletePerson)
    http.HandleFunc("GET /person/getage/", GetAge)

    port := os.Getenv("PORT")
    fmt.Printf("Server started on port %s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

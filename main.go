package main

import (
    "fmt"
    "log"
    "net/http"
    "examapp/config"
    "examapp/controllers"
)

func main() {
    client, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("Could not connect to MongoDB: %v", err)
    }
    defer client.Disconnect(nil)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        fmt.Fprintln(w, "Exam Application Project")
    })

    http.HandleFunc("/person/create", controllers.CreatePerson)
    http.HandleFunc("/person/getone", controllers.GetPerson)
    http.HandleFunc("/person/update", controllers.UpdatePerson)
    http.HandleFunc("/person/delete", controllers.DeletePerson)
    http.HandleFunc("/person/getage", controllers.GetAge)

    port := "8000"
    fmt.Printf("Server started on port %s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

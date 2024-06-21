package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Manipulador para a rota "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html") // Serve o arquivo HTML
    })

    fmt.Println("Servidor rodando em http://localhost:8080/")
    http.ListenAndServe(":8080", nil) 
}

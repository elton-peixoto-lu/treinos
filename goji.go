package main

import (
    "fmt"
    "net/http"

    "goji.io"
    "goji.io/pat"
)

func main() {
    mux := goji.NewMux()

    // Rota principal
    mux.HandleFunc(pat.Get("/"), func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Bem-vindo ao meu site com Goji!")
    })

    // Rota para saudação personalizada
    mux.HandleFunc(pat.Get("/ola/:nome"), func(w http.ResponseWriter, r *http.Request) {
        nome := pat.Param(r, "nome")
        fmt.Fprintf(w, "Olá, %s!", nome)
    })

    // Inicia o servidor
    fmt.Println("Servidor rodando na porta 8000...")
    http.ListenAndServe(":8000", mux)
}

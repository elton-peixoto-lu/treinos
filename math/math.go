package main

import (
    "fmt"
    "math/big"
    "net/http"
    "strconv"
)

func fatorial(n int64) *big.Int {
    result := big.NewInt(1)
    for i := int64(2); i <= n; i++ {
        result.Mul(result, big.NewInt(i))
    }
    return result
}

func fatorialHandler(w http.ResponseWriter, r *http.Request) {
    numeroStr := r.URL.Query().Get("numero")

    numero, err := strconv.Atoi(numeroStr)
    if err != nil || numero < 0 {
        http.Error(w, "Número inválido", http.StatusBadRequest)
        return
    }

    resultado := fatorial(int64(numero))
    fmt.Fprintf(w, "O fatorial de %d é %s", numero, resultado.String())
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        html := `
<!DOCTYPE html>
<html>
<head>
    <title>Calculadora de Fatorial</title>
</head>
<body>
    <h1>Calculadora de Fatorial</h1>
    <form method="GET" action="/fatorial">
        <input type="number" name="numero" min="0" required>
        <button type="submit">Calcular</button>
    </form>
</body>
</html>
`
        fmt.Fprint(w, html)
    })

    http.HandleFunc("/fatorial", fatorialHandler)

    fmt.Println("Servidor iniciado na porta 8080")
    http.ListenAndServe(":8080", nil)
}

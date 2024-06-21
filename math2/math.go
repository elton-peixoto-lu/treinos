package main

import (
    "fmt"
    "math/big"
    "net/http"
    "strconv"
)

func fatorial(n int64) *big.Int {
    if n < 0 {
        return big.NewInt(0)
    } else if n == 0 {
        return big.NewInt(1)
    } else {
        result := big.NewInt(1)
        for i := int64(2); i <= n; i++ {
            result.Mul(result, big.NewInt(i))
        }
        return result
    }
}

func fatorialHandler(w http.ResponseWriter, r *http.Request) {
    numeroStr := r.URL.Query().Get("numero")
    numero, err := strconv.Atoi(numeroStr)
    if err != nil || numero < 0 {
        http.Error(w, "Número inválido", http.StatusBadRequest)
        return
    }
    resultado := fatorial(int64(numero))
    fmt.Fprintf(w, resultado.String())
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        html := `
<!DOCTYPE html>
<html>
<head>
    <title>Calculadora de Fatorial</title>
    <script>
        function calcularFatorial() {
            var numero = document.getElementById("numero").value;
            var xhr = new XMLHttpRequest();
            xhr.open("GET", "/fatorial?numero=" + numero, true);
            xhr.onload = function() {
                if (xhr.status == 200) {
                    document.getElementById("resultado").textContent = xhr.responseText;
                } else {
                    document.getElementById("resultado").textContent = "Erro ao calcular.";
                }
            };
            xhr.send();
        }
    </script>
</head>
<body>
    <h1>Calculadora de Fatorial</h1>
    <form>
        <input type="number" id="numero" min="0" required>
        <button type="button" onclick="calcularFatorial()">Calcular</button>
    </form>
    <p>Resultado: <span id="resultado"></span></p>
</body>
</html>
`
        fmt.Fprint(w, html)
    })

    http.HandleFunc("/fatorial", fatorialHandler)

    fmt.Println("Servidor iniciado na porta http:localhost:9090")
    http.ListenAndServe(":9090", nil)
}


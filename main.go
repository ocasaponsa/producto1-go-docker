package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Soy alumno de la UOC")
}

func main() {
    
  // Servir archivos estáticos: /static/...
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // Página principal: devuelve el HTML
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
  })

  http.ListenAndServe(":8080", nil)
}

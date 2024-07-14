// cmd/main.go

package main

import (
	"log"
	"net/http"

	"github.com/adriannylelis/dance-moves-api/internal/database"
	"github.com/adriannylelis/dance-moves-api/internal/movement"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar arquivo .env: %v", err)
    }

    // Inicializa o banco de dados
    database.InitDB()

    // Inicializa o roteador mux
    r := mux.NewRouter()

    // Inicializa o handler dos movimentos de dança
    movementHandler := movement.NewMovementHandler()

    // Define as rotas
    r.HandleFunc("/api/movements", movementHandler.GetMovements).Methods("GET")
    r.HandleFunc("/api/movements/{id}", movementHandler.GetMovementByID).Methods("GET")
    r.HandleFunc("/api/movements", movementHandler.CreateMovement).Methods("POST")
    r.HandleFunc("/api/movements/{id}", movementHandler.UpdateMovement).Methods("PUT")
    r.HandleFunc("/api/movements/{id}", movementHandler.DeleteMovement).Methods("DELETE")

    // Inicia o servidor HTTP
    log.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

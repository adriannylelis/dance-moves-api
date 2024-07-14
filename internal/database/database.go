// internal/database/database.go

package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar arquivo .env: %v", err)
    }

    // Obtém as variáveis de ambiente
    dataSourceName := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

    // Inicializa a conexão com o banco de dados MySQL
    var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

    // Verifica se a conexão com o banco de dados é bem-sucedida
    err = db.Ping()
    if err != nil {
        log.Fatalf("Erro ao pingar o banco de dados: %v", err)
    }

    log.Println("Conexão com o banco de dados estabelecida")
}

func GetDB() *sql.DB {
    return db
}

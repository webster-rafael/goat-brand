package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"goatbrand-backend/internal/repository"
	transporthttp "goatbrand-backend/internal/transport/http"
	"goatbrand-backend/internal/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Falha ao cerregar o arquivo .env, o sistema procurará variaveis nativas de ambiente")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL não configurada no env")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Erro abrindo conexão com banco de dados:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Erro contantando banco de dados:", err)
	}

	runMigrations(db)

	userRepo := repository.NewUserPostgres(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := transporthttp.NewUserHandler(userUsecase)
	router := transporthttp.NewRouter(userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func runMigrations(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		phone VARCHAR(50) NOT NULL,
		code VARCHAR(10) NOT NULL,
		payment_status VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Falha crítica ao executar migrations de database: %v", err)
	}
	log.Println("Migrations da tabela 'users' rodadas com sucesso!")
}

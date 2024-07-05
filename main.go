package main

import (
	"log"
	"net/http"
	"os"
	"sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/Mafaz03/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()


	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in env")
	}

	DB_url := os.Getenv("DB_URL")
	if DB_url == "" {
		log.Fatal("DB URL not found in env")
	}

	conn, err := sql.Open(("postgres", DB_url)
	if err != nil {
		log.Fatal("Connection to DB Failed")
	}

	type apiConfig struct {
		DB *database.Queries
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r1router := chi.NewRouter()
	r1router.Get("/healthz", handler_readiness)
	r1router.Get("/err", handler_err)

	router.Mount("/v1", r1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("listening on Port number: %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	println("Port: ", portString)
}

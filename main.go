package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/Mafaz03/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	feed, _ := urlToFeed("https://wagslane.dev/index.xml")
	fmt.Println(feed)

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in env")
	}

	DB_url := os.Getenv("DB_URL")
	if DB_url == "" {
		log.Fatal("DB URL not found in env")
	}

	conn, err := sql.Open("postgres", DB_url)
	if err != nil {
		log.Fatal("Connection to DB Failed")
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
	r1router.Post("/users", apiCfg.handler_CreateUsers)
	r1router.Get("/getusers", apiCfg.GetUsersByKey)
	r1router.Get("/users", apiCfg.middlewareAuth(apiCfg.GetUsersByAuth))
	r1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handler_CreateFeed))
	r1router.Get("/feeds", apiCfg.handler_GetFeed)
	r1router.Post("/feeds_follow", apiCfg.middlewareAuth(apiCfg.handler_CreateFeedFollow))
	r1router.Get("/feeds_follow", apiCfg.handler_GetFeedFollow)
	r1router.Delete("/feeds_follow/{feedfollowid}", apiCfg.middlewareAuth(apiCfg.DeleteFeedsFollow))

	router.Mount("/v1", r1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("listening on Port number: %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	println("Port: ", portString)
}

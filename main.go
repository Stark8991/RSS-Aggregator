package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Stark8991/RSSAgg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	dbString := os.Getenv("DB_URL")

	if dbString == "" {
		log.Fatal("Database connection not found")
	}

	if portString == "" {
		log.Fatal("Port not found")
	}

	conn, err := sql.Open("pgx", dbString)

	if err != nil {

		log.Fatal("Failed to open a db connection")

	}

	queries := database.New(conn)

	apiConfig := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/users", apiConfig.middlewareAuth(apiConfig.handlUsersGet))
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiConfig.handlerCreateUser)
	v1Router.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlerFeedCreate))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port: %s", portString)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}

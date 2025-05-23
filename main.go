package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/HtetAungKhant23/blogAgg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// encode(sha256(random()::text::bytea), 'hex') -> for postgres
// uuid_generate_v4()::string -> for cockroach

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Can't connect to database: %v", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	// go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Get("/user/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPost))
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerFeedFollow))
	v1Router.Get("/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed-follow/{feedFollowId}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	v2Router := chi.NewRouter()
	v2Router.Get("/car-list", apiCfg.handlerCarList)

	router.Mount("/v2", v2Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting at port %v", portString)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal("Server starting error ", err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/frangdelsolar/todo_cli/pkg/config"
	"github.com/frangdelsolar/todo_cli/pkg/data"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/gorilla/mux"
)

const (
	PKG_NAME    = "Server PKG"
	PKG_VERSION = "1.0.1"
	PUBLIC_DIR  = "./server/public"
)

var (
	cfg *config.Config
	log *logger.Logger
	db  *data.Database
)

func main() {

	// Setup Global Configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("Failed to load config: %v", err)
	}

	// Define Logger
	log = logger.NewLogger(logger.LoggerConfig{
		PackageName:    PKG_NAME,
		PackageVersion: PKG_VERSION,
	})

	log.Info().Msgf("Running %s v%s", PKG_NAME, PKG_VERSION)
	log.Debug().Interface("Config", cfg).Msg("Loaded Config")

	// Connect to Database
	db, err = data.LoadDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}
	log.Debug().Msgf("Loaded Database: %s", db.Name())

	// CSRF

	// Define Router
	r := mux.NewRouter()

	// Middlewares
	r.Use(mux.CORSMethodMiddleware(r))
	// r.Use(csrf.Middleware(csrfConfig))

	// Public Routes
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(PUBLIC_DIR))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Auth Middleware

	// Protected Routes

	// Start Server
	http.Handle("/", r)

	// Start Server
	err = http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	log.Info().Msgf("Starting Server on Port: %s", cfg.ServerPort)

}

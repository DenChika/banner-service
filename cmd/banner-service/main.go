package main

import (
	"banner_service/internal/config"
	"banner_service/internal/repository"
	"banner_service/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// load envs
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// read configs
	cfg := config.MustLoad()

	// init logger
	log := initLogger()
	log.Info("Starting banner-service")

	// init database
	db := initDatabase(log, cfg.Db)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Error("Error closing db connection", log.With("err", err.Error()))
		}
	}(db)
	log.Info("Database connection established")

	// init Routes
	initRoutes(log)

	// start server
	srv := &http.Server{
		Addr:         ":" + cfg.HttpServer.Port,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("Failed starting banner-service")
	}
	log.Info("Shutting down banner-service")
}

func initLogger() *slog.Logger {
	var log *slog.Logger

	slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}

func initDatabase(log *slog.Logger, cfg config.Db) *sqlx.DB {
	dbConfig := &repository.DbConfig{
		User:     cfg.User,
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     cfg.Port,
		Name:     cfg.Name,
		Ssl:      cfg.Ssl,
		Driver:   cfg.Driver,
	}

	db, err := repository.ConnectToDb(dbConfig, postgres.NewDriver())
	if err != nil {
		log.Error("Failed connecting to database", log.With("error", err))
	}

	return db
}

func initRoutes(log *slog.Logger) {
	// TODO: uncomment when BannerGetter will be implemented
	//router := chi.NewRouter()
	//
	//router.Route("/banner", func(r chi.Router) {
	//	r.Get("/", banner.New(log))
	//})
}

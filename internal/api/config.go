package api

import (
	"log"
	"net/http"
	"os"

	Logger "github.com/binsabit/url_shortener/pkg/log"
	"github.com/joho/godotenv"
)

type config struct {
	port string
}

type application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	config      config
}

func ConfigureAndStart() {
	infoLogger, err := Logger.RegisterLogger("INFO", "logs")
	if err != nil {
		log.Fatal("connot configure logger", err)
	}
	errorLogger, err := Logger.RegisterLogger("ERROR", "logs")
	if err != nil {
		log.Fatal("connot configure logger", err)
	}
	var cfg config

	err = godotenv.Load(".env")

	if err != nil {
		errorLogger.Fatal("cannot configure server port ", err)
	}
	cfg.port = os.Getenv("SERVER_PORT")

	app := application{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
		config:      cfg,
	}

	srv := &http.Server{
		Addr:    ":" + app.config.port,
		Handler: app.router(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		app.ErrorLogger.Fatal("server cannot be startted")
	}

}

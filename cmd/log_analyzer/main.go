package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/KazikovAP/log_analyzer/config"
	"github.com/KazikovAP/log_analyzer/internal/application"
	"github.com/KazikovAP/log_analyzer/internal/infrastructure"
)

// go run cmd/log_analyzer/main.go
// --path https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs
// --from 2015-05-18T00:00:00Z --to 2015-05-30T00:00:00Z --format markdown

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Ошибка инициализации файла конфигурации: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ioAdapter := infrastructure.NewIOAdapter(os.Stdin, os.Stdout, logger)

	app := application.NewApp(cfg, ioAdapter)
	if err := app.Start(); err != nil {
		logger.Error("Application failed to start", "error", err)
	}
}

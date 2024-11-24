package config

import (
	"errors"
	"flag"
	"time"
)

type Config struct {
	LogURL       string
	From         time.Time
	To           time.Time
	ReportFormat string
}

func Init() (*Config, error) {
	pathFlag := flag.String("path", "", "Путь к лог-файлам (локальный путь или URL)")
	fromFlag := flag.String("from", "", "Начало временного диапазона в формате ISO8601 (например, 2024-11-01T00:00:00Z)")
	toFlag := flag.String("to", "", "Конец временного диапазона в формате ISO8601 (например, 2024-11-10T23:50:00Z)")
	formatFlag := flag.String("format", "markdown", "Формат отчета (markdown/adoc)")

	flag.Parse()

	if *pathFlag == "" {
		return nil, errors.New("обязательный флаг -path не указан")
	}

	cfg := &Config{
		LogURL:       *pathFlag,
		ReportFormat: *formatFlag,
	}

	if *fromFlag != "" {
		fromTime, err := time.Parse(time.RFC3339, *fromFlag)
		if err == nil {
			cfg.From = fromTime
		} else {
			return nil, errors.New("неверный формат времени 'from': " + err.Error())
		}
	}

	if *toFlag != "" {
		toTime, err := time.Parse(time.RFC3339, *toFlag)
		if err == nil {
			cfg.To = toTime
		} else {
			return nil, errors.New("неверный формат времени 'to': " + err.Error())
		}
	}

	return cfg, nil
}

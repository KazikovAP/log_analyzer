package application

import (
	"fmt"
	"net/url"
	"time"

	"github.com/KazikovAP/log_analyzer/config"
	"github.com/KazikovAP/log_analyzer/internal/domain"
)

type ioAdapter interface {
	Output(content string)
}

type App struct {
	io  ioAdapter
	cfg *config.Config
}

func NewApp(cfg *config.Config, io ioAdapter) *App {
	return &App{cfg: cfg, io: io}
}

func (a *App) Start() error {
	path := a.cfg.LogURL
	from := a.cfg.From
	to := a.cfg.To
	format := a.cfg.ReportFormat

	records, err := a.loadLogs(path)
	if err != nil {
		return err
	}

	filteredRecords := records.Filter(
		FilterByTimeRange(from, to),
	)

	a.io.Output(fmt.Sprintf("\nЗагружено логов после фильтрации: %d\n", len(filteredRecords)))

	report := a.generateReport(filteredRecords, format, path, from, to)
	a.io.Output(report)

	return nil
}

func (a *App) loadLogs(path string) (domain.LogRecords, error) {
	var records []domain.LogRecord

	var err error

	var logReader domain.LogReader

	parser := domain.NewLogParser()

	if isURL(path) {
		logReader = domain.NewURLLogReader(path)
	} else {
		logReader = domain.NewFileLogReader(path)
	}

	records, err = logReader.ReadLogs(parser)

	if err != nil {
		return nil, err
	}

	a.io.Output(fmt.Sprintf("\nЗагружено логов: %d", len(records)))

	return records, nil
}

func (a *App) generateReport(records []domain.LogRecord, format, path string, from, to time.Time) string {
	analyzer := domain.NewLogAnalyzer(records)

	analysisResult := analyzer.Analyze()

	filename := "url"
	if !isURL(path) {
		filename = path
	}

	startDate, endDate := "-", "-"
	if !from.IsZero() {
		startDate = from.Format("02.01.2006")
	}

	if !to.IsZero() {
		endDate = to.Format("02.01.2006")
	}

	reportData := domain.ReportData{
		Filename:       filename,
		StartDate:      startDate,
		EndDate:        endDate,
		TotalRequests:  analysisResult.TotalRequests,
		ResourceCount:  analysisResult.ResourceCount,
		StatusCount:    analysisResult.StatusCount,
		AvgSize:        analysisResult.AvgSize,
		PercentileSize: analysisResult.PercentileSize,
	}

	var formatter domain.Formatter

	if format == "adoc" {
		formatter = &domain.Asciidoc{}
	} else {
		formatter = &domain.Markdown{}
	}

	return formatter.Render(&reportData)
}

func FilterByTimeRange(from, to time.Time) domain.LogFilter {
	return func(record domain.LogRecord) bool {
		return (from.IsZero() || record.TimeLocal.After(from)) &&
			(to.IsZero() || record.TimeLocal.Before(to))
	}
}

func isURL(path string) bool {
	parsedURL, err := url.Parse(path)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

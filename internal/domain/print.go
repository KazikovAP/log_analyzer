package domain

import (
	"fmt"
	"net/http"
	"strings"
)

type ReportData struct {
	Filename       string
	StartDate      string
	EndDate        string
	TotalRequests  int
	ResourceCount  map[string]int
	StatusCount    map[int]int
	AvgSize        float64
	PercentileSize int
}

type Formatter interface {
	Render(data *ReportData) string
}

type Markdown struct{}

func (m *Markdown) Render(data *ReportData) string {
	var report strings.Builder

	report.WriteString("\n#### Общая информация\n\n")
	report.WriteString("|        Метрика        |       Значение     |\n|:---------------------:|-------------------:|\n")
	report.WriteString(fmt.Sprintf("|       Файл(-ы)        | %18s |\n", data.Filename))
	report.WriteString(fmt.Sprintf("|    Начальная дата     | %18s |\n", data.StartDate))
	report.WriteString(fmt.Sprintf("|     Конечная дата     | %18s |\n", data.EndDate))
	report.WriteString(fmt.Sprintf("|  Количество запросов  | %18d |\n", data.TotalRequests))
	report.WriteString(fmt.Sprintf("| Средний размер ответа | %17.2fb |\n", data.AvgSize))
	report.WriteString(fmt.Sprintf("|  95p размера ответа   | %17db |\n\n", data.PercentileSize))

	report.WriteString("#### Запрашиваемые ресурсы\n\n| Ресурс                              | Количество |\n")
	report.WriteString("|:-----------------------------------:|-----------:|\n")

	for resource, count := range data.ResourceCount {
		report.WriteString(fmt.Sprintf("| %-35s | %10d |\n", resource, count))
	}

	report.WriteString("\n#### Коды ответа\n\n| Код |                 Имя                 | Количество |\n")
	report.WriteString("|:---:|:-----------------------------------:|-----------:|\n")

	for status, count := range data.StatusCount {
		report.WriteString(fmt.Sprintf("| %3d | %-35s | %10d |\n", status, http.StatusText(status), count))
	}

	return report.String()
}

type Asciidoc struct{}

func (a *Asciidoc) Render(data *ReportData) string {
	var report strings.Builder

	report.WriteString("\n=== Общая информация\n\n")
	report.WriteString("| Метрика                  | Значение     |\n")
	report.WriteString("|--------------------------|--------------|\n")
	report.WriteString(fmt.Sprintf("| Файл(-ы)                 | %12s |\n", data.Filename))
	report.WriteString(fmt.Sprintf("| Начальная дата           | %12s |\n", data.StartDate))
	report.WriteString(fmt.Sprintf("| Конечная дата            | %12s |\n", data.EndDate))
	report.WriteString(fmt.Sprintf("| Количество запросов      | %12d |\n", data.TotalRequests))
	report.WriteString(fmt.Sprintf("| Средний размер ответа    | %11.2fb |\n", data.AvgSize))
	report.WriteString(fmt.Sprintf("| 95-й перцентиль размера  | %11db |\n\n", data.PercentileSize))

	report.WriteString("=== Запрашиваемые ресурсы\n\n| Ресурс                              | Количество |\n")
	report.WriteString("|-------------------------------------|------------|\n")

	for resource, count := range data.ResourceCount {
		report.WriteString(fmt.Sprintf("| %-35s | %10d |\n", resource, count))
	}

	report.WriteString("\n=== Коды ответа\n\n| Код |                 Имя                 | Количество |\n")
	report.WriteString("|-----|-------------------------------------|------------|\n")

	for status, count := range data.StatusCount {
		report.WriteString(fmt.Sprintf("| %3d | %-35s | %10d |\n", status, http.StatusText(status), count))
	}

	return report.String()
}

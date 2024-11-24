package domain_test

import (
	"testing"

	"github.com/KazikovAP/log_analyzer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestAnalyze_EmptyLogs(t *testing.T) {
	analyzer := domain.NewLogAnalyzer([]domain.LogRecord{})
	result := analyzer.Analyze()

	assert.Equal(t, 0, result.TotalRequests)
	assert.Empty(t, result.ResourceCount)
	assert.Empty(t, result.StatusCount)
	assert.Equal(t, 0.0, result.AvgSize)
	assert.Equal(t, 0, result.PercentileSize)
}

func TestAnalyze_SingleRecord(t *testing.T) {
	records := []domain.LogRecord{
		{Request: "/index", Status: 200, BodyBytesSent: 100},
	}
	analyzer := domain.NewLogAnalyzer(records)
	result := analyzer.Analyze()

	assert.Equal(t, 1, result.TotalRequests)
	assert.Equal(t, map[string]int{"/index": 1}, result.ResourceCount)
	assert.Equal(t, map[int]int{200: 1}, result.StatusCount)
	assert.Equal(t, 100.0, result.AvgSize)
	assert.Equal(t, 100, result.PercentileSize)
}

func TestAnalyze_MultipleRecords(t *testing.T) {
	records := []domain.LogRecord{
		{Request: "/index", Status: 200, BodyBytesSent: 100},
		{Request: "/index", Status: 404, BodyBytesSent: 200},
		{Request: "/about", Status: 200, BodyBytesSent: 300},
	}
	analyzer := domain.NewLogAnalyzer(records)
	result := analyzer.Analyze()

	assert.Equal(t, 3, result.TotalRequests)
	assert.Equal(t, map[string]int{"/index": 2, "/about": 1}, result.ResourceCount)
	assert.Equal(t, map[int]int{200: 2, 404: 1}, result.StatusCount)
	assert.Equal(t, 200.0, result.AvgSize)
	assert.Equal(t, 300, result.PercentileSize)
}

func TestAnalyze_LargeBodyBytesSent(t *testing.T) {
	records := []domain.LogRecord{
		{Request: "/index", Status: 200, BodyBytesSent: 500},
		{Request: "/about", Status: 200, BodyBytesSent: 1000},
		{Request: "/contact", Status: 404, BodyBytesSent: 1500},
		{Request: "/index", Status: 200, BodyBytesSent: 200},
		{Request: "/about", Status: 500, BodyBytesSent: 700},
	}
	analyzer := domain.NewLogAnalyzer(records)
	result := analyzer.Analyze()

	assert.Equal(t, 5, result.TotalRequests)
	assert.Equal(t, map[string]int{"/index": 2, "/about": 2, "/contact": 1}, result.ResourceCount)
	assert.Equal(t, map[int]int{200: 3, 404: 1, 500: 1}, result.StatusCount)
	assert.Equal(t, 780.0, result.AvgSize)
	assert.Equal(t, 1500, result.PercentileSize)
}

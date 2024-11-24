package application_test

import (
	"testing"
	"time"

	"github.com/KazikovAP/log_analyzer/config"
	"github.com/KazikovAP/log_analyzer/internal/application"
	"github.com/KazikovAP/log_analyzer/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockIOAdapter struct {
	mock.Mock
}

func (m *MockIOAdapter) Output(content string) {
	m.Called(content)
}

type MockLogReader struct {
	mock.Mock
}

func (m *MockLogReader) ReadLogs(parser *domain.LogParser) ([]domain.LogRecord, error) {
	args := m.Called(parser)
	return args.Get(0).([]domain.LogRecord), args.Error(1)
}

func TestApp_Start(t *testing.T) {
	mockIO := new(MockIOAdapter)
	mockIO.On("Output", mock.Anything).Return()

	cfg := &config.Config{
		LogURL:       "http://example.com/logs",
		From:         time.Now().Add(-24 * time.Hour),
		To:           time.Now(),
		ReportFormat: "markdown",
	}

	app := application.NewApp(cfg, mockIO)

	err := app.Start()
	assert.NoError(t, err)

	mockIO.AssertExpectations(t)
}

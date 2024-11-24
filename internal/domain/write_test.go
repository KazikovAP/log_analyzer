package domain_test

import (
	"testing"

	"github.com/KazikovAP/log_analyzer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestLogParser_ParseLine_InvalidLogFormat(t *testing.T) {
	parser := domain.NewLogParser()
	line := `invalid log format`
	_, err := parser.ParseLine(line)
	assert.Error(t, err)
	assert.Equal(t, "лог не соответствует формату: invalid log format", err.Error())
}

func TestLogParser_ParseLine_InvalidTimeFormat(t *testing.T) {
	parser := domain.NewLogParser()
	line := `127.0.0.1 - - [invalid-time] "GET / HTTP/1.1" 200 512 "http://example.com" "Mozilla/5.0"`
	_, err := parser.ParseLine(line)
	assert.Error(t, err)
}

func TestLogParser_ParseLine_InvalidStatusCode(t *testing.T) {
	parser := domain.NewLogParser()
	line := `127.0.0.1 - - [17/Nov/2024:10:15:30 +0000] "GET / HTTP/1.1" invalid-status 512 "http://example.com" "Mozilla/5.0"`
	_, err := parser.ParseLine(line)
	assert.Error(t, err)
}

func TestLogParser_ParseLine_InvalidBodyBytesSent(t *testing.T) {
	parser := domain.NewLogParser()
	line := `127.0.0.1 - - [17/Nov/2024:10:15:30 +0000] "GET / HTTP/1.1" 200 invalid-bytes "http://example.com" "Mozilla/5.0"`
	_, err := parser.ParseLine(line)
	assert.Error(t, err)
}

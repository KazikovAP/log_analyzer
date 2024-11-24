package domain_test

import (
	"testing"

	"github.com/KazikovAP/log_analyzer/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestFileLogReader_ReadLogs_EmptyPath(t *testing.T) {
	parser := &domain.LogParser{}
	reader := &domain.FileLogReader{Path: ""}
	_, err := reader.ReadLogs(parser)
	assert.Error(t, err, "путь не может быть пустым")
}

func TestFileLogReader_ReadLogs_FileNotFound(t *testing.T) {
	parser := &domain.LogParser{}
	reader := &domain.FileLogReader{Path: "non_existent_file.log"}
	_, err := reader.ReadLogs(parser)
	assert.Error(t, err, "не удалось открыть файл")
}

func TestURLLogReader_ReadLogs_FailToFetch(t *testing.T) {
	parser := &domain.LogParser{}
	reader := &domain.URLLogReader{URL: "http://invalid-url"}
	_, err := reader.ReadLogs(parser)
	assert.Error(t, err, "ошибка при загрузке URL")
}

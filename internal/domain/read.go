package domain

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

type LogReader interface {
	ReadLogs(parser *LogParser) ([]LogRecord, error)
}

type FileLogReader struct {
	Path string
}

type URLLogReader struct {
	URL string
}

func NewFileLogReader(path string) *FileLogReader {
	return &FileLogReader{Path: path}
}

func NewURLLogReader(url string) *URLLogReader {
	return &URLLogReader{URL: url}
}

func (f *FileLogReader) ReadLogs(parser *LogParser) ([]LogRecord, error) {
	if f.Path == "" {
		return nil, fmt.Errorf("путь не может быть пустым")
	}

	file, err := os.Open(f.Path)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	var records []LogRecord

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record, err := parser.ParseLine(scanner.Text())
		if err == nil {
			records = append(records, *record)
		}
	}

	return records, scanner.Err()
}

func (u *URLLogReader) ReadLogs(parser *LogParser) ([]LogRecord, error) {
	resp, err := http.Get(u.URL)
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке URL: %v", err)
	}
	defer resp.Body.Close()

	var records []LogRecord

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		record, err := parser.ParseLine(scanner.Text())
		if err == nil {
			records = append(records, *record)
		}
	}

	return records, scanner.Err()
}

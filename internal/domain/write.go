package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type LogRecord struct {
	RemoteAddr    string
	TimeLocal     time.Time
	Request       string
	Status        int
	BodyBytesSent int
	HTTPReferer   string
	HTTPUserAgent string
}

type LogParser struct{}

func NewLogParser() *LogParser {
	return &LogParser{}
}

func (p *LogParser) ParseLine(line string) (*LogRecord, error) {
	pattern := `^(\S+) - (\S+) \[(.*?)\] "(.*?)" (\d{3}) (\d+) "(.*?)" "(.*?)"$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("лог не соответствует формату: %s", line)
	}

	timeLayout := "02/Jan/2006:15:04:05 -0700"
	timeParsed, err := time.Parse(timeLayout, matches[3])

	if err != nil {
		return nil, err
	}

	status, _ := strconv.Atoi(matches[5])
	bodyBytesSent, _ := strconv.Atoi(matches[6])

	return &LogRecord{
		RemoteAddr:    matches[1],
		TimeLocal:     timeParsed,
		Request:       matches[4],
		Status:        status,
		BodyBytesSent: bodyBytesSent,
		HTTPReferer:   matches[7],
		HTTPUserAgent: matches[8],
	}, nil
}

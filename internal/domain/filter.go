package domain

type LogRecords []LogRecord

type LogFilter func(record LogRecord) bool

func (records LogRecords) Filter(filters ...LogFilter) []LogRecord {
	var filtered []LogRecord

outer:
	for _, record := range records {
		for _, filter := range filters {
			if !filter(record) {
				continue outer
			}
		}
		filtered = append(filtered, record)
	}

	return filtered
}

package domain

import "sort"

type LogAnalyzer struct {
	Records []LogRecord
}

type LogAnalysisResult struct {
	TotalRequests  int
	ResourceCount  map[string]int
	StatusCount    map[int]int
	AvgSize        float64
	PercentileSize int
}

func NewLogAnalyzer(records []LogRecord) *LogAnalyzer {
	return &LogAnalyzer{Records: records}
}

func (a *LogAnalyzer) Analyze() LogAnalysisResult {
	if len(a.Records) == 0 {
		return LogAnalysisResult{}
	}

	var totalSize int

	resourceCount := make(map[string]int)
	statusCount := make(map[int]int)
	totalRequests := 0

	for _, record := range a.Records {
		totalRequests++
		totalSize += record.BodyBytesSent

		resourceCount[record.Request]++
		statusCount[record.Status]++
	}

	avgSize := float64(totalSize) / float64(totalRequests)

	sortedSizes := make([]int, len(a.Records))
	for i, record := range a.Records {
		sortedSizes[i] = record.BodyBytesSent
	}

	sort.Ints(sortedSizes)

	percentileIndex := int(0.95 * float64(len(sortedSizes)))
	percentileSize := sortedSizes[percentileIndex]

	return LogAnalysisResult{
		TotalRequests:  totalRequests,
		ResourceCount:  resourceCount,
		StatusCount:    statusCount,
		AvgSize:        avgSize,
		PercentileSize: percentileSize,
	}
}

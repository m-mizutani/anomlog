package anomlog

import (
	"time"
)

// Model is current result of format analysis
type Model interface {
	read(log *Log) *Format
	count() int
	formats() []*Format
}

// Format describes fixed literal and variable of log message
type Format struct {
	Count     int       `json:"count"`
	Chunks    []Chunk   `json:"chunks"`
	Timestamp time.Time `json:"timestamp"`
}

func newFormat(log *Log) *Format {
	fmt := &Format{}
	fmt.Timestamp = time.Now()
	fmt.Count = 1
	fmt.Chunks = make([]Chunk, len(log.chunks))
	for idx, c := range log.chunks {
		fmt.Chunks[idx].Data = c.Data
	}
	return fmt
}

func (x *Format) matchRatio(log *Log) float64 {
	if len(log.chunks) == len(x.Chunks) {
		matched := 0
		for idx, chunk := range x.Chunks {
			if chunk.equals(log.chunks[idx]) {
				matched++
			}
		}
		return float64(matched) / float64(len(x.Chunks))
	}
	return 0
}

func (x *Format) merge(log *Log) {
	for idx, c := range log.chunks {
		x.Chunks[idx].merge(c)
	}
	x.Count++
}

func (x *Format) String() string {
	s := ""
	for _, c := range x.Chunks {
		s += c.String()
	}
	return s
}

// SimpleModel is most simple format analysis model
type SimpleModel struct {
	Formats []*Format
}

func findClosestFormat(formats []*Format, log *Log) (*Format, float64) {
	maxIdx := -1
	maxScore := 0.0

	for idx, fmt := range formats {
		score := fmt.matchRatio(log)
		if score > maxScore {
			// logger.Println(maxIdx, "->", idx)
			maxIdx = idx
			maxScore = score
		}
	}

	if maxIdx < 0 {
		return nil, 0
	}

	return formats[maxIdx], maxScore
}

func (x *SimpleModel) read(log *Log) *Format {
	if len(log.chunks) == 0 {
		return nil
	}

	fmt, maxScore := findClosestFormat(x.Formats, log)
	if len(x.Formats) == 0 || maxScore < 0.7 {
		fmt = newFormat(log)
		x.Formats = append(x.Formats, fmt)
	} else {
		fmt.merge(log)
	}

	return fmt
}

func (x *SimpleModel) count() int {
	sum := 0
	for _, fmt := range x.Formats {
		sum += fmt.Count
	}
	return sum
}

func (x *SimpleModel) formats() []*Format {
	return x.Formats
}

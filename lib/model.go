package anomlog

// Model is current result of format analysis
type Model interface {
	read(log *Log) *Format
	count() int
	dump() ([]byte, error)
}

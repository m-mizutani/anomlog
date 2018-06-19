package anomlog

// Chunk is data structure for a word
type Chunk struct {
	Data    string `json:"data"`
	IsParam bool   `json:"is_param"`
	freezed bool
}

// Log is a sequence of Chunk
type Log struct {
	text   string
	chunks []*Chunk
}

func newLog(line string, sp Splitter) *Log {
	log := Log{}
	log.text = line
	log.chunks = sp.Split(line)
	return &log
}

func (x *Log) String() string {
	// red := color.New(color.FgRed).SprintFunc()

	// s := fmt.Sprintf("[%s] ", x.format.id())
	s := ""
	for _, c := range x.chunks {
		s += c.Data
		/*
			if x.format.Segments[idx].Fixed() {
				s += c.Data
			} else {
				s += red(c.Data)
			}
		*/
	}
	return s
}

func newChunk(d string) *Chunk {
	c := Chunk{}
	c.Data = d
	c.IsParam = false
	c.freezed = false
	return &c
}

// Clone is duplicate a Chunk data structure
func (x *Chunk) Clone() *Chunk {
	c := newChunk(x.Data)
	c.IsParam = x.IsParam
	c.freezed = x.freezed
	return c
}

func (x *Chunk) String() string {
	return x.Data
}

func (x *Chunk) equals(chunk *Chunk) bool {
	return x.Data == chunk.Data
}

func (x *Chunk) merge(chunk *Chunk) {
	if x.Data != chunk.Data {
		x.IsParam = true
		x.Data = "*"
	}
}

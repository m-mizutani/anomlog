package anomlog

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Stream is a main module of logseer, should be used mainly as interface
type Stream struct {
	splitter Splitter
	model    Model
}

// NewStream is a constructor
func NewStream() Stream {
	s := Stream{}
	s.splitter = NewSimpleSplitter()
	s.model = &SimpleModel{}
	return s
}

// Read method import log data to a model
func (x *Stream) Read(text string) *Format {
	log := newLog(text, x.splitter)
	fmt := x.model.read(log)
	return fmt
}

// Count returns number of logs which are already analyzed
func (x *Stream) Count() int {
	return x.model.count()
}

// Save command stores model data to file system
func (x *Stream) Save(fpath string) error {
	data, err := json.Marshal(x.model)
	if err != nil {
		return errors.Wrap(err, "Json dump error")
	}

	f, err := os.Create(fpath)
	if err != nil {
		return errors.Wrap(err, "Model file open error: "+fpath)
	}
	defer f.Close()

	f.Write(data)
	return nil
}

// DumpModel returns encoded model data
func (x *Stream) DumpModel() []byte {
	data, err := json.Marshal(x.model)
	if err != nil { // be must encodable
		panic(err)
	}

	return data
}

// Load command imports model data from file system
func (x *Stream) Load(fpath string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return errors.Wrap(err, "Model file open error: "+fpath)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "Fail to read model data: "+fpath)
	}

	model := &SimpleModel{}
	err = json.Unmarshal(data, model)
	if err != nil {
		return errors.Wrap(err, "Fail to load model data:"+fpath)
	}

	x.model = model

	return nil
}

// Formats returns list of Format structure from model.
func (x *Stream) Formats() []*Format {
	return x.model.formats()
}

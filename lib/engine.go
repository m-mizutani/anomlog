package anomlog

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Engine is a main module of logseer, should be used mainly as interface
type Engine struct {
	splitter Splitter
	model    Model
}

// NewEngine is a constructor
func NewEngine() Engine {
	s := Engine{}
	s.splitter = NewSimpleSplitter()
	s.model = &SimpleModel{}
	return s
}

// Read method import log data to a model
func (x *Engine) Read(text string) *Format {
	log := newLog(text, x.splitter)
	fmt := x.model.read(log)
	return fmt
}

// Count returns number of logs which are already analyzed
func (x *Engine) Count() int {
	return x.model.count()
}

// Save command stores model data to file system
func (x *Engine) Save(fpath string) error {
	data := x.DumpModel()

	f, err := os.Create(fpath)
	if err != nil {
		return errors.Wrap(err, "Model file open error: "+fpath)
	}
	defer f.Close()

	f.Write(data)
	return nil
}

// DumpModel returns encoded model data
func (x *Engine) DumpModel() []byte {
	data, err := x.model.dump()
	if err != nil { // be must encodable
		panic(err)
	}

	return data
}

// Load command imports model data from file system
func (x *Engine) Load(fpath string) error {
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
func (x *Engine) Formats() []*Format {
	return []*Format{}
}

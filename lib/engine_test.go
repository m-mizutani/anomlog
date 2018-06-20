package anomlog_test

import (
	"os"
	"testing"

	"io/ioutil"

	anomlog "github.com/m-mizutani/anomlog/lib"
	"github.com/stretchr/testify/assert"
)

func TestEngine1(t *testing.T) {
	engine := anomlog.NewEngine()
	f1 := engine.Read("abc def cde")
	f2 := engine.Read("xzv ckf afe")
	assert.NotEqual(t, f1, f2)
	assert.Equal(t, engine.Count(), 2)
}

func TestEngine2(t *testing.T) {
	engine := anomlog.NewEngine()
	f1 := engine.Read("a b c d x e")
	f2 := engine.Read("a b c d y e")
	assert.Equal(t, f1, f2)
	assert.Equal(t, engine.Count(), 2)
}

func TestEngineSaveModel(t *testing.T) {
	engine := anomlog.NewEngine()
	engine.Read("a b c d x e")
	engine.Read("a b c d y e")
	engine.Read("x y z c")
	engine.Read("x y z a")

	assert.Equal(t, 4, engine.Count())

	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	err = engine.Save(f.Name())
	assert.Nil(t, err)

	s2 := anomlog.NewEngine()
	assert.Equal(t, 0, s2.Count())
	s2.Load(f.Name())
	assert.Equal(t, 4, s2.Count())
}

func TestDumpModel(t *testing.T) {
	engine := anomlog.NewEngine()
	engine.Read("a b c d x e")
	engine.Read("a b c d y e")
	engine.Read("x y z c")
	engine.Read("x y z a")

	data := engine.DumpModel()
	assert.NotEqual(t, 0, len(data))

	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	err = engine.Save(f.Name())
	assert.Nil(t, err)
	f.Close()

	rf, err := os.Open(f.Name())
	assert.Nil(t, err)
	fileData, err := ioutil.ReadAll(rf)
	assert.Nil(t, err)

	assert.Equal(t, fileData, data)
}

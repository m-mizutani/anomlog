package anomlog_test

import (
	"os"
	"testing"

	"io/ioutil"

	anomlog "github.com/m-mizutani/anomlog/lib"
	"github.com/stretchr/testify/assert"
)

func TestStream1(t *testing.T) {
	stream := anomlog.NewStream()
	f1 := stream.Read("abc def cde")
	f2 := stream.Read("xzv ckf afe")
	assert.NotEqual(t, f1, f2)
	assert.Equal(t, stream.Count(), 2)
}

func TestStream2(t *testing.T) {
	stream := anomlog.NewStream()
	f1 := stream.Read("a b c d x e")
	f2 := stream.Read("a b c d y e")
	assert.Equal(t, f1, f2)
	assert.Equal(t, stream.Count(), 2)
}

func TestStreamSaveModel(t *testing.T) {
	stream := anomlog.NewStream()
	stream.Read("a b c d x e")
	stream.Read("a b c d y e")
	stream.Read("x y z c")
	stream.Read("x y z a")

	assert.Equal(t, 4, stream.Count())

	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	err = stream.Save(f.Name())
	assert.Nil(t, err)

	s2 := anomlog.NewStream()
	assert.Equal(t, 0, s2.Count())
	s2.Load(f.Name())
	assert.Equal(t, 4, s2.Count())
}

func TestDumpModel(t *testing.T) {
	stream := anomlog.NewStream()
	stream.Read("a b c d x e")
	stream.Read("a b c d y e")
	stream.Read("x y z c")
	stream.Read("x y z a")

	data := stream.DumpModel()
	assert.NotEqual(t, 0, len(data))

	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	err = stream.Save(f.Name())
	assert.Nil(t, err)
	f.Close()

	rf, err := os.Open(f.Name())
	assert.Nil(t, err)
	fileData, err := ioutil.ReadAll(rf)
	assert.Nil(t, err)

	assert.Equal(t, fileData, data)
}

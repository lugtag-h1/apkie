package files_test

import (
	"apkie/helper/files"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestExists(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), strconv.Itoa(rand.Int()))
	if err != nil {
		t.Fail()
	}

	if files.Exists(f.Name()) {
		t.Fail()
	}

	f.Write([]byte{0x13, 0x37})
	if !files.Exists(f.Name()) {
		t.Fail()
	}
}

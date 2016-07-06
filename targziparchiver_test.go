package compressutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTarGzipArchiver(t *testing.T) {
	archiver, err := NewTarGzipArchiver()
	require.NoError(t, err)

	dir, err := ioutil.TempDir("", "test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	dst, err := ioutil.TempDir("", "test")
	require.NoError(t, err)
	defer os.RemoveAll(dst)

	archive := filepath.Join(dir, "archive")
	folder := filepath.Join(dir, "folder")
	os.Mkdir(folder, 0777)
	err = ioutil.WriteFile(filepath.Join(folder, "file"),
		[]byte("hello, world"), 0666)
	require.NoError(t, err)

	err = archiver.PackFrom(dir, archive, "folder")
	require.NoError(t, err)

	err = archiver.Unpack(archive, dst)
	if assert.NoError(t, err) {
		data, err := ioutil.ReadFile(
			filepath.Join(filepath.Join(dst, "folder", "file")))
		if assert.NoError(t, err) {
			assert.EqualValues(t, "hello, world", data)
		}
	}
}

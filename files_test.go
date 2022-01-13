package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {

	files := []string{
		"templates/sample.txt",
		"templates/{{ .Values.key1 }}/{{ .Values.key2 }}.txt",
	}

	t.Run("ReadFile", func(t *testing.T) {
		m := &morph{}
		m.parseArgs("templates", ".", "", "")
		content, err := m.readFile(files[0])
		assert.NoError(t, err, "no error expected")
		assert.Equal(t, "{{ .Values.key3 }}", content, "file content should be read")
	})

	t.Run("FailReadFile", func(t *testing.T) {
		m := &morph{}
		m.parseArgs("templates", ".", "", "")
		_, err := m.readFile("non-existing-file")
		assert.Error(t, err, "read file should fail for non existing file")
	})

	t.Run("GetFile", func(t *testing.T) {
		tmp, _ := ioutil.TempFile("", "")
		tmp.Close()
		os.Remove(tmp.Name())
		_, err := getFile(tmp.Name())
		assert.NoError(t, err, "file should be created when no directory or file exists")
	})

	t.Run("WriteFile", func(t *testing.T) {
		tmp, _ := ioutil.TempFile("", "")
		defer func() {
			os.Remove(tmp.Name())
		}()
		txt := "test"
		err := writeFile(tmp.Name(), txt)
		assert.NoError(t, err, "write to file should succeed")
	})

	t.Run("ListFiles", func(t *testing.T) {
		m := &morph{}
		m.parseArgs("templates", ".", "", "")
		res, err := m.listFiles()
		assert.NoError(t, err, "files need to be listed")
		assert.ElementsMatch(t, files, res, "all files should be listed")
	})

}

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {

	files := []string{
		"sample.txt",
		"value1/value2.txt",
	}

	t.Run("Generate", func(t *testing.T) {
		dst, _ := ioutil.TempDir("", "")
		defer func() {
			os.Remove(dst)
		}()
		m := newMorph()
		err := m.parseArgs("templates", dst, "", "key1=value1,key2=value2")
		assert.NoError(t, err, "argument needs to be parsed")
		err = m.generate()
		assert.NoError(t, err, "generate failed", err)

		var genFiles []string
		err = filepath.Walk(dst, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			genFiles = append(genFiles, strings.TrimPrefix(path, dst+"/"))
			return nil
		})
		assert.NoError(t, err, "generated files cannot be listed")
		assert.ElementsMatch(t, genFiles, files, "generated files name may not be correct")
	})

	t.Run("ReplaceString", func(t *testing.T) {
		m := newMorph()
		err := m.parseArgs("templates", "dst", "", "key1=value1,key2=value2,key3=value3")
		assert.NoError(t, err, "argument needs to be parsed")

		content, _ := m.readFile("templates/sample.txt")
		replaced := m.replaceString(content)
		assert.Equal(t, "value3", replaced, "file content should be replaced")
	})

	t.Run("ReplaceFileName", func(t *testing.T) {
		m := newMorph()
		err := m.parseArgs("templates", "dst", "", "key1=value1,key2=value2")
		assert.NoError(t, err, "argument needs to be parsed")
		replaced := m.replaceFileName("templates/{{ .Values.key1 }}/{{ .Values.key2 }}.txt")
		assert.Equal(t, "dst/value1/value2.txt", replaced, "file name should be replaced")
	})

}

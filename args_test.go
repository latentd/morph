package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgs(t *testing.T) {

	t.Run("ParseParams", func(t *testing.T) {
		m := newMorph()
		err := m.parseArgs("templates", ".", "", "key=value")
		assert.NoError(t, err, "parsing parameter should succeed when given in the form of key=value")
	})

	t.Run("FailParseParamsNotKeyValue", func(t *testing.T) {
		m := newMorph()
		err := m.parseArgs("templates", ".", "", "onlyKeyOrValue")
		assert.Error(t, err, "parsing parameter should fail when parameter is not given in the form of key=value")
	})

	t.Run("FailParseParamsWhenDirectorySame", func(t *testing.T) {
		m := newMorph()
		err := m.parseArgs(".", ".", "", "")
		assert.Error(t, err, "parsing parameter should fail when template-dir and output-dir are the same, resulting in overriting template files.")
	})

}

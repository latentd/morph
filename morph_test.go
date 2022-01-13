package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMorph(t *testing.T) {

	t.Run("Main", func(t *testing.T) {
		d, _ := ioutil.TempDir("", "")
		rc.SetArgs([]string{"-t", "templates", "-o", d})
		defer func() {
			os.Remove(d)
		}()

		err := rc.Execute()
		assert.NoError(t, err, "no error expected when running")
	})

}

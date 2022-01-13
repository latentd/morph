package main

import (
	"fmt"
	"strings"
)

func (m *morph) parseArgs(templateDir, outputDir, repo, set string) error {
	m.templateDir = templateDir
	m.outputDir = outputDir
	m.repo = repo
	if m.templateDir == "" && m.repo == "" {
		return fmt.Errorf("either template-dir or repo needs to be specified")
	}
	if m.repo == "" && m.templateDir == m.outputDir {
		return fmt.Errorf("template-dir and output-dir needs to be different")
	}
	for _, arg := range strings.Split(set, ",") {
		if arg == "" {
			continue
		}
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			return fmt.Errorf("parameter needs to be specified as key=value. %s", arg)
		}
		m.params[kv[0]] = kv[1]
	}
	return nil
}

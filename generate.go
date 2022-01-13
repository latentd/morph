package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func (m *morph) generate() error {
	files, err := m.listFiles()
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}
	for _, filename := range files {
		content, err := m.readFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if err = writeFile(
			m.replaceFileName(filename),
			m.replaceString(content),
		); err != nil {
			return fmt.Errorf("failed to save replaced file: %w", err)
		}
	}
	return nil
}

func (m *morph) replaceString(str string) string {
	for k, v := range m.params {
		ph := fmt.Sprintf("{{ %s.%s }}", PREFIX, k)
		str = strings.Replace(str, ph, v, -1)
	}
	return str
}

func (m *morph) replaceFileName(filename string) string {
	return filepath.Join(m.outputDir, strings.TrimPrefix(m.replaceString(filename), m.templateDir))
}

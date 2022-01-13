package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func (m *morph) readFile(filename string) (string, error) {
	if m.repo != "" {
		file, err := m.fs.Open(filename)
		if err != nil {
			return "", err
		}
		buff := new(bytes.Buffer)
		_, err = io.Copy(buff, file)
		return string(buff.Bytes()), err
	}
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getFile(filename string) (*os.File, error) {
	dir, _ := filepath.Split(filename)
	os.MkdirAll(dir, os.ModePerm)
	fp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return fp, nil
}

func writeFile(filename string, txt string) error {
	fp, err := getFile(filename)
	if err != nil {
		return fmt.Errorf("failed to get file: %w", err)
	}
	defer fp.Close()
	_, err = fmt.Fprintf(fp, txt)
	if err != nil {
		return err
	}
	return nil
}

func (m *morph) listFiles() ([]string, error) {
	var filenames []string
	var err error
	if m.repo != "" {
		rfs, err := m.listRepoFiles()
		if err != nil {
			return nil, err
		}
		for _, rf := range rfs {
			if strings.HasPrefix(rf, m.templateDir) {
				filenames = append(filenames, rf)
			}
		}
	} else {
		err = filepath.Walk(m.templateDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			filenames = append(filenames, path)
			return nil
		})
	}
	return filenames, err
}

func (m *morph) listRepoFiles() ([]string, error) {
	var filenames []string
	r, err := git.Clone(memory.NewStorage(), m.fs, &git.CloneOptions{
		URL:   m.repo,
		Depth: 1,
	})
	if err != nil {
		return nil, err
	}
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}
	tree.Files().ForEach(func(f *object.File) error {
		filenames = append(filenames, f.Name)
		return nil
	})
	return filenames, nil
}

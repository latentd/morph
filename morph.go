package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/spf13/cobra"
)

var (
	PREFIX = ".Values"
)

type morph struct {
	templateDir string
	outputDir   string
	repo        string
	fs          billy.Filesystem
	params      map[string]string
}

func newMorph() *morph {
	return &morph{
		fs:     memfs.New(),
		params: make(map[string]string),
	}
}

func (m *morph) run() error {
	if err := m.parseArgs(templateDir, outputDir, repo, set); err != nil {
		return fmt.Errorf("failed to validate args: %w", err)
	}
	if err := m.generate(); err != nil {
		return fmt.Errorf("failed to generate template: %w", err)
	}
	return nil
}

var rc = &cobra.Command{
	Use:   "morph",
	Short: "Morph is a simple CLI tool to generate files from templates",
	Long: `A CLI tool to generate files from templates.
Documentation is available at https://pkg.go.dev/github.com/latentd/morph`,
	Run: func(c *cobra.Command, args []string) {
		m := newMorph()
		if err := m.run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var templateDir, outputDir, repo, set string

func init() {
	rc.Flags().StringVarP(&templateDir, "template-dir", "t", "", "template directory")
	rc.Flags().StringVarP(&outputDir, "output-dir", "o", ".", "output directory")
	rc.Flags().StringVarP(&repo, "repo", "r", "", "github repository")
	rc.Flags().StringVar(&set, "set", "", "parameter")
}

func main() {
	if err := rc.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

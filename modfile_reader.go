package gobump

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/mod/modfile"
)

func ReadWriter(path, target string) error {

	f, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	parsedFile, err := modfile.Parse(path, content, func(path, version string) (string, error) {
		return version, nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse file %s: %w", path, err)
	}

	if err := parsedFile.AddGoStmt(target); err != nil {
		return fmt.Errorf("cannot add Go statement (path: %s): %w", path, err)
	}

	content, err = parsedFile.Format()
	if err != nil {
		return fmt.Errorf("failed to format file %s: %w", path, err)
	}

	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file content %s: %w", path, err)
	}

	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek file %s: %w", path, err)
	}

	if n, err := f.Write(content); err != nil || n != len(content) {
		return fmt.Errorf("failed to write entire content to file %s: %w", path, err)
	}

	return f.Sync()
}

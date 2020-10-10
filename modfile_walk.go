package gobump

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/mod/modfile"
)

// SkipFile is a function that returns true whether the given file path must be skipped.
type SkipFile func(path string) bool

// WalkFunc is a function that applies transformation to the given file.
//
// It takes the file path and the target version.
type WalkFunc func(path, target string) error

func ApplyTarget(target, basePath string, walkFunc WalkFunc, skips ...SkipFile) error {

	if err := validateTarget(target); err != nil {
		return err
	}

	return filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		return walk(path, info, err, target, walkFunc, skips...)
	})
}

type dirDetector interface {
	IsDir() bool
}

type dirDetectorFunc func() bool

func (d dirDetectorFunc) IsDir() bool {
	return d()
}

func walk(path string, info dirDetector, err error, target string, walkFunc WalkFunc, skips ...SkipFile) error {

	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	for _, s := range skips {
		if s(path) {
			return nil
		}
	}

	log.Println("Applying walkFunc to path", path)
	return walkFunc(path, target)
}

func validateTarget(target string) error {
	if !modfile.GoVersionRE.MatchString(target) {
		return fmt.Errorf("invalid target version %s", target)
	}
	return nil
}

func NonGoMod(path string) bool {
	return !strings.HasSuffix(path, "go.mod")
}

func Vendor(path string) bool {
	return strings.Contains(path, "vendor")
}

func Exclude(regex *regexp.Regexp) SkipFile {
	return regex.MatchString
}

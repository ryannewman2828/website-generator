package utility

import (
	"io/ioutil"
	"strings"

	"github.com/xchapter7x/lo"
)

// ReadFilesAndDirectories - Reads in the non-private files from the `path` directory
func ReadFilesAndDirectories(path string) ([]string, error) {
	lo.G.Info("Reading the files in the", path, "directory")
	foundFiles, err := ioutil.ReadDir("./")
	files := []string{}
	if err != nil {
		lo.G.Error(err.Error())
		return []string{}, err
	}
	for _, f := range foundFiles {
		if !strings.HasPrefix(f.Name(), ".") && len(f.Name()) > 0 {
			lo.G.Info("Found file:", f.Name())
			files = append(files, f.Name())
		}
	}
	return files, err
}

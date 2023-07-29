package filesystem

import (
	"backup/constants"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type LocalClient struct {
}

func (lc LocalClient) GetFileNames(path string) []string {
	var result []string
	var adjustedPath string = path
	if !strings.HasSuffix(path, constants.Separator) {
		adjustedPath = path + constants.Separator
	}

	entries, err := os.ReadDir(adjustedPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}
		if e.IsDir() {
			result = append(result, lc.GetFileNames(adjustedPath+e.Name())...)
		} else {
			result = append(result, adjustedPath+e.Name())
		}
	}

	return result
}

func (lc LocalClient) ValidatePath(path string) string {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if !info.IsDir() {
		panic(errors.New("Path provided must be a Folder."))
	}

	if !strings.HasSuffix(path, constants.Separator) {
		return path + constants.Separator
	}

	return path
}

func (lc LocalClient) ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func (lc LocalClient) Close() {
}
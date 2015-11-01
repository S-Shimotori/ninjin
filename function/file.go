package function

import (
	"io/ioutil"
)

func listFiles(directoryPath string) ([]string, error) {
	result := []string{}
	files, readError := ioutil.ReadDir(directoryPath)
	if readError != nil {
		return result, readError
	}

	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, readError
}

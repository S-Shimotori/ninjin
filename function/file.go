package function

import (
	"io/ioutil"
)

func listFiles(directoryPath string) ([]string, error) {
	result := []string{}
	files, error := ioutil.ReadDir(directoryPath)
	if error != nil {
		return result, error
	}

	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, error
}

package function

import (
	"io/ioutil"
)

func listApplications(directoryPath string) ([]string, error) {
	result := []string{}
	files, readError := ioutil.ReadDir(directoryPath)
	if readError != nil {
		return result, readError
	}

	for _, file := range files {
		if isApplicationDirectory(directoryPath + "/" + file.Name()) {
			result = append(result, file.Name())
		}
	}

	return result, readError
}

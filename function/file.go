package function

import (
	"io/ioutil"
)

const ApplicationsPath string = "/Applications/"

func GenerateFullPathForFileInApplications(fileName string) string {
	return ApplicationsPath + "/" + fileName
}

func ListApplications(directoryPath string) ([]string, error) {
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

package function

import (
	"io/ioutil"
)

const ApplicationsPath string = "/Applications/"

func GenerateFullPathForFileInApplications(fileName string) string {
	return ApplicationsPath + "/" + fileName
}

func ListXcodes(rootPath string) ([]string, error) {
	result := []string{}
	files, readError := ioutil.ReadDir(rootPath)
	if readError != nil {
		return result, readError
	}

	for _, file := range files {
		filePath := rootPath + "/" + file.Name()
		if isApplicationDirectory(filePath) && isXcode(filePath) {
			result = append(result, file.Name())
		}
	}

	return result, readError
}

package function

import (
	"../model"
	"os"
	"os/exec"
	"encoding/xml"
)

const applicationsPath string = "/Applications/"
const infoPlistPath string = "/Contents/Info.plist"
const versionPlistPath string = "/Contents/version.plist"
const applicationNameKey string = "CFBundleExecutable"
const shortVersionKey string = "CFBundleShortVersionString"
const productBuildVersionKey string = "ProductBuildVersion"
const plutilCommand string = "plutil"

func generateExtractCommand(key string, plistPath string) []string {
	return []string{"-extract", key, "xml1", plistPath, "-o", "-"}
}

func getApplicationName(appFileName string) (string, error) {
	infoPlistFullPath := applicationsPath + appFileName + infoPlistPath
	_, existError := os.Stat(infoPlistFullPath)
	if existError != nil {
		return "", existError
	}

	options := generateExtractCommand(applicationNameKey, infoPlistFullPath)
	execOut, execError := exec.Command(plutilCommand, options...).Output()
	if execError != nil {
		return "", execError
	}

	parseOut := model.Plist{}
	parseError := xml.Unmarshal(execOut, &parseOut)
	return parseOut.String, parseError
}

func getVersions(appFileName string) (string, string) {
	shortVersion := ""
	buildVersion := ""

	infoPlistFullPath := applicationsPath + appFileName + infoPlistPath
	versionPlistFullPath := applicationsPath + appFileName + versionPlistPath
	_, infoPlistExistError := os.Stat(infoPlistFullPath)
	_, versionPlistExistError := os.Stat(versionPlistFullPath)

	switch {
	case versionPlistExistError == nil:
		buildVersionOptions := generateExtractCommand(productBuildVersionKey, versionPlistFullPath)
		buildVersionExecOut, buildVersionExecError := exec.Command(plutilCommand, buildVersionOptions...).Output()
		if buildVersionExecError == nil {
			buildVersionParseOut := model.Plist{}
			buildVersionParseError := xml.Unmarshal(buildVersionExecOut, &buildVersionParseOut)
			if buildVersionParseError == nil {
				buildVersion = buildVersionParseOut.String
			}
		}
		shortVersionOptions := generateExtractCommand(shortVersionKey, versionPlistFullPath)
		shortVersionExecOut, shortVersionExecError := exec.Command(plutilCommand, shortVersionOptions...).Output()
		if shortVersionExecError == nil {
			shortVersionParseOut := model.Plist{}
			shortVersionParseError := xml.Unmarshal(shortVersionExecOut, &shortVersionParseOut)
			if shortVersionParseError == nil {
				shortVersion = shortVersionParseOut.String
				break
			}
		}

		fallthrough

	case infoPlistExistError == nil:
		shortVersionOptions := generateExtractCommand(shortVersionKey, infoPlistFullPath)
		shortVersionExecOut, shortVersionExecError := exec.Command(plutilCommand, shortVersionOptions...).Output()
		if shortVersionExecError == nil {
			shortVersionParseOut := model.Plist{}
			shortVersionParseError := xml.Unmarshal(shortVersionExecOut, &shortVersionParseOut)
			if shortVersionParseError == nil {
				shortVersion = shortVersionParseOut.String
			}
		}
	}

	return shortVersion, buildVersion
}

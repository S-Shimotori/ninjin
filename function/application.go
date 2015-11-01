package function

import "github.com/S-Shimotori/ninjin/model"
import (
	"os"
	"encoding/xml"
	"strings"
	"io/ioutil"
)

const ApplicationsPath string = "/Applications"
const pathToInfoPlistPath string = "/Contents/Info.plist"
const pathToVersionPlistPath string = "/Contents/version.plist"
const pathToDeveloperDirectoryPath string = "/Contents/Developer"

const applicationNameKey string = "CFBundleExecutable"
const shortVersionKey string = "CFBundleShortVersionString"
const productBuildVersionKey string = "ProductBuildVersion"

const appExtension = ".app"

func isApplicationDirectory(filePath string) bool {
	directoryInfo, directoryExistsError := os.Stat(filePath)
	if directoryExistsError != nil {
		return false
	}
	if !directoryInfo.IsDir() || !strings.HasSuffix(directoryInfo.Name(), appExtension) {
		return false
	}

	_, plistExistsError := os.Stat(filePath + pathToInfoPlistPath)
	if plistExistsError != nil {
		return false
	}

	return true
}

func isXcode(appFilePath string) bool {
	if !isApplicationDirectory(appFilePath) {
		return false
	}

	execOut, execError := execPlutilExtractOutput(applicationNameKey, appFilePath + pathToInfoPlistPath)
	if execError != nil {
		return false
	}

	parseOut := model.Plist{}
	parseError := xml.Unmarshal(execOut, &parseOut)
	if parseError != nil || parseOut.String != "Xcode" {
		return false
	} else {
		return true
	}
}

func IsActiveXcode(appFileFullPath string) bool {
	if !isXcode(appFileFullPath) {
		return false
	}

	execOut, execError := execXcodeSelectPrintOutput()
	if execError != nil {
		return false
	}

	activeDeveloperDirectoryPath := strings.TrimSpace(string(execOut[:]))
	if activeDeveloperDirectoryPath == appFileFullPath + pathToDeveloperDirectoryPath {
		return true
	} else {
		return false
	}
}

func GetVersions(appFilePath string) (string, string) {
	shortVersion := ""
	buildVersion := ""

	infoPlistPath := appFilePath + pathToInfoPlistPath
	versionPlistPath := appFilePath + pathToVersionPlistPath
	_, infoPlistExistError := os.Stat(infoPlistPath)
	_, versionPlistExistError := os.Stat(versionPlistPath)

	switch {
	case versionPlistExistError == nil:
		buildVersionExecOut, buildVersionExecError := execPlutilExtractOutput(productBuildVersionKey, versionPlistPath)
		if buildVersionExecError == nil {
			buildVersionParseOut := model.Plist{}
			buildVersionParseError := xml.Unmarshal(buildVersionExecOut, &buildVersionParseOut)
			if buildVersionParseError == nil {
				buildVersion = buildVersionParseOut.String
			}
		}
		shortVersionExecOut, shortVersionExecError := execPlutilExtractOutput(shortVersionKey, versionPlistPath)
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
		shortVersionExecOut, shortVersionExecError := execPlutilExtractOutput(shortVersionKey, infoPlistPath)
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
		if isXcode(filePath) {
			result = append(result, file.Name())
		}
	}

	return result, readError
}

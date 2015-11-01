package function

import "github.com/S-Shimotori/ninjin/model"
import (
	"os"
	"os/exec"
	"encoding/xml"
	"strings"
)

const pathToInfoPlistPath string = "/Contents/Info.plist"
const pathToVersionPlistPath string = "/Contents/version.plist"
const applicationNameKey string = "CFBundleExecutable"
const shortVersionKey string = "CFBundleShortVersionString"
const productBuildVersionKey string = "ProductBuildVersion"
const plutilCommand string = "plutil"
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
	appName, getNameError := getApplicationName(appFilePath)
	if getNameError != nil || appName != "Xcode" {
		return false
	} else {
		return true
	}
}

func generateExtractCommand(key string, plistPath string) []string {
	return []string{"-extract", key, "xml1", plistPath, "-o", "-"}
}

func getApplicationName(appFilePath string) (string, error) {
	infoPlistPath := appFilePath + pathToInfoPlistPath
	_, existError := os.Stat(infoPlistPath)
	if existError != nil {
		return "", existError
	}

	options := generateExtractCommand(applicationNameKey, infoPlistPath)
	execOut, execError := exec.Command(plutilCommand, options...).Output()
	if execError != nil {
		return "", execError
	}

	parseOut := model.Plist{}
	parseError := xml.Unmarshal(execOut, &parseOut)
	return parseOut.String, parseError
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
		buildVersionOptions := generateExtractCommand(productBuildVersionKey, versionPlistPath)
		buildVersionExecOut, buildVersionExecError := exec.Command(plutilCommand, buildVersionOptions...).Output()
		if buildVersionExecError == nil {
			buildVersionParseOut := model.Plist{}
			buildVersionParseError := xml.Unmarshal(buildVersionExecOut, &buildVersionParseOut)
			if buildVersionParseError == nil {
				buildVersion = buildVersionParseOut.String
			}
		}
		shortVersionOptions := generateExtractCommand(shortVersionKey, versionPlistPath)
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
		shortVersionOptions := generateExtractCommand(shortVersionKey, infoPlistPath)
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

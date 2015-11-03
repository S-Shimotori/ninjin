package function

import "github.com/S-Shimotori/ninjin/model"
import (
	"os"
	"encoding/xml"
	"strings"
	"io/ioutil"
	"errors"
)

const ApplicationsPath string = "/Applications"
const pathToInfoPlistPath string = "/Contents/Info.plist"
const pathToVersionPlistPath string = "/Contents/version.plist"
const PathToDeveloperDirectoryPath string = "/Contents/Developer"

const applicationNameKey string = "CFBundleExecutable"
const shortVersionKey string = "CFBundleShortVersionString"
const productBuildVersionKey string = "ProductBuildVersion"

const appExtension = ".app"
const xcodeAppName = "Xcode"

func getVersions(appFilePath string) (string, string) {
	versionPlistPath := appFilePath + pathToVersionPlistPath

	shortVersionExecOut, shortVersionExecError := execPlutilExtractOutput(shortVersionKey, versionPlistPath)
	buildVersionExecOut, buildVersionExecError := execPlutilExtractOutput(productBuildVersionKey, versionPlistPath)
	if shortVersionExecError != nil || buildVersionExecError != nil {
		return "", ""
	}

	shortVersionParseOut := model.Plist{}
	shortVersionParseError := xml.Unmarshal(shortVersionExecOut, &shortVersionParseOut)
	buildVersionParseOut := model.Plist{}
	buildVersionParseError := xml.Unmarshal(buildVersionExecOut, &buildVersionParseOut)
	if shortVersionParseError != nil || buildVersionParseError != nil {
		return "", ""
	}

	return shortVersionParseOut.String, buildVersionParseOut.String
}

func ListXcodes(rootPath string) ([]model.Xcode, error) {
	result := model.XcodeSlice{}
	files, readError := ioutil.ReadDir(rootPath)
	if readError != nil {
		return result, readError
	}

	for _, file := range files {
		xcode, generateError := GenerateXcode(rootPath + "/" + file.Name())
		if generateError == nil {
			result = append(result, xcode)
		}
	}

	result.Sort()
	return result, readError
}

func GenerateXcode(appPath string) (model.Xcode, error) {
	var xcode model.Xcode

	//existsfile
	appFileInfo, appExistsError := os.Stat(appPath)
	if appExistsError != nil {
		return xcode, appExistsError
	}

	//isAppDirectory?
	if !appFileInfo.IsDir() || !strings.HasSuffix(appFileInfo.Name(), appExtension) {
		return xcode, errors.New("not .app directory")
	}

	//isApplication?
	_, infoExistsError := os.Stat(appPath + pathToInfoPlistPath)
	if infoExistsError != nil {
		return xcode, infoExistsError
	}

	//isXcode?
	executableName, executableNameError := execPlutilExtractOutput(applicationNameKey, appPath + pathToInfoPlistPath)
	if executableNameError != nil {
		return xcode, executableNameError
	}
	_, versionExistsError := os.Stat(appPath + pathToInfoPlistPath)
	if versionExistsError != nil {
		return xcode, versionExistsError
	}

	parsedExecutableName := model.Plist{}
	parseError := xml.Unmarshal(executableName, &parsedExecutableName)
	if parseError != nil {
		return xcode, parseError
	}
	if parsedExecutableName.String != xcodeAppName {
		return xcode, errors.New("not Xcode")
	}

	shortVersion, buildVersion := getVersions(appPath)
	if shortVersion == "" || buildVersion == "" {
		return xcode, errors.New("not Xcode")
	}

	xcode = model.Xcode{
		AppPath: appPath,
		AppName: appFileInfo.Name(),
		ShortVersion: shortVersion,
		ProductBuildVersion: buildVersion,
	}

	return xcode, nil
}

func GetActiveDeveloperDirectoryPath() (string, error) {
	execOut, execError := execXcodeSelectPrintOutput()
	if execError != nil {
		return "", execError
	}

	return strings.TrimSpace(string(execOut[:])), execError
}

package function

import (
	"os/exec"
)

const xcodeSelectCommand string = "xcode-select"

func execXcodeSelectSwitchOutput(developerDirectoryPath string) ([]byte, error) {
	return exec.Command(xcodeSelectCommand, "-s", developerDirectoryPath).Output()
}

func execXcodeSelectPrintOutput() ([]byte, error) {
	return exec.Command(xcodeSelectCommand, "-p").Output()
}

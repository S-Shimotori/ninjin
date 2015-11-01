package function

import (
	"os/exec"
)

const xcodeSelectCommand string = "xcode-select"

func ExecXcodeSelectSwitchOutput(developerDirectoryPath string) ([]byte, error) {
	return exec.Command("sudo", xcodeSelectCommand, "-s", developerDirectoryPath).Output()
}

func execXcodeSelectPrintOutput() ([]byte, error) {
	return exec.Command(xcodeSelectCommand, "-p").Output()
}

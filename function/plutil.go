package function

import (
	"os/exec"
)

const plutilCommand string = "plutil"

func execPlutilExtractOutput(key string, plistPath string) ([]byte, error){
	return exec.Command(plutilCommand, "-extract", key, "xml1", plistPath, "-o", "-").Output()
}

package function
import "os"

func Exists(filePath string) bool {
	_, error := os.Stat(filePath)
	if error == nil {
		return true
	} else {
		return false
	}
}
package steamlocate

import (
	"fmt"
	"os"
)

func PathExists(path string) error {
	info, err := os.Stat(path)
	if err == nil || info.IsDir() {
		return nil
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}

	return err
}

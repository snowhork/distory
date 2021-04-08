package dio

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func ListHistory(distoryRoot string) error {
	dir, err := targetDirectory(distoryRoot)
	if err != nil {
		return err
	}

	historyFilePath := path.Join(dir, historyFileName)
	if _, err := os.Stat(historyFilePath); err != nil {
		return nil
	}

	body, err := ioutil.ReadFile(historyFilePath)
	fmt.Print(string(body))

	return nil
}

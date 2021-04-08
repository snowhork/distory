package dio

import (
	"fmt"
	"os"
	"path"
)

var historyFileName = "dir_history"

func targetDirectory(distoryRoot string) (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(distoryRoot, cur), nil
}

func WriteHistory(distoryRoot, command string) error {
	dir, err := targetDirectory(distoryRoot)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0775); err != nil {
		return err
	}

	historyFilePath := path.Join(dir, historyFileName)

	if _, err := os.Stat(historyFilePath); err != nil {
		f, err := os.Create(historyFilePath)
		if err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(historyFilePath, os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s\n", command)); err != nil {
		return err
	}

	return nil
}

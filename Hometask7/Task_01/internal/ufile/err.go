package ufile

import "os"

const (
	errFilePath = "./assets/errors.txt"
)

func LogError(err error) error {
	file, err := os.OpenFile(errFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(err.Error() + "\n")
	return err
}

func LogErrors(errs []error) error {
	file, err := os.OpenFile(errFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, err := range errs {
		_, err = file.WriteString(err.Error() + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

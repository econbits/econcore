//Copyright (C) 2020  Germán Fuentes Capella

package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Downloads the content of a URL to a file
func Download(fromURL string, toPath string) error {
	resp, err := http.Get(fromURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Received response code: %d", resp.StatusCode)
	}

	file, err := os.Create(toPath)
	if err != nil {
		return err
	} else {
		fmt.Print(file.Name())
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

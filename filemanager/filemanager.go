package filemanager

import (
	"os"
	"bufio"
	"errors"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read content of the file")
	}
	return lines, nil
}

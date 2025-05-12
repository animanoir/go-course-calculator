package filemanager

import (
	"os"
	"encoding/json"
	"bufio"
	"errors"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	
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

// interface{} is like "any" type
func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}
	
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	} 
	
	file.Close()
	
	return nil
}

// function/constructor for the FileManager struct
func New(inputPath, outputPath string) FileManager {
	return FileManager {
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}

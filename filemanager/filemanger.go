package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fmInput *FileManager) Read() ([]string, error) {
	file, err := os.Open(fmInput.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Failed to read file")
	}

	defer file.Close()
	return lines, nil

}

func (fmOutput *FileManager) Write(data string) error {
	file, err := os.Create(fmOutput.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to write to file")
	}

	defer file.Close()
	return nil
}

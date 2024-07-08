package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Filemanager struct {
	InputFile  string
	OutputFile string
}

func New(input, output string) Filemanager {
	return Filemanager{
		InputFile:  input,
		OutputFile: output,
	}
}

func (fm Filemanager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFile)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	scaner := bufio.NewScanner(file)

	lines := []string{}
	for scaner.Scan() {
		lines = append(lines, scaner.Text())
	}
	err = scaner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read the file")
	}
	return lines, nil
}

func (fm Filemanager) WriteJson(data any) error {
	file, err := os.Create(fm.OutputFile)
	defer FileClose(file)
	if err != nil {
		return errors.New("falied to create a file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		//file.Close()
		return errors.New("failed to encode in json")
	}
	//file.Close()
	return nil
}

func FileClose(file *os.File) {
	file.Close()
}

package utils

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"os"
)

// Read json from file
func ReadJson(dir string, target any) error {

	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(target); err != nil {
		return err
	}

	return nil
}

// Write json to file
func WriteJson(dir string, target any) error {

	file, err := os.OpenFile(dir, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(target); err != nil {
		return err
	}

	return nil
}

// Read text from file
func ReadText(dir string) ([]string, error) {

	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sentences := []string{}
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}

	return sentences, nil
}

// Write text to file
func WriteText(dir string, target []string) error {

	file, err := os.OpenFile(dir, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, data := range target {
		writer.WriteString(data + "\n")
	}

	return nil
}

// Reads a csv file
func ReadCsv(path string) ([][]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Write text to file
func WriteCsv(path string, input [][]string) error {

	file, err := os.OpenFile(path, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, data := range input {
		writer.Write(data)
	}

	return nil
}

// Read all files in directory and returns list of filenames
func FilesInDirectory(dir string) ([]string, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fileNameList := []string{}
	for _, file := range files {
		fileNameList = append(fileNameList, file.Name())
	}

	return fileNameList, nil
}

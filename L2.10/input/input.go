package input

import (
	"bufio"
	"fmt"
	"os"
)

// ReadStdin считывает строки из стандартного ввода и возвращает их в виде слайса.
func ReadStdin() ([]string, error) {
	return readLines(os.Stdin)
}

// ReadFile открывает файл по имени и возвращает строки, содержащиеся в нём.
func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error closing file %s: %v\n", filename, cerr)
		}
	}()

	return readLines(file)
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

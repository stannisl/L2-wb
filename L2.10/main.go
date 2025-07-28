package main

import (
	"L2-10/args"
	"L2-10/input"
	"L2-10/sort"
	"fmt"
	"os"
)

func main() {
	flagsSettings, files, err := args.ParseArgs(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, "flag error:", err)
	}

	var lines []string

	if len(files) > 0 {
		for _, file := range files {
			fileLines, err := input.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", file, err)
				continue
			}

			lines = append(lines, fileLines...)
		}
	} else {
		fileLines, err := input.ReadStdin()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		}

		lines = append(lines, fileLines...)
	}

	sortedLines := sort.Lines(lines, flagsSettings)
	for _, line := range sortedLines {
		fmt.Println(line)
	}
}

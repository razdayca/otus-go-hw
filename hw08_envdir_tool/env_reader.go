package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(filepath.Join(".", dir))
	if err != nil {
		return nil, errors.New("cant open directory")
	}

	out := Environment{}

	for _, file := range files {
		if strings.Contains(file.Name(), "=") {
			continue
		}

		openedFile, err := os.Open(filepath.Join(".", dir, file.Name()))
		if err != nil {
			fmt.Println("cant read file")
			continue
		}

		reader := bufio.NewReader(openedFile)
		line, _, err := reader.ReadLine()
		if err != nil {
			openedFile.Close()
			if errors.Is(err, io.EOF) {
				fmt.Println("file is empty")
			}
		}
		trimedLine := bytes.TrimRight(line, " \t")
		validLine := bytes.ReplaceAll(trimedLine, []byte("\x00"), []byte("\n"))
		if len(validLine) == 0 {
			openedFile.Close()
			continue
		}
		out[file.Name()] = EnvValue{Value: string(validLine)}
		openedFile.Close()
	}

	return out, nil
}

package main

import (
	"errors"
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
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	env := make(Environment)
	for _, file := range files {
		value, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return env, err
		}
		if strings.Contains(file.Name(), "=") {
			return nil, errors.New("env file name error")
		}
		str1 := strings.Split(string(value), "\n")[0]
		str2 := strings.ReplaceAll(str1, "\x00", "\n")
		cleanValue := strings.TrimRight(str2, " \t")
		env[file.Name()] = EnvValue{Value: cleanValue, NeedRemove: len(cleanValue) == 0}
	}
	return env, nil
}

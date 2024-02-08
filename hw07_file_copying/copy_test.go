package main

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("file not found", func(t *testing.T) {
		err := Copy("", "", 0, 0)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})
	t.Run("offset exceeds file size", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./unit-test/", 10000, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
		os.RemoveAll("./unit-test/")
	})
	t.Run("file limited size", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./unit-test/", 0, 10)
		files, err := os.ReadDir("./unit-test/")
		if err != nil {
			require.Error(t, err)
		}
		var findFile string
		for _, file := range files {
			findFile = file.Name()
		}
		fileStat, err := os.Stat(filepath.Join("./unit-test/", findFile))
		if err != nil {
			require.Error(t, err)
		}
		require.Equal(t, int64(10), fileStat.Size())
		os.RemoveAll("./unit-test/")
	})
	t.Run("file offset", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./unit-test/", 100, 0)
		files, err := os.ReadDir("./unit-test/")
		if err != nil {
			require.Error(t, err)
		}
		var findFile string
		for _, file := range files {
			findFile = file.Name()
		}
		fileR, err := os.Open(filepath.Join("./unit-test/", findFile))
		if err != nil {
			require.Error(t, err)
		}
		scanner := bufio.NewScanner(fileR)
		var line string
		for i := 1; scanner.Scan(); i++ {
			if i > 1 {
				break
			}
			line = scanner.Text()
		}
		fileR.Close()
		require.Equal(t, "our installation", line)
		os.RemoveAll("./unit-test/")
	})

}

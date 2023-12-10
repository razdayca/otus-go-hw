package copy

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("file not found", func(t *testing.T) {
		//os.Args = append(os.Args, "-from=./testdata/input.txt")
		//os.Args = append(os.Args, " -to=./")
		err := Copy("", "", 0, 0)
		require.ErrorIs(t, err, ErrNotFound)
	})
	t.Run("offset exceeds file size", func(t *testing.T) {
		pwd, _ := os.Getwd()
		err := Copy("-from=/Users/razdajbeden/golang/otus-go-hw/hw07_file_copying/testdata/input.txt", "-to=./", 100000, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize, pwd)
	})

}

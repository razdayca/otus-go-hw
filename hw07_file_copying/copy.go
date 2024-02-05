package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileFromCopy, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	fileStat, _ := os.Stat(fromPath)
	defer fileFromCopy.Close()
	var bar *pb.ProgressBar
	if limit < fileStat.Size() {
		bar = pb.Full.Start64(limit)
	} else {
		bar = pb.Full.Start64(fileStat.Size() - offset)
	}
	barReader := bar.NewProxyReader(fileFromCopy)
	defer bar.Finish()
	err = os.MkdirAll(toPath, 0755)
	if err != nil {
		return errors.New("can not create directory")
	}
	fileToCopy, err := os.CreateTemp(toPath, "hw07_file_copying.*.txt")
	if err != nil {
		return errors.New("can not create file")
	}
	defer fileToCopy.Close()

	if offset > 0 {
		if offset > fileStat.Size() {
			return ErrOffsetExceedsFileSize
		}
		_, err = fileFromCopy.Seek(offset, io.SeekStart)
		if err != nil {
			defer fileToCopy.Close()
		}
	}

	if limit == 0 {
		if _, err := io.Copy(fileToCopy, barReader); err != nil {
			return err
		}
		err := fileToCopy.Sync()
		if err != nil {
			return err
		}
	}
	if _, err := io.CopyN(fileToCopy, barReader, limit); err != nil {
		if errors.Is(err, io.EOF) {
			err := fileToCopy.Sync()
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	err = fileToCopy.Sync()
	if err != nil {
		return err
	}

	return nil
}

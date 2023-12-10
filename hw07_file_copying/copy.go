package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
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
	bar := pb.Full.Start64(fileStat.Size() - limit - offset)
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
	io.Copy(fileToCopy, barReader)
	if offset > 0 {
		if offset > fileStat.Size() {
			return ErrOffsetExceedsFileSize
		}
		_, err = fileFromCopy.Seek(offset, io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
	}

	if limit > 0 && limit < fileStat.Size() {
		if _, err := io.CopyN(fileToCopy, fileFromCopy, limit); err != nil {
			return err
		}
		fileToCopy.Sync()
		return nil
	} else {
		if _, err := io.Copy(fileToCopy, fileFromCopy); err != nil {
			return err
		}
		fileToCopy.Sync()
		return nil
	}
}

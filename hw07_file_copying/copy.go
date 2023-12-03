package main

import (
	"errors"
	"io"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileW, err := os.Open(fromPath)
	if err != nil {
		return errors.New("file not found")
	}
	fileStat, _ := os.Stat(fromPath)
	defer fileW.Close()
	fileR, err := os.CreateTemp(toPath, "hw07_file_copying-")
	if err != nil {
		return errors.New("can not create file")
	}
	defer fileR.Close()

	if offset > 0 {
		if offset > fileStat.Size() {
			return ErrOffsetExceedsFileSize
		}
		_, err = fileW.Seek(offset, io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
	}

	if limit < fileStat.Size() {
		if _, err := io.CopyN(fileR, fileW, limit); err != nil {
			return err
		}
		fileR.Sync()
	} else {
		if _, err := io.Copy(fileR, fileW); err != nil {
			return err
		}

		fileR.Sync()
	}

	return nil
}

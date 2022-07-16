package archive

import (
	"archive/zip"
	"io"
	"os"
)

type IZipWriter interface {
	Create(file string) (io.Writer, error)
	Close() error
}

type ZipWriter struct {
	zipWriter IZipWriter
}

func NewZipWriter(file string) (ArchiveWriter, error) {
	archive, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	zipWriter := zip.NewWriter(archive)
	return &ZipWriter{zipWriter: zipWriter}, nil
}

func (c *ZipWriter) WriteFile(file, zipFile string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := c.zipWriter.Create(zipFile)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, f)
	return err
}

func (c *ZipWriter) Close() error {
	return c.zipWriter.Close()
}

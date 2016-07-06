package compressutil

import (
	"flag"
	"os/exec"
	"path/filepath"
)

var tarExecutable = flag.String("tar-path", "tar", "Path to tar executable")

type TarGzipArchiver struct {
	binary string
}

func NewTarGzipArchiver() (*TarGzipArchiver, error) {
	binary, err := exec.LookPath(*tarExecutable)
	if err != nil {
		return nil, err
	}
	return &TarGzipArchiver{
		binary: binary,
	}, nil
}

func (a *TarGzipArchiver) PackFrom(directory, archive, filename string) error {
	return exec.Command(a.binary, "--create", "--gzip",
		"--file", archive,
		"--directory", directory,
		"--", filename).Run()
}

func (a *TarGzipArchiver) Unpack(archive, destination string) error {
	archive, err := filepath.Abs(archive)
	if err != nil {
		return err
	}
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}
	return exec.Command(a.binary, "--extract", "--gzip",
		"--file", archive,
		"--directory", destination).Run()
}

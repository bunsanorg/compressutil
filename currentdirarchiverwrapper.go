package compressutil

import (
	"path/filepath"
)

type CurrentDirArchiverWrapper struct {
	archiver CurrentDirArchiver
}

func WrapCurrentDirArchiver(archiver CurrentDirArchiver) Archiver {
	return &CurrentDirArchiverWrapper{
		archiver: archiver,
	}
}

func (a *CurrentDirArchiverWrapper) Pack(archive, filename string) error {
	filename, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	archive, err = filepath.Abs(archive)
	if err != nil {
		return err
	}
	return a.archiver.PackFrom(filepath.Dir(filename),
		archive, filepath.Base(filename))
}

func (a *CurrentDirArchiverWrapper) PackContents(archive, root string) error {
	filename, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	archive, err = filepath.Abs(archive)
	if err != nil {
		return err
	}
	return a.archiver.PackFrom(filename, archive, ".")
}

func (a *CurrentDirArchiverWrapper) Unpack(archive, destination string) error {
	return a.archiver.Unpack(archive, destination)
}

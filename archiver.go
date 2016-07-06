package compressutil

//go:generate bunsan-mockgen -gofile=$GOFILE

type Archiver interface {
	Pack(archive, filename string) error
	PackContents(archive, root string) error
	Unpack(archive, destination string) error
}

type CurrentDirArchiver interface {
	PackFrom(directory, archive, filename string) error
	Unpack(archive, destination string) error
}

package compressutil

import (
	"testing"

	"github.com/bunsanorg/compressutil/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCurrentDirArchiveWrapperPack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cdarchiver := mock_compressutil.NewMockCurrentDirArchiver(ctrl)
	archiver := WrapCurrentDirArchiver(cdarchiver)

	cdarchiver.EXPECT().PackFrom("/home/user", "/tmp/archive.tgz", "folder")
	err := archiver.Pack("/tmp/archive.tgz", "/home/user/folder")
	assert.NoError(t, err)

	cdarchiver.EXPECT().PackFrom("/home/user/folder", "/tmp/archive.tgz", ".")
	err = archiver.PackContents("/tmp/archive.tgz", "/home/user/folder")
	assert.NoError(t, err)

	cdarchiver.EXPECT().Unpack("/tmp/archive.tgz", "/home/user/folder")
	err = archiver.Unpack("/tmp/archive.tgz", "/home/user/folder")
	assert.NoError(t, err)
}

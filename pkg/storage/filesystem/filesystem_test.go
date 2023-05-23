package filesystem_test

import (
	"testing"

	"github.com/maciejgaleja/gosimple/pkg/storage/filesystem"
	"github.com/maciejgaleja/gosimple/test"
)

func TestFilesystem(t *testing.T) {
	fs := filesystem.FilesystemStore{"../../../test/"}
	test.DoTest(t, fs)
}

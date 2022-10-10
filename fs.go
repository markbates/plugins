package plugins

import "io/fs"

type FSable interface {
	FileSystem() fs.FS
}

type FSSetable interface {
	SetFileSystem(fs.FS)
}

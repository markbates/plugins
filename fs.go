package plugins

import "io/fs"

// FSable can be implemented to return an fs.FS
type FSable interface {
	FileSystem() fs.FS
}

// FSSetable can be implemented to receive an fs.FS
type FSSetable interface {
	SetFileSystem(fs.FS)
}

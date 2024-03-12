package plugins

import "io/fs"

// FSable can be implemented to return an fs.FS
type FSable interface {
	Plugin
	FileSystem() (fs.FS, error)
}

// FSSetable can be implemented to receive an fs.FS
type FSSetable interface {
	Plugin
	SetFileSystem(fs.FS) error
}

package plugtest

import (
	"fmt"
	"io/fs"
	"sync"

	"github.com/markbates/plugins"
)

var _ plugins.FSSetable = &FSable{}
var _ plugins.FSable = &FSable{}

type FSable struct {
	FS fs.FS

	mu sync.Mutex
}

func (f *FSable) FileSystem() (fs.FS, error) {
	if f == nil {
		return nil, fmt.Errorf("FSable instance cannot be nil")
	}

	if f.FS == nil {
		return nil, fmt.Errorf("no filesystem configured")
	}

	return f.FS, nil
}

func (f *FSable) SetFileSystem(fs fs.FS) error {
	if f == nil {
		return fmt.Errorf("FSable instance cannot be nil")
	}

	if fs == nil {
		return fmt.Errorf("filesystem cannot be nil")
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	f.FS = fs
	return nil
}

func (f *FSable) PluginName() string {
	return fmt.Sprintf("%T", f)
}

package plugcmd

// Describer is called by `Print` and can be
// implemented to print a short, single line,
// description of the plugin. `-h`
type Describer interface {
	Commander
	Description() string
}

type DescriberFn func() string

func (fn DescriberFn) Description() string {
	return fn()
}

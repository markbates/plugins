package plugcmd

// Namer is a command that provides a different name for the
// command than the name of the Plugin
type Namer interface {
	Commander
	CmdName() string
}

type NamerFn func() string

func (fn NamerFn) CmdName() string {
	return fn()
}

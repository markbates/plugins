package plugcmd

// Aliaser is a command that provides aliases for itself
type Aliaser interface {
	Commander
	CmdAliases() []string
}

type AliaserFn func() []string

func (fn AliaserFn) CmdAliases() []string {
	return fn()
}

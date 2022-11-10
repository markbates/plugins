package plugcmd

type Hider interface {
	Commander
	HidePlugin()
}

type HiderFn func()

func (fn HiderFn) HidePlugin() {
	fn()
}

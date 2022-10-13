package plugcmd

type Hider interface {
	HidePlugin()
}

type HiderFn func()

func (fn HiderFn) HidePlugin() {
	fn()
}

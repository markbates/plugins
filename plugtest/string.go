package plugtest

type StringPlugin string

func (s StringPlugin) PluginName() string {
	return string(s)
}

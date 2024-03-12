package plugins

// Plugin is the most basic interface a plugin can implement.
type Plugin interface {
	PluginName() string
}

// Scoper can be implemented to return a slice of plugins that
// are important to the type defining it.
type Scoper interface {
	ScopedPlugins() Plugins
}

// FeederFn is a function that is used to feed plugins
// into a Needer implementation.
type FeederFn func() Plugins

// Feeder can be implemented to return a FeederFn.
type Feeder interface {
	PluginFeeder() FeederFn
}

// Needer can be implemented to receive a Feeder function
// that can be used to gain access to other plugins in the system.
type Needer interface {
	WithPlugins(FeederFn) error
}

// AvailabilityChecker can be implemented to check if a plugin
// is available to be used at the given root.
type AvailabilityChecker interface {
	PluginAvailable(root string) bool
}

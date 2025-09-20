# plugins

A flexible and extensible plugin system for Go applications built with modern Go idioms.

## Features

- 🎯 **Type-safe plugin discovery** using Go generics
- 🔌 **Extensible interfaces** for specialized plugin behavior
- 🏗️ **Builder pattern support** for complex configurations
- 📊 **Structured logging** with slog integration
- 🔍 **Runtime availability checking** for plugins
- ⚡ **Performance optimized** with modern slice operations
- 🧪 **Comprehensive test coverage** with benchmarks

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/markbates/plugins"
)

// Define your plugin
type MyPlugin struct {
    name string
}

func (p MyPlugin) PluginName() string {
    return p.name
}

func main() {
    // Create a plugin collection
    plugs := plugins.Plugins{
        MyPlugin{name: "plugin1"},
        MyPlugin{name: "plugin2"},
    }

    // Validate the collection
    if err := plugs.Validate(); err != nil {
        panic(err)
    }

    // Find plugins by type
    myPlugins := plugins.ByType[MyPlugin](plugs)
    fmt.Printf("Found %d plugins\n", len(myPlugins))

    // Get plugin names
    names := plugs.Names()
    fmt.Printf("Plugin names: %v\n", names)
}
```

## Plugin Interfaces

### Core Interfaces

- **`Plugin`**: Basic interface that all plugins must implement
- **`Scoper`**: Plugins that can return scoped plugin collections
- **`Feeder/Needer`**: Plugin communication and dependency injection
- **`AvailabilityChecker`**: Runtime availability checking

### I/O and Filesystem

- **`IOSetable/IOable`**: I/O configuration and access
- **`FSSetable/FSable`**: Filesystem configuration and access

### Command-line (plugcmd package)

- **`Commander`**: CLI command implementation
- **`Namer`**: Custom naming for commands
- **`Flagger`**: Flag definition and parsing

## Advanced Usage

### Plugin Validation

```go
plugs := plugins.Plugins{
    MyPlugin{name: "plugin1"},
    MyPlugin{name: "plugin2"},
}

// Validate for common issues
if err := plugs.Validate(); err != nil {
    log.Fatalf("Plugin validation failed: %v", err)
}
```

### Runtime Availability

```go
type ConditionalPlugin struct {
    name string
}

func (p ConditionalPlugin) PluginName() string {
    return p.name
}

func (p ConditionalPlugin) PluginAvailable(root string) bool {
    // Custom logic to determine availability
    return true
}

// Filter to only available plugins
available := plugs.Available("/project/root")
```

### I/O Configuration

```go
// Configure I/O for all compatible plugins
io := someIOInstance
if err := plugs.SetStdio(io); err != nil {
    log.Fatalf("Failed to configure I/O: %v", err)
}
```

## Requirements

- Go 1.24 or later
- Modern Go toolchain with generics support

## Dependencies

- `github.com/markbates/iox` - I/O utilities
- `github.com/stretchr/testify` - Testing framework (dev dependency)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
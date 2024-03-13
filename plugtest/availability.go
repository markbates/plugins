package plugtest

import (
	"fmt"

	"github.com/markbates/plugins"
)

var _ plugins.AvailabilityChecker = AvailabilityChecker(false)

type AvailabilityChecker bool

func (a AvailabilityChecker) PluginName() string {
	return fmt.Sprintf("%T", a)
}

func (a AvailabilityChecker) PluginAvailable(root string) bool {
	return bool(a)
}

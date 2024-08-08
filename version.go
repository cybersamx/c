package c

import (
	"fmt"
	"github.com/cybersamx/e"
)

func Version() string {
	return fmt.Sprintf("c: v1.2.0\n%s", e.Version())
}

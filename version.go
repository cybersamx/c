package c

import (
	"fmt"
	"github.com/cybersamx/e"
)

func Version() string {
	return fmt.Sprintf("c: v1.3.0\n%s", e.Version())
}

package a

import (
	"fmt"

	"github.com/jasonphi/gomod_hang/b/v2"
)

const Apple = "apple"

func AppleBadger() string {
	return fmt.Sprintf("%s %d", Apple, b.Badger)
}

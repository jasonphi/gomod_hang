package a

import (
	"fmt"

	"github.com/jasonphi/gomod_hang/b"
)

const Apple = "apple"

func AppleBadger() string {
	return fmt.Sprintf("%s %s", Apple, b.Badger)
}

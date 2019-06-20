package lenconv

import (
	"fmt"
)

type Meter float64
type Foot float64

func (f Foot) String() string { return fmt.Sprintf("%g", f) }
func (m Meter) String() string { return fmt.Sprintf("%g", m) }

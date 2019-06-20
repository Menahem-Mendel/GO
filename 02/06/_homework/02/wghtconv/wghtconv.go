package wghtconv

import (
	"fmt"
)

type Pound float64
type Kilogram float64

func (p Pound) String() string { return fmt.Sprintf("%g", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g", k) }

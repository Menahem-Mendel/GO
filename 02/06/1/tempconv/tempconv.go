// Package tempconv ...
// добвьте в пакет tempconv типы, констант и функции для роботы с температурой по шкале Кельвина, 
// в которой нуль градусов соответствует температуре -273,15°C, а разница теммператур в 1К имеет ту же величину, что и 1°C
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

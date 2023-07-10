// https://go.dev/doc/effective_go#constants
package examples

import "fmt"

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

/*
	Stringify ByteSize value + add unit.

Adding this method onto the `ByteSize` type will change the output of the call
in `Iota()` from the raw float output, to a nicely string formatted output with
the unit attached.
*/
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func Iota() {
	bss := []ByteSize{KB, MB, GB, TB, PB, EB, ZB, YB}
	fmt.Printf(`
Example that prints out values when multiplied against the respective ByteSize.
This uses "iota" to generate a list of constants.
Slice of ByteSizes: %f.
`, bss)
	for _, bs := range bss {
		fmt.Printf("5 * ByteSize = %v.\n", 5*bs)
	}
}

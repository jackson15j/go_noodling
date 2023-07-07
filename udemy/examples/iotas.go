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

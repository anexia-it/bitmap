package bitmap_test

import (
	"fmt"

	"github.com/anexia-it/bitmap"
)

func ExampleSet() {
	m := make([]byte, 4)
	bitmap.Set(m, 0)
	bitmap.Set(m, 31)

	fmt.Printf("%x\n", m)
	// Output: 80000001
}

func ExampleClear() {
	m := []byte{0x80 | 0x40, 0x00, 0x00, 0x01}
	bitmap.Clear(m, 31)
	fmt.Printf("%x\n", m)
	// Output: 40000001
}

func ExampleIsSet() {
	m := []byte{0xf0, 0x0f}

	fmt.Printf("%t %t %t %t %t %t %t %t %t %t %t %t %t %t %t %t",
		bitmap.IsSet(m, 0),
		bitmap.IsSet(m, 1),
		bitmap.IsSet(m, 2),
		bitmap.IsSet(m, 3),
		bitmap.IsSet(m, 4),
		bitmap.IsSet(m, 5),
		bitmap.IsSet(m, 6),
		bitmap.IsSet(m, 7),

		bitmap.IsSet(m, 8),
		bitmap.IsSet(m, 9),
		bitmap.IsSet(m, 10),
		bitmap.IsSet(m, 11),
		bitmap.IsSet(m, 12),
		bitmap.IsSet(m, 13),
		bitmap.IsSet(m, 14),
		bitmap.IsSet(m, 15),
	)
	// Output: true true true true false false false false false false false false true true true true
}

func ExampleMask() {
	m := []byte{0x75}
	mask := []byte{0x5b}

	masked, err := bitmap.Mask(m, mask)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", masked)
	// Output: 51
}

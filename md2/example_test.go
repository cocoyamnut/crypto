package md2_test

import (
	"fmt"
	"io"

	"github.com/cocoyamnut/crypto/md2"
)

func ExampleNew() {
	h := md2.New()
	_, _ = io.WriteString(h, "hello")
	fmt.Printf("%x\n", h.Sum(nil))
	// Output:
	// a9046c73e00331af68917d3804f70655
}

func ExampleSum() {
	digest := md2.Sum([]byte("hello"))
	fmt.Printf("%x\n", digest)
	// Output:
	// a9046c73e00331af68917d3804f70655
}

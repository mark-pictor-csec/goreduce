package crasher

import (
	"unsafe"
	foo "errors"
)

// Crasher just crashes.
func Crasher() {
	_, _ = foo.New(""), unsafe.Sizeof(0)
	panic("foocrash")
}

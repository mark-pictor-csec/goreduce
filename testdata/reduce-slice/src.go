package crasher

func Crasher() {
	a := ""
	_ = a[1:2]
}

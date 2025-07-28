package util

func Assert(condition bool, msg string) {
	if !condition {
		panic("assertion failed: " + msg)
	}
}

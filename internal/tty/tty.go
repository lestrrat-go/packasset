package tty

type fder interface {
	Fd() uintptr
}

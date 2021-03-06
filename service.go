package goshneh

/*
#cgo LDFLAGS: -L/usr/local/lib -lavahi-client -lavahi-common
#include "service.h"
*/
import "C"

type Service struct {
	Name      string
	Type      string
	Domain    *string
	Host      *string
	Port      uint16
	Collision uint8
}

var ClientFailedCallback func(err error)

//export clientFailedCallback
func clientFailedCallback(err uint8, strerr *C.char) {
	if ClientFailedCallback != nil && err != 0 {
		ClientFailedCallback(constructError(err, strerr))
	}
}

var context C.Context

func Setup() {
	C.setup(&context)
}

func Run() {
	go C.run(&context)
}

func Clean() {
	C.clean(&context)
}

func Quit() {
	C.quit(&context)
}

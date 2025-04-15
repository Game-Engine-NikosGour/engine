package types

//#include "../windows_os/defines.h"
import "C"
import (
	"runtime"
	"unsafe"
)

var pinner = runtime.Pinner{}

type PlatformState struct {
	Internal_state   unsafe.Pointer
	C_platform_state *C.struct_platform_state
}

func NewPlatformState() *PlatformState {
	this := &PlatformState{}
	p := C.struct_platform_state{}
	pinner.Pin(this)
	this.C_platform_state = &p
	return this
}

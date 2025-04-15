package platform

////#cgo CFLAGS: -g -Wall -Wextra
//#include <stdlib.h>
//#include "windows_os/windows_platform.c"
import "C"
import (
	"unsafe"

	platform_types "github.com/NikosGour/Game-Engine/src/lib/platform/types"
)

func Platform_startup(platform_state *platform_types.PlatformState, app_name string, x, y, width, height int32) error {
	str := C.CString(app_name)
	defer C.free(unsafe.Pointer(str))
	C.platform_startup(unsafe.Pointer(platform_state.C_platform_state), str, C.i32(x), C.i32(y), C.i32(width), C.i32(height))
	for {
		b, _ := C.platform_pump_messages(unsafe.Pointer(platform_state.C_platform_state))
		x := C.GoBytes(unsafe.Pointer(&b), 1)
		y := !(x[0] == 0)

		if !y {
			break

		}
	}

	return nil
}

func Platform_shutdown(platform_state *platform_types.PlatformState) {}

func Platform_pump_messages(platform_state *platform_types.PlatformState) {}

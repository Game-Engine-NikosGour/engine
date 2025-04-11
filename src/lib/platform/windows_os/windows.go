// DISABLED   go:build windows_os

package windows_os

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output syscalls.go syscall_definitions.go

import (
	"fmt"

	platform "github.com/NikosGour/Game-Engine/src/lib/platform/types"
	log "github.com/NikosGour/logging/src"
	winapi "golang.org/x/sys/windows"
)

// var (
// 	_kernel32DLL      windows.Handle
// 	_GetModuleHandleA uintptr
// )

type internal_state_windows struct {
	h_instance winapi.Handle
	hwnd       winapi.HWND
}

func Platform_startup_windows(platform_state *platform.PlatformState, app_name string, x, y, width, height int32) error {
	log.Debug("Windows Startup")
	// load_windows_api()

	internal_state := &internal_state_windows{}
	platform_state.Internal_state = internal_state

	handle, err := GetModuleHandleA(nil)
	if err != nil {
		return fmt.Errorf("Couldn't GetModuleHandleA: `%s`", err)
	}
	log.Debug("handle=%#v", handle)

	// unload_windows_api()
	return nil
}

// func load_windows_api() {
// 	var err error
// 	_kernel32DLL, err = windows.LoadLibrary("kernel32.dll")
// 	if err != nil {
// 		log.Error("Cannot load kernel32.dll")
// 	}

// 	_GetModuleHandleA, err = windows.GetProcAddress(_kernel32DLL, "GetModuleHandleA")
// 	if err != nil {
// 		log.Error("Cannot load `GetModuleHandleA`")
// 	}
// }
// func unload_windows_api() {
// 	err := windows.FreeLibrary(_kernel32DLL)
// 	if err != nil {
// 		log.Error("Couldn't free kernel32")
// 	}
// }

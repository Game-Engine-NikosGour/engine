// DISABLED   go:build windows_os

package platform

import (
	"syscall"

	log "github.com/NikosGour/logging/src"
	"golang.org/x/sys/windows"
)

var (
	_kernel32         windows.Handle
	_GetModuleHandleA uintptr
)

type internal_state_windows struct {
	h_instance windows.Handle
	hwnd       windows.HWND
}

func platform_startup_windows(platform_state *PlatformState, app_name string, x, y, width, height int32) error {
	log.Debug("Windows Startup")
	load_windows_api()

	internal_state := &internal_state_windows{}
	platform_state.internal_state = internal_state

	r, _, _ := syscall.SyscallN(_GetModuleHandleA, 0)
	internal_state.h_instance = windows.Handle(r)
	log.Debug("Got Module Handle")

	unload_windows_api()
	return nil
}

func load_windows_api() {
	var err error
	_kernel32, err = windows.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Error("Cannot load kernel32.dll")
	}

	_GetModuleHandleA, err = windows.GetProcAddress(_kernel32, "GetModuleHandleA")
	if err != nil {
		log.Error("Cannot load `GetModuleHandleA`")
	}
}
func unload_windows_api() {
	err := windows.FreeLibrary(_kernel32)
	if err != nil {
		log.Error("Couldn't free kernel32")
	}
}

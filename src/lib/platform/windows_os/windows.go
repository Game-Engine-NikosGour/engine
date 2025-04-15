// DISABLED   go:build windows_os

package windows_os

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output syscalls.go syscall_definitions.go

import "C"
import (
	"fmt"

	platform "github.com/NikosGour/Game-Engine/src/lib/platform/types"
	log "github.com/NikosGour/logging/src"
	"golang.org/x/sys/windows"
)

type internal_state_windows struct {
	h_instance windows.Handle
	hwnd       windows.HWND
}

func win32_process_message(hwnd windows.HWND, msg uint32, w_param uintptr, l_param uintptr) uintptr {
	return 0
}

func Platform_startup_windows_(platform_state *platform.PlatformState, app_name string, x, y, width, height int32) error {
	log.Debug("Windows Startup")

	internal_state := &internal_state_windows{}
	platform_state.Internal_state = internal_state

	handle, err := GetModuleHandleA(nil)
	if err != nil {
		return fmt.Errorf("Couldn't GetModuleHandleA: `%s`", err)
	}
	log.Debug("handle=%#v", handle)
	internal_state.h_instance = handle

	icon, err := LoadIconA(internal_state.h_instance, nil)
	if err != nil {
		return fmt.Errorf("Couldn't Load Icon: `%s`", err)
	}
	var wc WNDCLASSA
	wc.Style = CS_DBLCLCKS
	proc := WNDPROC(win32_process_message)
	wc.WndProc = &proc
	wc.ClsExtra = 0
	wc.WndExtra = 0
	wc.HInstance = internal_state.h_instance
	wc.HIcon = icon
	// cursor, err := LoadCursorA(, nil)

	return nil
}

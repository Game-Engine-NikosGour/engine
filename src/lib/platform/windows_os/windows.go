// DISABLED   go:build windows_os

package windows_os

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output syscalls.go syscall_definitions.go

import (
	"fmt"

	platform "github.com/NikosGour/Game-Engine/src/lib/platform/types"
	log "github.com/NikosGour/logging/src"
	winapi "golang.org/x/sys/windows"
)

type internal_state_windows struct {
	h_instance winapi.Handle
	hwnd       winapi.HWND
}

func Platform_startup_windows(platform_state *platform.PlatformState, app_name string, x, y, width, height int32) error {
	log.Debug("Windows Startup")

	internal_state := &internal_state_windows{}
	platform_state.Internal_state = internal_state

	handle, err := GetModuleHandleA(nil)
	if err != nil {
		return fmt.Errorf("Couldn't GetModuleHandleA: `%s`", err)
	}
	log.Debug("handle=%#v", handle)

	return nil
}

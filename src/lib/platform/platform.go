package platform

import (
	platform "github.com/NikosGour/Game-Engine/src/lib/platform/types"
	"github.com/NikosGour/Game-Engine/src/lib/platform/windows_os"
)

func Platform_startup(platform_state *platform.PlatformState, app_name string, x, y, width, height int32) error {
	return windows_os.Platform_startup_windows(platform_state, app_name, x, y, width, height)
}

func Platform_shutdown(platform_state *platform.PlatformState) {}

func Platform_pump_messages(platform_state *platform.PlatformState) {}

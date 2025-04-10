package platform

type PlatformState struct {
	internal_state interface{}
}

func Platform_startup(platform_state *PlatformState, app_name string, x, y, width, height int32) error {
	return platform_startup_windows(platform_state, app_name, x, y, width, height)
}

func Platform_shutdown(platform_state *PlatformState) {}

func Platform_pump_messages(platform_state *PlatformState) {}

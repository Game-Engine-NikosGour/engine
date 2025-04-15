package windows_os

// run `go generate -v -x ./src/lib/platform/windows_os` to generate all syscalls
//sys GetModuleHandleA(module_name *byte) (handle windows.Handle, err error)
//sys LoadIconA(hInstance windows.Handle, szIconName *byte) (handle windows.Handle, err error) = user32.LoadIconA
//sys LoadCursorA(hInstance windows.Handle, szCursorName *byte) (handle windows.Handle, err error) = user32.LoadCursorA

package windows_os

import (
	"encoding/binary"

	"golang.org/x/sys/windows"
)

type WNDCLASSA struct {
	Style       uint32
	WndProc     *WNDPROC
	ClsExtra    int32
	WndExtra    int32
	HInstance   windows.Handle
	HIcon       windows.Handle
	HCursor     windows.Handle
	HBrush      windows.Handle
	SzMenuName  *byte
	SzClassName *byte
}

type WNDPROC func(hwnd windows.HWND, msg uint32, w_param uintptr, l_param uintptr) uintptr

const (
	CS_DBLCLCKS = 0x0008
)

func InttoByteArray(i int32) (arr [4]byte) {
	binary.LittleEndian.PutUint32(arr[0:4], uint32(i))
	return
}

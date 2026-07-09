//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

func getProcessCPUTime() int64 {
	var creationTime, exitTime, kernelTime, userTime syscall.Filetime
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return 0
	}
	err = syscall.GetProcessTimes(handle, &creationTime, &exitTime, &kernelTime, &userTime)
	if err != nil {
		return 0
	}
	return (int64(kernelTime.HighDateTime)<<32 + int64(kernelTime.LowDateTime) +
		int64(userTime.HighDateTime)<<32 + int64(userTime.LowDateTime)) * 100
}

func getTermSize() (int, int) {
	var info struct {
		Size              struct{ X, Y int16 }
		CursorPosition    struct{ X, Y int16 }
		Attributes        uint16
		Window            struct{ Left, Top, Right, Bottom int16 }
		MaximumWindowSize struct{ X, Y int16 }
	}
	handle, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return 24, 80
	}
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getConsoleScreenBufferInfo := kernel32.NewProc("GetConsoleScreenBufferInfo")
	r, _, _ := getConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&info)))
	if r == 0 {
		return 24, 80
	}
	return int(info.Window.Bottom - info.Window.Top + 1), int(info.Window.Right - info.Window.Left + 1)
}
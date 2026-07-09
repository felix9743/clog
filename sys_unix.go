//go:build !windows

package main

import (
	"syscall"
	"unsafe"
)

func getProcessCPUTime() int64 {
	var usage syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &usage)
	return usage.Utime.Nano() + usage.Stime.Nano()
}

func getTermSize() (int, int) {
	type winsize struct{ Row, Col, X, Y uint16 }
	ws := &winsize{}
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(ws)))
	return int(ws.Row), int(ws.Col)
}
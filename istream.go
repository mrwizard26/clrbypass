package main

import (
	"syscall"
	"unsafe"
)

type _d9a435f2 struct {
	_711ec703 *_80860867
}

type _80860867 struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	_330f8d34	uintptr
	_81da36c3	uintptr
	_67b9604b	uintptr
	_b6be9a8e	uintptr
	_56821026	uintptr
	_67b52a59	uintptr
	_2fc4b891	uintptr
	_9695df2c	uintptr
	_09f5cddf	uintptr
	_62988453	uintptr
	_5c029663	uintptr
}

// Read reads from the stream into a buffer
func (_61de9d6e *_d9a435f2) _330f8d34(_73513484 []byte) (uint32, error) {
	var _4bfab562 uint32
	_8d822e59, _, _ := syscall.SyscallN(_61de9d6e._711ec703._330f8d34,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_73513484[0])),
		uintptr(len(_73513484)),
		uintptr(unsafe.Pointer(&_4bfab562)))

	if _8d822e59 != 0 {
		return 0, syscall.Errno(_8d822e59)
	}
	return _4bfab562, nil
}

// Write writes to the stream
func (_61de9d6e *_d9a435f2) _81da36c3(_73513484 []byte) (uint32, error) {
	var _d0c05010 uint32
	_8d822e59, _, _ := syscall.SyscallN(_61de9d6e._711ec703._81da36c3,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_73513484[0])),
		uintptr(len(_73513484)),
		uintptr(unsafe.Pointer(&_d0c05010)))

	if _8d822e59 != 0 {
		return 0, syscall.Errno(_8d822e59)
	}
	return _d0c05010, nil
}

// Seek moves the stream pointer
func (_61de9d6e *_d9a435f2) _67b9604b(_64966fc7 int64, _69411024 uint32) (uint64, error) {
	var _d5d28bf0 uint64
	_8d822e59, _, _ := syscall.SyscallN(_61de9d6e._711ec703._67b9604b,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(_64966fc7),
		uintptr(_69411024),
		uintptr(unsafe.Pointer(&_d5d28bf0)))

	if _8d822e59 != 0 {
		return 0, syscall.Errno(_8d822e59)
	}
	return _d5d28bf0, nil
}

// Release decrements the reference count
func (_61de9d6e *_d9a435f2) _2819d4ac() uint32 {
	_8d822e59, _, _ := syscall.SyscallN(_61de9d6e._711ec703._2819d4ac, uintptr(unsafe.Pointer(_61de9d6e)))
	return uint32(_8d822e59)
}

package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _54f666f3 struct {
	_319e9646 *_54f666f3
}
type _45a76a21 *_54f666f3

type _7b81e2df struct {
	_4e87a8bd	_54f666f3
	_7c6fe200	uintptr
	_639590b7	uintptr
	_6bb07601	uint32
}

type _d5604117 struct {
	_711ec703	*_412f9b5f
	_dc943a03	uint32
	_f5a2c8a2	uintptr
	_ba71542a	*_7b81e2df
}

type _412f9b5f struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	_d8c6563d	uintptr
	_3b87237b	uintptr
	_efa2efdf	uintptr
}

func _03ce8162(_967e0e2c *_d5604117, _c963af98 *windows.GUID, _27e0e8a3 **uintptr) uintptr {
	if !_50a2c617(_c963af98, &_87874f80) || !_50a2c617(_c963af98, &_46427002) {
		*_27e0e8a3 = nil
		return uintptr(_482e4fcc)
	}
	*_27e0e8a3 = (*uintptr)(unsafe.Pointer(_967e0e2c))

	// Calling AddRef
	syscall.SyscallN(_967e0e2c._711ec703._aca9fb9e, uintptr(unsafe.Pointer(_967e0e2c)))
	return uintptr(_56702a9c)
}

func _78a2a67d(_967e0e2c *_d5604117) uintptr {
	_967e0e2c._dc943a03++
	return uintptr(_967e0e2c._dc943a03)
}

func _2425c871(_967e0e2c *_d5604117) uintptr {
	_967e0e2c._dc943a03--
	if _967e0e2c._dc943a03 == 0 {
		_58ef3fa7(uintptr(unsafe.Pointer(_967e0e2c)))
	}
	return uintptr(_967e0e2c._dc943a03)
}

func _14cca52e(_967e0e2c *_d5604117, _4c5b9bd7 uintptr, _77e28305 uint32, _2313cf19 **uintptr) uintptr {
	_6acf5dbf := windows.NewLazyDLL("kernel32.dll")
	_450810dc := _6acf5dbf.NewProc("HeapAlloc")
	_8cfca4d7, _, _ := _450810dc.Call(_967e0e2c._f5a2c8a2, 0, _4c5b9bd7)
	*_2313cf19 = (*uintptr)(unsafe.Pointer(_8cfca4d7))
	if *_2313cf19 == nil {
		return 0x8007000E
	}

	return _56702a9c
}

func _b1bb160e(_967e0e2c *_d5604117, _4c5b9bd7 uintptr, _77e28305 uint32, _347dc309 *byte, _26e912c2 int, _2313cf19 **uintptr) uintptr {
	_6acf5dbf := windows.NewLazyDLL("kernel32.dll")
	_450810dc := _6acf5dbf.NewProc("HeapAlloc")
	_8cfca4d7, _, _ := _450810dc.Call(_967e0e2c._f5a2c8a2, 0, _4c5b9bd7)
	*_2313cf19 = (*uintptr)(unsafe.Pointer(_8cfca4d7))
	if *_2313cf19 == nil {
		return 0x8007000E
	}

	return _56702a9c
}

func _3e03c4c1(_967e0e2c *_d5604117, _de42c457 uintptr) uintptr {
	_6acf5dbf := windows.NewLazyDLL("kernel32.dll")
	_450810dc := _6acf5dbf.NewProc("HeapFree")
	_450810dc.Call(_967e0e2c._f5a2c8a2, 0, _de42c457)
	return _56702a9c
}

func _64991bd4() *_d5604117 {
	var _b55f5491 _d5604117

	var _6a61c0c1 _412f9b5f
	_6a61c0c1._09f34622 = windows.NewCallback(_03ce8162)
	_6a61c0c1._aca9fb9e = windows.NewCallback(_78a2a67d)
	_6a61c0c1._2819d4ac = windows.NewCallback(_2425c871)
	_6a61c0c1._d8c6563d = windows.NewCallback(_14cca52e)
	_6a61c0c1._3b87237b = windows.NewCallback(_b1bb160e)
	_6a61c0c1._efa2efdf = windows.NewCallback(_3e03c4c1)

	_b55f5491._711ec703 = &_6a61c0c1

	return &_b55f5491

}

package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _42204c58 struct {
	_3c0f7f13	*_e83c9a18
	_54362c53	*_5662f202
	_e564777f	*_3c30a119
	_dc943a03	uint32
}

type _89eab626 struct {
	_3c0f7f13	*_e83c9a18
	_dc943a03	uint32
}

type _e83c9a18 struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	_b1e99289	uintptr
	_86e53a25	uintptr
}

func _432e7883(_967e0e2c *_42204c58, _c963af98 *windows.GUID, _27e0e8a3 **uintptr) uintptr {
	if !_50a2c617(_c963af98, &_87874f80) || !_50a2c617(_c963af98, &_46427002) {
		*_27e0e8a3 = nil
		return uintptr(_482e4fcc)
	}
	*_27e0e8a3 = (*uintptr)(unsafe.Pointer(_967e0e2c))

	// Calling AddRef
	syscall.SyscallN(_967e0e2c._3c0f7f13._aca9fb9e, uintptr(unsafe.Pointer(_967e0e2c)))
	return uintptr(_56702a9c)
}

func _d5035b25(_967e0e2c *_42204c58) uintptr {
	_967e0e2c._dc943a03++
	return uintptr(_967e0e2c._dc943a03)
}

func _1f1ba61f(_967e0e2c *_42204c58) uintptr {
	_967e0e2c._dc943a03--
	if _967e0e2c._dc943a03 == 0 {
		_58ef3fa7(uintptr(unsafe.Pointer(_967e0e2c)))
	}
	return uintptr(_967e0e2c._dc943a03)
}

func _6c0717dd(_967e0e2c *_42204c58, _6df31f60 uint32, _29706ccd *_1532ee38) uintptr {
	return uintptr(_7903a6c8)
}

func _e40ceaea(_967e0e2c *_42204c58, _8f61fa3d *windows.GUID, _a7b2a2be **uintptr) uintptr {
	if _50a2c617(_8f61fa3d, &_80465a68) {
		*_a7b2a2be = (*uintptr)(unsafe.Pointer(_967e0e2c._e564777f))
		return _56702a9c
	}
	if _50a2c617(_8f61fa3d, &_8f71bbda) {
		_c8f74722 := _ee6a3817()
		_5e32e111 := &_c8f74722
		_5e32e111._70b77075 = _967e0e2c._54362c53
		_5e32e111._de15efd3 = nil
		*_a7b2a2be = (*uintptr)(unsafe.Pointer(_5e32e111))
		return _56702a9c
	}
	*_a7b2a2be = nil
	return uintptr(_7903a6c8)
}

// Generate a New Custom HostControl
func _60cb63b3() _42204c58 {
	var _5d97c5f6 _e83c9a18
	var _b55f5491 _42204c58

	_47eaa0ff := windows.NewCallback(_432e7883)
	_28eecdef := windows.NewCallback(_d5035b25)
	_66f402e3 := windows.NewCallback(_1f1ba61f)
	_3d4f4c30 := windows.NewCallback(_6c0717dd)
	_f148c3e3 := windows.NewCallback(_e40ceaea)

	_5d97c5f6._09f34622 = _47eaa0ff
	_5d97c5f6._aca9fb9e = _28eecdef
	_5d97c5f6._2819d4ac = _66f402e3
	_5d97c5f6._86e53a25 = _3d4f4c30
	_5d97c5f6._b1e99289 = _f148c3e3

	_b55f5491._3c0f7f13 = &_5d97c5f6
	_b55f5491._dc943a03 = 0

	return _b55f5491
}

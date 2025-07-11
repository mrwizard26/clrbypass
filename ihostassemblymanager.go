package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _7b147c53 struct {
	_3c0f7f13	*_e0e34cfd
	_de15efd3	*_c6db52a1
	_70b77075	*_5662f202
	_dc943a03	uint32
}

type _e0e34cfd struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	_8473793f	uintptr
	_340769b2	uintptr
}

func _465ff9eb(_967e0e2c *_42204c58, _c963af98 *windows.GUID, _27e0e8a3 **uintptr) uintptr {
	if !_50a2c617(_c963af98, &_87874f80) || !_50a2c617(_c963af98, &_46427002) {
		*_27e0e8a3 = nil
		return uintptr(_482e4fcc)
	}
	*_27e0e8a3 = (*uintptr)(unsafe.Pointer(_967e0e2c))

	// Calling AddRef
	syscall.SyscallN(_967e0e2c._3c0f7f13._aca9fb9e, uintptr(unsafe.Pointer(_967e0e2c)))
	return uintptr(_56702a9c)

}

func _e5b49e26(_967e0e2c *_42204c58) uintptr {
	_967e0e2c._dc943a03++
	return uintptr(_967e0e2c._dc943a03)
}

func _68aa4399(_967e0e2c *_42204c58) uintptr {
	_967e0e2c._dc943a03--
	if _967e0e2c._dc943a03 == 0 {
		_58ef3fa7(uintptr(unsafe.Pointer(_967e0e2c)))
	}
	return uintptr(_967e0e2c._dc943a03)
}

func _affffc74(_967e0e2c *_7b147c53, _037b14f1 **_59e68d53) uintptr {
	*_037b14f1 = nil
	return _56702a9c
}

func _8357a6b9(_967e0e2c *_7b147c53, _4b7d7f45 **_c6db52a1) uintptr {
	_282bc663 := _7ff16972()
	_282bc663._97440d25 = _967e0e2c._70b77075
	_967e0e2c._de15efd3 = &_282bc663
	*_4b7d7f45 = _967e0e2c._de15efd3

	return _56702a9c
}

func _ee6a3817() _7b147c53 {

	var _b55f5491 _7b147c53
	var _5d97c5f6 _e0e34cfd

	_5d97c5f6._09f34622 = windows.NewCallback(_465ff9eb)
	_5d97c5f6._aca9fb9e = windows.NewCallback(_e5b49e26)
	_5d97c5f6._2819d4ac = windows.NewCallback(_68aa4399)
	_5d97c5f6._8473793f = windows.NewCallback(_affffc74)
	_5d97c5f6._340769b2 = windows.NewCallback(_8357a6b9)

	_b55f5491._3c0f7f13 = &_5d97c5f6
	return _b55f5491
}

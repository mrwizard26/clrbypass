package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _bb8daae4 struct {
	_3c0f7f13 *_37dade49
}

type _c6db52a1 struct {
	_3c0f7f13	*_37dade49
	_97440d25	*_5662f202
	_dc943a03	uint32
}

type _37dade49 struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	_fc1bec6b	uintptr
	_0f98d841	uintptr
}

type _33e8cd16 struct {
	_a68aa5de	uint32
	_df5fba22	*uint16
	_f1be5ded	*uint16
	_e3ea4136	uint32
}

type _68e77b0a struct {
	_a68aa5de	uint32
	_df5fba22	*uint16
	_3ad3fee2	*uint16
}

func _60272036(_967e0e2c *_42204c58, _c963af98 *windows.GUID, _27e0e8a3 **uintptr) uintptr {
	if !_50a2c617(_c963af98, &_87874f80) || !_50a2c617(_c963af98, &_46427002) {
		*_27e0e8a3 = nil
		return uintptr(_482e4fcc)
	}
	*_27e0e8a3 = (*uintptr)(unsafe.Pointer(_967e0e2c))

	// Calling AddRef
	syscall.SyscallN(_967e0e2c._3c0f7f13._aca9fb9e, uintptr(unsafe.Pointer(_967e0e2c)))
	return uintptr(_56702a9c)

}

func _757c3999(_967e0e2c *_42204c58) uintptr {
	_967e0e2c._dc943a03++
	return uintptr(_967e0e2c._dc943a03)
}

func _06254b91(_967e0e2c *_42204c58) uintptr {
	_967e0e2c._dc943a03--
	if _967e0e2c._dc943a03 == 0 {
		_58ef3fa7(uintptr(unsafe.Pointer(_967e0e2c)))
	}
	return uintptr(_967e0e2c._dc943a03)
}

func _ebed993b(_967e0e2c *_c6db52a1, _9966b6a1 *_33e8cd16, _30aec7d0 *uint64, _9c171129 *uint64, _cc0d4c90 **_d9a435f2, _ea11e2ff **_d9a435f2) uintptr {
	_8431581c := windows.UTF16PtrToString(_967e0e2c._97440d25._f8673b44)
	_528e01ad := windows.UTF16PtrToString(_9966b6a1._f1be5ded)
	if _8431581c == _528e01ad {
		*_9c171129 = 0
		*_30aec7d0 = 50000

		_1b5d0313 := _61d36a31(_967e0e2c._97440d25._a52f29f4, _967e0e2c._97440d25._a57adb70)
		*_cc0d4c90 = _1b5d0313

		return _56702a9c
	}
	return uintptr(_84d30ed8(2))
}
func _1ee81f0f(_967e0e2c *_c6db52a1, _9966b6a1 *_68e77b0a, _30aec7d0 *uint64, _9c171129 *uint64, _cc0d4c90 **_d9a435f2, _ea11e2ff **_d9a435f2) uintptr {
	return uintptr(_84d30ed8(2))
}

func _7ff16972() _c6db52a1 {

	var _b55f5491 _c6db52a1
	var _5d97c5f6 _37dade49

	_5d97c5f6._09f34622 = windows.NewCallback(_60272036)
	_5d97c5f6._aca9fb9e = windows.NewCallback(_757c3999)
	_5d97c5f6._2819d4ac = windows.NewCallback(_06254b91)
	_5d97c5f6._fc1bec6b = windows.NewCallback(_ebed993b)
	_5d97c5f6._0f98d841 = windows.NewCallback(_1ee81f0f)

	_b55f5491._3c0f7f13 = &_5d97c5f6
	return _b55f5491
}

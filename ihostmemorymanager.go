package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _3c30a119 struct {
	_10ffb06d	*_f0961395
	_dc943a03	uint32
	_8f6eff24	*_d5604117
	_ba71542a	*_7b81e2df
}

type _f0961395 struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	_115f1b0c	uintptr
	_79b84813	uintptr
	_65c2051d	uintptr
	_d5ae9e8d	uintptr
	_f5aeefe3	uintptr
	_ec3da33f	uintptr
	_6a77bfb3	uintptr
	_5727865e	uintptr
	_0a00dfeb	uintptr
	_ec37f7aa	uintptr
}

func _b82aa008(_967e0e2c *_3c30a119, _c963af98 *windows.GUID, _27e0e8a3 **uintptr) uintptr {
	if !_50a2c617(_c963af98, &_87874f80) || !_50a2c617(_c963af98, &_46427002) {
		*_27e0e8a3 = nil
		return uintptr(_482e4fcc)
	}
	*_27e0e8a3 = (*uintptr)(unsafe.Pointer(_967e0e2c))

	// Calling AddRef
	syscall.SyscallN(_967e0e2c._10ffb06d._aca9fb9e, uintptr(unsafe.Pointer(_967e0e2c)))
	return uintptr(_56702a9c)
}

func _7ce53e0c(_967e0e2c *_3c30a119) uintptr {
	_967e0e2c._dc943a03++
	return uintptr(_967e0e2c._dc943a03)
}

func _08d9b9d4(_967e0e2c *_3c30a119) uintptr {
	_967e0e2c._dc943a03--
	if _967e0e2c._dc943a03 == 0 {
		_58ef3fa7(uintptr(unsafe.Pointer(_967e0e2c)))
	}
	return uintptr(_967e0e2c._dc943a03)
}

func _e7375c9b(_967e0e2c *_3c30a119, _d5d2ef30 uint32, _08f42e8f **_d5604117) uintptr {
	_61df7279 := _64991bd4()

	var _dcf99d2c uintptr
	if _d5d2ef30&0x2 != 0 {
		_dcf99d2c = _f721ad6b(0x00040000)
	} else {
		_dcf99d2c = _f721ad6b(0)
	}

	_61df7279._f5a2c8a2 = _dcf99d2c
	_61df7279._ba71542a = _967e0e2c._ba71542a

	*_08f42e8f = _61df7279
	_967e0e2c._8f6eff24 = _61df7279

	return _56702a9c
}

func _cc475277(_967e0e2c *_3c30a119, _70ee2c25 uintptr, _bd651742 uintptr, _bd8c5d71 uint32, _7e1679b8 uint32, _77e28305 uint32, _2313cf19 **uintptr) uintptr {
	_4557ad9b, _ := windows.VirtualAlloc(_70ee2c25, _bd651742, _bd8c5d71, _7e1679b8)
	*_2313cf19 = (*uintptr)(unsafe.Pointer(_4557ad9b))
	fmt.Println("Called VirtualAlloc")
	return _56702a9c
}

func _237365b7(_967e0e2c *_3c30a119, _90829ac6 uintptr, _bd651742 uintptr, _a3f0f739 uint32) uintptr {
	windows.VirtualFree(_90829ac6, _bd651742, _a3f0f739)
	fmt.Println("Called VirtualFree")
	return _56702a9c
}

func _64127a6d(_967e0e2c *_3c30a119, _90829ac6 uintptr, _815820c5 *uintptr, _42939212 uintptr, _203abb6f *uintptr) uintptr {
	_2f858d34, _, _ := syscall.SyscallN(windows.NewLazyDLL("kernel32.dll").NewProc("VirtualQuery").Addr(), _90829ac6, uintptr(unsafe.Pointer(_815820c5)), uintptr(_42939212))
	*_203abb6f = _2f858d34
	fmt.Println("Called VirtualQuery")
	return _56702a9c
}

func _964e22dc(_967e0e2c *_3c30a119, _90829ac6 uintptr, _bd651742 uintptr, _ce17886a uint32, _164b4d07 *uint32) uintptr {
	windows.VirtualProtect(_90829ac6, _bd651742, _ce17886a, _164b4d07)
	fmt.Println("Called VirtualProtect")
	return _56702a9c
}

func _d0635829(_967e0e2c *_3c30a119, _f2d43568 *uint32, _4c2138b1 *uintptr) uintptr {
	*_f2d43568 = 30
	*_4c2138b1 = 100 * 1024 * 1024
	return _56702a9c
}

func _e36f6d64(_967e0e2c *_3c30a119, _90e399fc uintptr) uintptr {
	return _56702a9c
}

func _d5242507(_967e0e2c *_3c30a119, _57a8f587 *uintptr, _0df1a4e3 uintptr) uintptr {
	return _56702a9c
}

func _3bb1b85a(_967e0e2c *_3c30a119, _57a8f587 uintptr, _0df1a4e3 uintptr) uintptr {
	_99a787ba := &_7b81e2df{}
	_99a787ba._7c6fe200 = _57a8f587
	_99a787ba._639590b7 = _0df1a4e3
	_99a787ba._6bb07601 = 3

	(_45a76a21)(unsafe.Pointer(_99a787ba))._319e9646 = (_45a76a21)(unsafe.Pointer(_967e0e2c._ba71542a))._319e9646
	(_45a76a21)(unsafe.Pointer(_967e0e2c._ba71542a))._319e9646 = (_45a76a21)(unsafe.Pointer(_99a787ba))

	fmt.Println("Called VirtualProtect")

	return _56702a9c
}

func _86018406(_967e0e2c *_3c30a119, _57a8f587 *uintptr) uintptr {
	return _56702a9c
}

func _49be272b() *_3c30a119 {

	_fd1c4b44 := &_3c30a119{}
	_5d97c5f6 := &_f0961395{}
	_5d97c5f6._09f34622 = windows.NewCallback(_b82aa008)
	_5d97c5f6._aca9fb9e = windows.NewCallback(_7ce53e0c)
	_5d97c5f6._2819d4ac = windows.NewCallback(_08d9b9d4)
	_5d97c5f6._79b84813 = windows.NewCallback(_cc475277)
	_5d97c5f6._65c2051d = windows.NewCallback(_237365b7)
	_5d97c5f6._d5ae9e8d = windows.NewCallback(_64127a6d)
	_5d97c5f6._5727865e = windows.NewCallback(_d5242507)
	_5d97c5f6._0a00dfeb = windows.NewCallback(_3bb1b85a)
	_5d97c5f6._6a77bfb3 = windows.NewCallback(_e36f6d64)
	_5d97c5f6._ec37f7aa = windows.NewCallback(_86018406)
	_5d97c5f6._115f1b0c = windows.NewCallback(_e7375c9b)

	_fd1c4b44._10ffb06d = _5d97c5f6

	_568d52b9 := &_7b81e2df{}
	(_45a76a21)(unsafe.Pointer(_568d52b9))._319e9646 = nil

	_fd1c4b44._ba71542a = (*_7b81e2df)(unsafe.Pointer(&_568d52b9))
	return _fd1c4b44
}

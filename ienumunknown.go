//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type _856cf15a struct {
	_89dd8d85 *_05fed2b1
}

// IEnumUnknownVtbl Enumerates objects implementing the root COM interface, IUnknown.
// Commonly implemented by a component containing multiple objects. For more information, see IEnumUnknown.
// https://docs.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ienumunknown
type _05fed2b1 struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	// Next Retrieves the specified number of items in the enumeration sequence.
	_319e9646	uintptr
	// Skip Skips over the specified number of items in the enumeration sequence.
	_8a188643	uintptr
	// Reset Resets the enumeration sequence to the beginning.
	_26af8d02	uintptr
	// Clone Creates a new enumerator that contains the same enumeration state as the current one.
	_5c029663	uintptr
}

func (_61de9d6e *_856cf15a) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_856cf15a) _2819d4ac() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// Next retrieves the specified number of items in the enumeration sequence.
// HRESULT Next(
//
//	ULONG    celt,
//	IUnknown **rgelt,
//	ULONG    *pceltFetched
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ienumunknown-next
func (_61de9d6e *_856cf15a) _319e9646(_a3373a3f uint32, _7a69c82a unsafe.Pointer, _4c79aa45 *uint32) (_3df7d845 int, _a30f3ccf error) {
	_587fa871("Entering into ienumunknown.Next()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall6(
		_61de9d6e._89dd8d85._319e9646,
		4,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(_a3373a3f),
		uintptr(_7a69c82a),
		uintptr(unsafe.Pointer(_4c79aa45)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("there was an error calling the IEnumUnknown::Next method:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c && _a0e7c2cc != _4ef33ea8 {
		_a30f3ccf = fmt.Errorf("the IEnumUnknown::Next method method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	_3df7d845 = int(_a0e7c2cc)
	return
}

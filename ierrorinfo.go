//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _14313b5f struct {
	_89dd8d85 *_992d2e3d
}

// IErrorInfoVtbl returns information about an error in addition to the return code.
// It returns the error message, name of the component and GUID of the interface in
// which the error occurred, and the name and topic of the Help file that applies to the error.
// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/ms723041(v=vs.85)
type _992d2e3d struct {
	// QueryInterface Retrieves pointers to the supported interfaces on an object.
	_09f34622	uintptr
	// AddRef Increments the reference count for an interface pointer to a COM object.
	// You should call this method whenever you make a copy of an interface pointer.
	_aca9fb9e	uintptr
	// Release Decrements the reference count for an interface on a COM object.
	_2819d4ac	uintptr
	// GetDescription Returns a text description of the error
	_46b0dcbd	uintptr
	// GetGUID Returns the GUID of the interface that defined the error.
	_873e64eb	uintptr
	// GetHelpContext Returns the Help context ID for the error.
	_bada87c1	uintptr
	// GetHelpFile Returns the path of the Help file that describes the error.
	_0f14b75a	uintptr
	// GetSource Returns the name of the component that generated the error, such as "ODBC driver-name".
	_2b0c8f69	uintptr
}

// GetDescription Returns a text description of the error.
// HRESULT GetDescription (
//
//	BSTR *pbstrDescription);
//
// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/ms714318(v=vs.85)
func (_61de9d6e *_14313b5f) _46b0dcbd() (_c2009685 *string, _a30f3ccf error) {
	_587fa871("Entering into ierrorinfo.GetDescription()...")

	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._46b0dcbd,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_c2009685)),
		0,
	)

	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the IErrorInfo::GetDescription method returned an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the IErrorInfo::GetDescription method method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

// GetGUID Returns the globally unique identifier (GUID) of the interface that defined the error.
// HRESULT GetGUID(
//
//	GUID *pGUID
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ierrorinfo-getguid
func (_61de9d6e *_14313b5f) _873e64eb() (_2ce7c9ae *windows.GUID, _a30f3ccf error) {
	_587fa871("Entering into ierrorinfo.GetGUID()...")

	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._873e64eb,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(_2ce7c9ae)),
		0,
	)

	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the IErrorInfo::GetGUID method returned an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the IErrorInfo::GetGUID method method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

// GetErrorInfo Obtains the error information pointer set by the previous call to SetErrorInfo in the current logical thread.
// HRESULT GetErrorInfo(
//
//	ULONG      dwReserved,
//	IErrorInfo **pperrinfo
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-geterrorinfo
func _6443fdce() (_97d5e24a *_14313b5f, _a30f3ccf error) {
	_587fa871("Entering into ierrorinfo.GetErrorInfo()...")
	_26859f14 := syscall.MustLoadDLL("OleAut32.dll")
	_3dc06f89 := _26859f14.MustFindProc("GetErrorInfo")
	_a0e7c2cc, _, _a30f3ccf := _3dc06f89.Call(0, uintptr(unsafe.Pointer(&_97d5e24a)))
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the OleAu32.GetErrorInfo procedure call returned an error:\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the OleAu32.GetErrorInfo procedure call returned a non-zero HRESULT code: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

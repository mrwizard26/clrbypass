//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _1532ee38 struct {
	_89dd8d85 *_40faaed5
}

// IUnknownVtbl Enables clients to get pointers to other interfaces on a given object through the
// QueryInterface method, and manage the existence of the object through the AddRef and Release methods.
// All other COM interfaces are inherited, directly or indirectly, from IUnknown. Therefore, the three
// methods in IUnknown are the first entries in the vtable for every interface.
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type _40faaed5 struct {
	// QueryInterface Retrieves pointers to the supported interfaces on an object.
	_09f34622	uintptr
	// AddRef Increments the reference count for an interface pointer to a COM object.
	// You should call this method whenever you make a copy of an interface pointer.
	_aca9fb9e	uintptr
	// Release Decrements the reference count for an interface on a COM object.
	_2819d4ac	uintptr
}

// QueryInterface queries a COM object for a pointer to one of its interface;
// identifying the interface by a reference to its interface identifier (IID).
// If the COM object implements the interface, then it returns a pointer to that interface after calling IUnknown::AddRef on it.
// HRESULT QueryInterface(
//
//	REFIID riid,
//	void   **ppvObject
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (_61de9d6e *_1532ee38) _09f34622(_8f61fa3d windows.GUID, _60a5cb92 unsafe.Pointer) error {
	_587fa871("Entering into iunknown.QueryInterface()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._09f34622,
		3,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_8f61fa3d)),	// A reference to the interface identifier (IID) of the interface being queried for.
		uintptr(_60a5cb92),
	)
	if _a30f3ccf != syscall.Errno(0) {
		return fmt.Errorf("the IUknown::QueryInterface method returned an error:\r\n%s", _a30f3ccf)
	}
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the IUknown::QueryInterface method method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return nil
}

// AddRef Increments the reference count for an interface pointer to a COM object.
// You should call this method whenever you make a copy of an interface pointer
// ULONG AddRef();
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (_61de9d6e *_1532ee38) _aca9fb9e() (_56c57a83 uint32, _a30f3ccf error) {
	_587fa871("Entering into iunknown.AddRef()...")
	_8d822e59, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		return 0, fmt.Errorf("the IUnknown::AddRef method returned an error:\r\n%s", _a30f3ccf)
	}
	_a30f3ccf = nil
	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	_56c57a83 = *(*uint32)(unsafe.Pointer(_8d822e59))
	return
}

// Release Decrements the reference count for an interface on a COM object.
// ULONG Release();
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (_61de9d6e *_1532ee38) _2819d4ac() (_56c57a83 uint32, _a30f3ccf error) {
	_587fa871("Entering into iunknown.Release()...")
	_8d822e59, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		return 0, fmt.Errorf("the IUnknown::Release method returned an error:\r\n%s", _a30f3ccf)
	}
	_a30f3ccf = nil
	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	_56c57a83 = *(*uint32)(unsafe.Pointer(_8d822e59))
	return
}

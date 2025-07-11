//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// from mscorlib.tlh

type _6ef9d402 struct {
	_89dd8d85 *_c92d7b5d
}

// AssemblyVtbl is a COM virtual table of functions for the Assembly Class
// https://docs.microsoft.com/en-us/dotnet/api/system.reflection.assembly?view=netframework-4.8
type _c92d7b5d struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	_121da2b5	uintptr
	_68145cfe	uintptr
	_4db65ade	uintptr
	_62863c34	uintptr
	_2a990e1b	uintptr
	_62e37c64	uintptr
	_e5647db8	uintptr
	_5943ce66	uintptr
	_4d9e6df3	uintptr
	_5192989a	uintptr
	_c8670929	uintptr
	_208ed165	uintptr
	_623b227a	uintptr
	_ca293d8e	uintptr
	_0aee8693	uintptr
	_c8ca0ddd	uintptr
	_3899ff5d	uintptr
	_16e75d19	uintptr
	_4e3016b3	uintptr
	_13728ae0	uintptr
	_5abf78af	uintptr
	_e6150148	uintptr
	_70b6bfbd	uintptr
	_70c4fc07	uintptr
	_6f055d9c	uintptr
	_bf94590b	uintptr
	_e591348b	uintptr
	_4f53a52e	uintptr
	_69a89002	uintptr
	_5f89e438	uintptr
	_a46a0358	uintptr
	_be935f1f	uintptr
	_8c2994ff	uintptr
	_08650e8e	uintptr
	_a0a04803	uintptr
	_797d152b	uintptr
	_0f9e0b8c	uintptr
	_c73552b7	uintptr
	_3eea631e	uintptr
	_60f9d4ba	uintptr
	_c663e66b	uintptr
	_40a25b7e	uintptr
	_05c77555	uintptr
	_eadf0a1b	uintptr
	_46b521b3	uintptr
	_1a26d9a0	uintptr
	_b6b1ca44	uintptr
	_ed114c05	uintptr
}

func (_61de9d6e *_6ef9d402) _09f34622(_8f61fa3d *windows.GUID, _60a5cb92 *uintptr) uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._09f34622,
		3,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(_8f61fa3d)),
		uintptr(unsafe.Pointer(_60a5cb92)))
	return _8d822e59
}

func (_61de9d6e *_6ef9d402) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_6ef9d402) _2819d4ac() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// GetEntryPoint returns the assembly's MethodInfo
//
//	 virtual HRESULT __stdcall get_EntryPoint (
//	/*[out,retval]*/ struct _MethodInfo * * pRetVal ) = 0;
//
// https://docs.microsoft.com/en-us/dotnet/api/system.reflection.assembly.entrypoint?view=netframework-4.8#System_Reflection_Assembly_EntryPoint
// https://docs.microsoft.com/en-us/dotnet/api/system.reflection.methodinfo?view=netframework-4.8
func (_61de9d6e *_6ef9d402) GetEntryPoint() (_f613cbe1 *_3bb2fce1, _a30f3ccf error) {
	_587fa871("Entering into assembly.GetEntryPoint()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._ca293d8e,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_f613cbe1)),
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the Assembly::GetEntryPoint method returned an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the Assembly::GetEntryPoint method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

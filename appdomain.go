//go:build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// AppDomain is a Windows COM object interface pointer for the .NET AppDomain class.
// The AppDomain object represents an application domain, which is an isolated environment where applications execute.
// This structure only contains a pointer to the AppDomain's virtual function table
// https://docs.microsoft.com/en-us/dotnet/api/system.appdomain?view=netframework-4.8
type _f5d69006 struct {
	_89dd8d85 *_e7768609
}

// AppDomainVtbl is a Virtual Function Table for the AppDomain COM interface
// The Virtual Function Table contains pointers to the COM IUnkown interface
// functions (QueryInterface, AddRef, & Release) as well as the AppDomain object's methods
// https://docs.microsoft.com/en-us/dotnet/api/system.appdomain?view=netframework-4.8
type _e7768609 struct {
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
	_e414d397	uintptr
	_0f1710fc	uintptr
	_e591348b	uintptr
	_a659869d	uintptr
	_3dc916af	uintptr
	_d5cf005c	uintptr
	_66378de0	uintptr
	_2fcf6dee	uintptr
	_032c3b14	uintptr
	_08eb8626	uintptr
	_78fc551e	uintptr
	_00038b2d	uintptr
	_56a92ef4	uintptr
	_d3484b3e	uintptr
	_6685b734	uintptr
	_6e1ed10a	uintptr
	_c8297711	uintptr
	_53e06a1a	uintptr
	_91fa520c	uintptr
	_5ffcfcbb	uintptr
	_0da4a4db	uintptr
	_53e075d3	uintptr
	_57429460	uintptr
	_3b858c29	uintptr
	_b01bca66	uintptr
	_1ae68143	uintptr
	_3eea631e	uintptr
	_ad9182e2	uintptr
	_60f9d4ba	uintptr
	_50ca1baa	uintptr
	_c663e66b	uintptr
	_8d2bfa1a	uintptr
	_6636318f	uintptr
	_c41b379d	uintptr
	_13500b2b	uintptr
	_12e2552f	uintptr
	_fb260462	uintptr
	_e0903ddb	uintptr
	_15dcb043	uintptr
	_f69a55cd	uintptr
	_820a8a41	uintptr
	_36209a5e	uintptr
	_ddd3ebe6	uintptr
	_e68d535f	uintptr
	_e9165275	uintptr
	_354360c3	uintptr
	_20f04195	uintptr
	_a39f547e	uintptr
	_e8a266b4	uintptr
	_f3900f70	uintptr
	_2a81cee5	uintptr
	_5d6ec093	uintptr
	_eff3cb6c	uintptr
	_2a95fcc4	uintptr
	_b4575885	uintptr
	_d55303ad	uintptr
	_44fc7cfe	uintptr
	_0b7212e9	uintptr
	_3a53d9dc	uintptr
}

// GetAppDomain is a wrapper function that returns an appDomain from an existing ICORRuntimeHost object
func _dcf7a078(_d27a3928 *_f5ea272a) (_d2d3ad38 *_f5d69006, _a30f3ccf error) {
	_587fa871("Entering into appdomain.GetAppDomain()...")
	_e329efcd, _a30f3ccf := _d27a3928._042bf879()
	if _a30f3ccf != nil {
		return
	}
	_a30f3ccf = _e329efcd._09f34622(_978850c1, unsafe.Pointer(&_d2d3ad38))
	return
}

func (_61de9d6e *_f5d69006) _09f34622(_8f61fa3d *windows.GUID, _60a5cb92 *uintptr) uintptr {
	_587fa871("Entering into appdomain.QueryInterface()...")
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._09f34622,
		3,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(_8f61fa3d)),
		uintptr(unsafe.Pointer(_60a5cb92)))
	return _8d822e59
}

func (_61de9d6e *_f5d69006) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_f5d69006) _2819d4ac() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// GetHashCode serves as the default hash function.
// https://docs.microsoft.com/en-us/dotnet/api/system.object.gethashcode?view=netframework-4.8#System_Object_GetHashCode
func (_61de9d6e *_f5d69006) _e5647db8() (int32, error) {
	_587fa871("Entering into appdomain.GetHashCode()...")
	_8d822e59, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._e5647db8,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		return 0, fmt.Errorf("the appdomain.GetHashCode function returned an error:\r\n%s", _a30f3ccf)
	}
	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	return int32(_8d822e59), nil
}

// Load_3 Loads an Assembly into this application domain.
// virtual HRESULT __stdcall Load_3 (
// /*[in]*/ SAFEARRAY * rawAssembly,
// /*[out,retval]*/ struct _Assembly * * pRetVal ) = 0;
// https://docs.microsoft.com/en-us/dotnet/api/system.appdomain.load?view=net-5.0
func (_61de9d6e *_f5d69006) _13500b2b(_1bb6940e *_52965c54) (_fec9970f *_6ef9d402, _a30f3ccf error) {
	_587fa871("Entering into appdomain.Load_3()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._13500b2b,
		3,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(_1bb6940e)),
		uintptr(unsafe.Pointer(&_fec9970f)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		if _a30f3ccf != syscall.Errno(1150) {
			return
		}
	}

	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the appdomain.Load_3 function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil

	return
}

// ToString Obtains a string representation that includes the friendly name of the application domain and any context policies.
// https://docs.microsoft.com/en-us/dotnet/api/system.appdomain.tostring?view=net-5.0#System_AppDomain_ToString
func (_61de9d6e *_f5d69006) _2ae7e806() (_0404f73f string, _a30f3ccf error) {
	_587fa871("Entering into appdomain.ToString()...")
	var _5dab2c4f *string
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._2a990e1b,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_5dab2c4f)),
		0,
	)

	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the AppDomain.ToString method retured an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the AppDomain.ToString method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	_0404f73f = _6b7f1f9f(unsafe.Pointer(_5dab2c4f))
	return
}

func (_61de9d6e *_f5d69006) _c41b379d(_cd916822 string) (_fec9970f *_6ef9d402, _a30f3ccf error) {
	_9d13ff42, _a30f3ccf := _b8bd4b22(_cd916822)
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}
	_a0e7c2cc, _, _a30f3ccf := syscall.SyscallN(
		_61de9d6e._89dd8d85._c41b379d,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(_9d13ff42),
		uintptr(unsafe.Pointer(&_fec9970f)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		if _a30f3ccf != syscall.Errno(1150) {
			return
		}
	}

	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the appdomain.Load_2 function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil

	return
}

//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type _cde9a584 struct {
	_89dd8d85 *_3f70c0b6
}

// ICLRRuntimeHostVtbl provides functionality similar to that of the ICorRuntimeHost interface
// provided in the .NET Framework version 1, with the following changes
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimehost-interface
type _3f70c0b6 struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	// Start Initializes the CLR into a process.
	_3eabf128	uintptr
	// Stop Stops the execution of code by the runtime.
	_30de6460	uintptr
	// SetHostControl sets the host control interface. You must call SetHostControl before calling Start.
	_3b4fa241	uintptr
	// GetCLRControl gets an interface pointer of type ICLRControl that hosts can use to customize
	// aspects of the common language runtime (CLR).
	_0c1b0c60	uintptr
	// UnloadAppDomain Unloads the AppDomain that corresponds to the specified numeric identifier.
	_de8939ee	uintptr
	// ExecuteInAppDomain Specifies the AppDomain in which to execute the specified managed code.
	_ec73eda1	uintptr
	// GetCurrentAppDomainID gets the numeric identifier of the AppDomain that is currently executing.
	_3a117677	uintptr
	// ExecuteApplication used in manifest-based ClickOnce deployment scenarios to specify the application
	// to be activated in a new domain.
	_76f1bec9	uintptr
	// ExecuteInDefaultAppDomain Invokes the specified method of the specified type in the specified assembly.
	_ab9d7077	uintptr
}

// GetICLRRuntimeHost is a wrapper function that takes an ICLRRuntimeInfo object and
// returns an ICLRRuntimeHost and loads it into the current process
func _be1fe961(_4ed69db5 *_8196dea4) (*_cde9a584, error) {
	_587fa871("Entering into iclrruntimehost.GetICLRRuntimeHost()...")
	var _d27a3928 *_cde9a584
	_a30f3ccf := _4ed69db5._d797eaba(_216acd05, _52a3e46d, unsafe.Pointer(&_d27a3928))
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}

	_a30f3ccf = _d27a3928._3eabf128()
	return _d27a3928, _a30f3ccf
}

func (_61de9d6e *_cde9a584) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_cde9a584) _2819d4ac() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// Start Initializes the common language runtime (CLR) into a process.
// HRESULT Start();
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimehost-start-method
func (_61de9d6e *_cde9a584) _3eabf128() error {
	_587fa871("Entering into iclrruntimehost.Start()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._3eabf128,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		return fmt.Errorf("the ICLRRuntimeHost::Start method returned an error:\r\n%s", _a30f3ccf)
	}
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the ICLRRuntimeHost::Start method method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return nil
}

// ExecuteInDefaultAppDomain Calls the specified method of the specified type in the specified managed assembly.
// HRESULT ExecuteInDefaultAppDomain (
//
//	[in] LPCWSTR pwzAssemblyPath,
//	[in] LPCWSTR pwzTypeName,
//	[in] LPCWSTR pwzMethodName,
//	[in] LPCWSTR pwzArgument,
//
// [out] DWORD *pReturnValue
// );
// An LPCWSTR is a 32-bit pointer to a constant string of 16-bit Unicode characters, which MAY be null-terminated.
// Use syscall.UTF16PtrFromString to turn a string into a LPCWSTR
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimehost-executeindefaultappdomain-method
func (_61de9d6e *_cde9a584) _ab9d7077(_b4d24ac5, _2a5baa01, _dde51899, _5f37f2fe *uint16) (_96eb0709 *uint32, _a30f3ccf error) {
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall9(
		_61de9d6e._89dd8d85._ab9d7077,
		6,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(_b4d24ac5)),
		uintptr(unsafe.Pointer(_2a5baa01)),
		uintptr(unsafe.Pointer(_dde51899)),
		uintptr(unsafe.Pointer(_5f37f2fe)),
		uintptr(unsafe.Pointer(_96eb0709)),
		0,
		0,
		0)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeHost::ExecuteInDefaultAppDomain method returned an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeHost::ExecuteInDefaultAppDomain method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

// GetCurrentAppDomainID Gets the numeric identifier of the AppDomain that is currently executing.
// HRESULT GetCurrentAppDomainId(
//
//	[out] DWORD* pdwAppDomainId
//
// );
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimehost-getcurrentappdomainid-method
func (_61de9d6e *_cde9a584) GetCurrentAppDomainID() (_5b8a9aac uint32, _a30f3ccf error) {
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._3a117677,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_5b8a9aac)),
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeHost::GetCurrentAppDomainID method returned an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeHost::GetCurrentAppDomainID method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

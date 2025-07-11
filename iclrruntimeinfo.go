//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _8196dea4 struct {
	_89dd8d85 *_46587c57
}

// ICLRRuntimeInfoVtbl Provides methods that return information about a specific common language runtime (CLR),
// including version, directory, and load status. This interface also provides runtime-specific functionality
// without initializing the runtime. It includes the runtime-relative LoadLibrary method, the runtime
// module-specific GetProcAddress method, and runtime-provided interfaces through the GetInterface method.
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimeinfo-interface
type _46587c57 struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	// GetVersionString Gets common language runtime (CLR) version information associated with a given
	// ICLRRuntimeInfo interface. This method supersedes the GetRequestedRuntimeInfo and GetRequestedRuntimeVersion methods.
	_0a5d6565	uintptr
	// GetRuntimeDirectory Gets the installation directory of the CLR associated with this interface.
	// This method supersedes the GetCORSystemDirectory method.
	_dc8d9292	uintptr
	// IsLoaded Indicates whether the CLR associated with the ICLRRuntimeInfo interface is loaded into a process.
	_b2b0c1ba	uintptr
	// LoadErrorString Translates an HRESULT value into an appropriate error message for the specified culture.
	// This method supersedes the LoadStringRC and LoadStringRCEx methods.
	_8e1ab52a	uintptr
	// LoadLibrary Loads a library from the framework directory of the CLR represented by an ICLRRuntimeInfo interface.
	// This method supersedes the LoadLibraryShim method.
	_cb773f6f	uintptr
	// GetProcAddress Gets the address of a specified function that was exported from the CLR associated with
	// this interface. This method supersedes the GetRealProcAddress method.
	_725ed6c3	uintptr
	// GetInterface Loads the CLR into the current process and returns runtime interface pointers,
	// such as ICLRRuntimeHost, ICLRStrongName and IMetaDataDispenser. This method supersedes all the CorBindTo* functions.
	_d797eaba	uintptr
	// IsLoadable Indicates whether the runtime associated with this interface can be loaded into the current
	// process, taking into account other runtimes that might already be loaded into the process.
	_b2259ac6	uintptr
	// SetDefaultStartupFlags Sets the CLR startup flags and host configuration file.
	_fe718e12	uintptr
	// GetDefaultStartupFlags Gets the CLR startup flags and host configuration file.
	_1641422b	uintptr
	// BindAsLegacyV2Runtime Binds this runtime for all legacy CLR version 2 activation policy decisions.
	_199b68a4	uintptr
	// IsStarted Indicates whether the CLR that is associated with the ICLRRuntimeInfo interface has been started.
	_a0e51459	uintptr
}

// GetRuntimeInfo is a wrapper function to return an ICLRRuntimeInfo from a standard version string
func _97617b08(_71b93ef9 *_12e239b6, _da280012 string) (*_8196dea4, error) {
	_00b52116, _a30f3ccf := syscall.UTF16PtrFromString(_da280012)
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}
	return _71b93ef9._6f0b0691(_00b52116, _dc2ff3cb)
}

func (_61de9d6e *_8196dea4) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_8196dea4) _2819d4ac() uintptr {
	_587fa871("Entering into iclrruntimeinfo.Release()...")
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// GetVersionString gets common language runtime (CLR) version information associated with a given ICLRRuntimeInfo interface.
// HRESULT GetVersionString(
//
//	[out, size_is(*pcchBuffer)] LPWSTR pwzBuffer,
//	[in, out]  DWORD *pcchBuffer);
//
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimeinfo-getversionstring-method
func (_61de9d6e *_8196dea4) _0a5d6565() (_da280012 string, _a30f3ccf error) {
	_587fa871("Entering into iclrruntimeinfo.GetVersion()...")
	// [in, out] Specifies the size of pwzBuffer to avoid buffer overruns. If pwzBuffer is null, pchBuffer returns the required size of pwzBuffer to allow preallocation.
	var _821d434d uint32
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._0a5d6565,
		3,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		uintptr(unsafe.Pointer(&_821d434d)),
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("there was an error calling the ICLRRuntimeInfo::GetVersionString method during preallocation:\r\n%s", _a30f3ccf)
		return
	}
	// 0x8007007a = The data area passed to a system call is too small, expected when passing a nil buffer for preallocation
	if _a0e7c2cc != _56702a9c && _a0e7c2cc != 0x8007007a {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeInfo::GetVersionString method (preallocation) returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}

	_d41dd921 := make([]uint16, 20)

	_a0e7c2cc, _, _a30f3ccf = syscall.Syscall(
		_61de9d6e._89dd8d85._0a5d6565,
		3,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_d41dd921[0])),
		uintptr(unsafe.Pointer(&_821d434d)),
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("there was an error calling the ICLRRuntimeInfo::GetVersionString method:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeInfo::GetVersionString method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	_da280012 = syscall.UTF16ToString(_d41dd921)
	return
}

// GetInterface loads the CLR into the current process and returns runtime interface pointers,
// such as ICLRRuntimeHost, ICLRStrongName, and IMetaDataDispenserEx.
// HRESULT GetInterface(
//
//	[in]  REFCLSID rclsid,
//	[in]  REFIID   riid,
//	[out, iid_is(riid), retval] LPVOID *ppUnk); unsafe pointer of a pointer to an object pointer
//
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimeinfo-getinterface-method
func (_61de9d6e *_8196dea4) _d797eaba(_63ff2525 windows.GUID, _8f61fa3d windows.GUID, _d00a2a30 unsafe.Pointer) error {
	_587fa871("Entering into iclrruntimeinfo.GetInterface()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall6(
		_61de9d6e._89dd8d85._d797eaba,
		4,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_63ff2525)),
		uintptr(unsafe.Pointer(&_8f61fa3d)),
		uintptr(_d00a2a30),
		0,
		0,
	)
	// The syscall returns "The requested lookup key was not found in any active activation context." in the error position
	// TODO Why is this error message returned?
	if _a30f3ccf != syscall.Errno(0) && _a30f3ccf.Error() != "The requested lookup key was not found in any active activation context." {
		return fmt.Errorf("the ICLRRuntimeInfo::GetInterface method returned an error:\r\n%s", _a30f3ccf)
	}
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the ICLRRuntimeInfo::GetInterface method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return nil
}

// BindAsLegacyV2Runtime binds the current runtime for all legacy common language runtime (CLR) version 2 activation policy decisions.
// HRESULT BindAsLegacyV2Runtime ();
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimeinfo-bindaslegacyv2runtime-method
func (_61de9d6e *_8196dea4) _199b68a4() error {
	_587fa871("Entering into iclrruntimeinfo.BindAsLegacyV2Runtime()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._199b68a4,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		return fmt.Errorf("the ICLRRuntimeInfo::BindAsLegacyV2Runtime method returned an error:\r\n%s", _a30f3ccf)
	}
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the ICLRRuntimeInfo::BindAsLegacyV2Runtime method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return nil
}

// IsLoadable indicates whether the runtime associated with this interface can be loaded into the current process,
// taking into account other runtimes that might already be loaded into the process.
// HRESULT IsLoadable(
//
//	[out, retval] BOOL *pbLoadable);
//
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrruntimeinfo-isloadable-method
func (_61de9d6e *_8196dea4) _b2259ac6() (_8760900f bool, _a30f3ccf error) {
	_587fa871("Entering into iclrruntimeinfo.IsLoadable()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._b2259ac6,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_8760900f)),
		0)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeInfo::IsLoadable method returned an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICLRRuntimeInfo::IsLoadable method  returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

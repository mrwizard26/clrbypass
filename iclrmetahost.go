//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Couldnt have done any of this without this SO answer I stumbled on:
// https://stackoverflow.com/questions/37781676/how-to-use-com-component-object-model-in-golang

// ICLRMetaHost Interface from metahost.h
type _12e239b6 struct {
	_89dd8d85 *_16f61f1e
}

// ICLRMetaHostVtbl provides methods that return a specific version of the common language runtime (CLR)
// based on its version number, list all installed CLRs, list all runtimes that are loaded in a specified
// process, discover the CLR version used to compile an assembly, exit a process with a clean runtime
// shutdown, and query legacy API binding.
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrmetahost-interface
type _16f61f1e struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	// GetRuntime gets the ICLRRuntimeInfo interface that corresponds to a particular CLR version.
	// This method supersedes the CorBindToRuntimeEx function used with the STARTUP_LOADER_SAFEMODE flag.
	_6f0b0691	uintptr
	// GetVersionFromFile gets the assembly's original .NET Framework compilation version (stored in the metadata),
	// given its file path. This method supersedes GetFileVersion.
	_010ba1d0	uintptr
	// EnumerateInstalledRuntimes returns an enumeration that contains a valid ICLRRuntimeInfo interface
	// pointer for each CLR version that is installed on a computer.
	_3ce71b97	uintptr
	// EnumerateLoadedRuntimes returns an enumeration that contains a valid ICLRRuntimeInfo interface
	// pointer for each CLR that is loaded in a given process. This method supersedes GetVersionFromProcess.
	_45f10549	uintptr
	// RequestRuntimeLoadedNotification guarantees a callback to the specified function pointer when a
	// CLR version is first loaded, but not yet started. This method supersedes LockClrVersion
	_95867ed1	uintptr
	// QueryLegacyV2RuntimeBinding returns an interface that represents a runtime to which legacy activation policy
	// has been bound, for example by using the useLegacyV2RuntimeActivationPolicy attribute on the <startup> Element
	// configuration file entry, by direct use of the legacy activation APIs, or by calling the
	// ICLRRuntimeInfo::BindAsLegacyV2Runtime method.
	_8670b49e	uintptr
	// ExitProcess attempts to shut down all loaded runtimes gracefully and then terminates the process.
	_13a043fd	uintptr
}

// CLRCreateInstance provides one of three interfaces: ICLRMetaHost, ICLRMetaHostPolicy, or ICLRDebugging.
// HRESULT CLRCreateInstance(
//
//	[in]  REFCLSID  clsid,
//	[in]  REFIID     riid,
//	[out] LPVOID  * ppInterface
//
// );
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/clrcreateinstance-function
func _f5b5f41e(_c3eb91f4, _8f61fa3d windows.GUID) (_f6ddca31 *_12e239b6, _a30f3ccf error) {
	_587fa871("Entering into iclrmetahost.CLRCreateInstance()...")

	if _c3eb91f4 != _43849882 {
		_a30f3ccf = fmt.Errorf("the input Class ID (CLSID) is not supported: %s", _c3eb91f4)
		return
	}

	_e09b1ec7 := syscall.MustLoadDLL("mscoree.dll")
	_5024242b := _e09b1ec7.MustFindProc("CLRCreateInstance")

	// For some reason this procedure call returns "The specified procedure could not be found." even though it works
	_a0e7c2cc, _, _a30f3ccf := _5024242b.Call(
		uintptr(unsafe.Pointer(&_c3eb91f4)),
		uintptr(unsafe.Pointer(&_8f61fa3d)),
		uintptr(unsafe.Pointer(&_f6ddca31)),
	)

	if _a30f3ccf != nil {
		// TODO Figure out why "The specified procedure could not be found." is returned even though everything works fine?
		_587fa871(fmt.Sprintf("the mscoree!CLRCreateInstance function returned an error:\r\n%s", _a30f3ccf))
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the mscoree!CLRCreateInstance function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

func (_61de9d6e *_12e239b6) _09f34622(_8f61fa3d windows.GUID, _60a5cb92 unsafe.Pointer) error {
	_587fa871("Entering into icorruntimehost.QueryInterface()...")
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

func (_61de9d6e *_12e239b6) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_12e239b6) _2819d4ac() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// EnumerateInstalledRuntimes returns an enumeration that contains a valid ICLRRuntimeInfo interface for each
// version of the common language runtime (CLR) that is installed on a computer.
// HRESULT EnumerateInstalledRuntimes (
//
//	[out, retval] IEnumUnknown **ppEnumerator);
//
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrmetahost-enumerateinstalledruntimes-method
func (_61de9d6e *_12e239b6) _3ce71b97() (_8b81c014 *_856cf15a, _a30f3ccf error) {
	_587fa871("Entering into iclrmetahost.EnumerateInstalledRuntimes()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._3ce71b97,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_8b81c014)),
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("there was an error calling the ICLRMetaHost::EnumerateInstalledRuntimes method:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICLRMetaHost::EnumerateInstalledRuntimes method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

// GetRuntime gets the ICLRRuntimeInfo interface that corresponds to a particular version of the common language runtime (CLR).
// This method supersedes the CorBindToRuntimeEx function used with the STARTUP_LOADER_SAFEMODE flag.
// HRESULT GetRuntime (
//
//	[in] LPCWSTR pwzVersion,
//	[in] REFIID riid,
//	[out,iid_is(riid), retval] LPVOID *ppRuntime
//
// );
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/iclrmetahost-getruntime-method
func (_61de9d6e *_12e239b6) _6f0b0691(_00b52116 *uint16, _8f61fa3d windows.GUID) (_e593a24f *_8196dea4, _a30f3ccf error) {
	_587fa871("Entering into iclrmetahost.GetRuntime()...")

	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall6(
		_61de9d6e._89dd8d85._6f0b0691,
		4,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(_00b52116)),
		uintptr(unsafe.Pointer(&_dc2ff3cb)),
		uintptr(unsafe.Pointer(&_e593a24f)),
		0,
		0,
	)

	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("there was an error calling the ICLRMetaHost::GetRuntime method:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICLRMetaHost::GetRuntime method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

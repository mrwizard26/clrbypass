//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type _f5ea272a struct {
	_89dd8d85 *_5751af2e
}

// ICORRuntimeHostVtbl Provides methods that enable the host to start and stop the common language runtime (CLR)
// explicitly, to create and configure application domains, to access the default domain, and to enumerate all
// domains running in the process.
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/icorruntimehost-interface
type _5751af2e struct {
	_09f34622	uintptr
	_aca9fb9e	uintptr
	_2819d4ac	uintptr
	// CreateLogicalThreadState Do not use.
	_10e6e104	uintptr
	// DeleteLogicalThreadSate Do not use.
	_1c3c1278	uintptr
	// SwitchInLogicalThreadState Do not use.
	_4e24c43d	uintptr
	// SwitchOutLogicalThreadState Do not use.
	_07e27fdb	uintptr
	// LocksHeldByLogicalThreadState Do not use.
	_f033829f	uintptr
	// MapFile Maps the specified file into memory. This method is obsolete.
	_d97697aa	uintptr
	// GetConfiguration Gets an object that allows the host to specify the callback configuration of the CLR.
	_f6f25095	uintptr
	// Start Starts the CLR.
	_3eabf128	uintptr
	// Stop Stops the execution of code in the runtime for the current process.
	_30de6460	uintptr
	// CreateDomain Creates an application domain. The caller receives an interface pointer of
	// type _AppDomain to an instance of type System.AppDomain.
	_b887e357	uintptr
	// GetDefaultDomain Gets an interface pointer of type _AppDomain that represents the default domain for the current process.
	_042bf879	uintptr
	// EnumDomains Gets an enumerator for the domains in the current process.
	_b002a8fe	uintptr
	// NextDomain Gets an interface pointer to the next domain in the enumeration.
	_04a3b46f	uintptr
	// CloseEnum Resets a domain enumerator back to the beginning of the domain list.
	_ec417ad3	uintptr
	// CreateDomainEx Creates an application domain. This method allows the caller to pass an
	// IAppDomainSetup instance to configure additional features of the returned _AppDomain instance.
	_6618f1b2	uintptr
	// CreateDomainSetup Gets an interface pointer of type IAppDomainSetup to an AppDomainSetup instance.
	// IAppDomainSetup provides methods to configure aspects of an application domain before it is created.
	_72f6c6cb	uintptr
	// CreateEvidence Gets an interface pointer of type IIdentity, which allows the host to create security
	// evidence to pass to CreateDomain or CreateDomainEx.
	_7faa4227	uintptr
	// UnloadDomain Unloads the specified application domain from the current process.
	_d2f05c82	uintptr
	// CurrentDomain Gets an interface pointer of type _AppDomain that represents the domain loaded on the current thread.
	_167e1a61	uintptr
}

// GetICORRuntimeHost is a wrapper function that takes in an ICLRRuntimeInfo and returns an ICORRuntimeHost object
// and loads it into the current process. This is the "deprecated" API, but the only way currently to load an assembly
// from memory (afaict)
func _0af4f255(_4ed69db5 *_8196dea4) (*_f5ea272a, error) {
	_587fa871("Entering into icorruntimehost.GetICORRuntimeHost()...")
	var _d27a3928 *_f5ea272a
	_a30f3ccf := _4ed69db5._d797eaba(_251588fb, _d40bda53, unsafe.Pointer(&_d27a3928))
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}

	_a30f3ccf = _d27a3928._3eabf128()
	return _d27a3928, _a30f3ccf
}

func (_61de9d6e *_f5ea272a) _09f34622(_8f61fa3d windows.GUID, _60a5cb92 unsafe.Pointer) error {
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

func (_61de9d6e *_f5ea272a) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_f5ea272a) _2819d4ac() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// Start starts the common language runtime (CLR).
// HRESULT Start ();
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/icorruntimehost-start-method
func (_61de9d6e *_f5ea272a) _3eabf128() error {
	_587fa871("Entering into icorruntimehost.Start()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._3eabf128,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		// The system could not find the environment option that was entered.
		// TODO Why is this error message returned?
		_587fa871(fmt.Sprintf("the ICORRuntimeHost::Start method returned an error:\r\n%s", _a30f3ccf))
	}
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the ICORRuntimeHost::Start method method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return nil
}

// GetDefaultDomain gets an interface pointer of type System._AppDomain that represents the default domain for the current process.
// HRESULT GetDefaultDomain (
//
//	[out] IUnknown** pAppDomain
//
// );
// https://docs.microsoft.com/en-us/dotnet/framework/unmanaged-api/hosting/icorruntimehost-getdefaultdomain-method
func (_61de9d6e *_f5ea272a) _042bf879() (_1532ee38 *_1532ee38, _a30f3ccf error) {
	_587fa871("Entering into icorruntimehost.GetDefaultDomain()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._042bf879,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_1532ee38)),
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		// The specified procedure could not be found.
		// TODO Why is this error message returned?
		_587fa871(fmt.Sprintf("the ICORRuntimeHost::GetDefaultDomain method returned an error:\r\n%s", _a30f3ccf))
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICORRuntimeHost::GetDefaultDomain method method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

// CreateDomain Creates an application domain. The caller receives an interface pointer of type _AppDomain to an instance of type System.AppDomain.
// HRESULT CreateDomain (
//
//	[in] LPWSTR    pwzFriendlyName,
//	[in] IUnknown* pIdentityArray,
//	[out] void   **pAppDomain
//
// );
// https://docs.microsoft.com/en-us/previous-versions/dotnet/netframework-4.0/ms164322(v=vs.100)
func (_61de9d6e *_f5ea272a) _b887e357(_ca455e93 *uint16) (_cb204c04 *_f5d69006, _a30f3ccf error) {
	_587fa871("Entering into icorruntimehost.CreateDomain()...")
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall6(
		_61de9d6e._89dd8d85._b887e357,
		4,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(_ca455e93)),	// [in] LPWSTR    pwzFriendlyName - An optional parameter used to give a friendly name to the domain
		uintptr(unsafe.Pointer(nil)),		// [in] IUnknown* pIdentityArray - An optional array of pointers to IIdentity instances that represent evidence mapped through security policy to establish a permission set
		uintptr(unsafe.Pointer(&_cb204c04)),	// [out] IUnknown** pAppDomain
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		// The specified procedure could not be found.
		// TODO Why is this error message returned?
		_587fa871(fmt.Sprintf("the ICORRuntimeHost::CreateDomain method returned an error:\r\n%s", _a30f3ccf))
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICORRuntimeHost::CreateDomain method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

// EnumDomains Gets an enumerator for the domains in the current process.
// HRESULT EnumDomains (
//
//	[out] HCORENUM *hEnum
//
// );
func (_61de9d6e *_f5ea272a) _b002a8fe() (_22630080 *uintptr, _a30f3ccf error) {
	_587fa871("Enterin into icorruntimehost.EnumDomains()...")

	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._b002a8fe,
		(uintptr(unsafe.Pointer(_22630080))),
		0,
		0,
		0,
	)

	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the ICORRuntimeHost::EnumDomains method returned an error:\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the ICORRuntimeHost::EnumDomains method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	return
}

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

type _3bb2fce1 struct {
	_89dd8d85 *_d8e395dd
}

// MethodInfoVtbl Discovers the attributes of a method and provides access to method metadata.
// Inheritance: Object -> MemberInfo -> MethodBase -> MethodInfo
// MethodInfo Class: https://docs.microsoft.com/en-us/dotnet/api/system.reflection.methodinfo?view=net-5.0
// MethodBase Class: https://docs.microsoft.com/en-us/dotnet/api/system.reflection.methodbase?view=net-5.0
// MemberInfo Class: https://docs.microsoft.com/en-us/dotnet/api/system.reflection.memberinfo?view=net-5.0
// Object Class: https://docs.microsoft.com/en-us/dotnet/api/system.object?view=net-5.0
type _d8e395dd struct {
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
	_de2a9196	uintptr
	_4bf51d8c	uintptr
	_0cf509c6	uintptr
	_03085471	uintptr
	_4f53a52e	uintptr
	_69a89002	uintptr
	_5f89e438	uintptr
	_935051ad	uintptr
	_203305cf	uintptr
	_4341a974	uintptr
	_19c70b7b	uintptr
	_8b192530	uintptr
	_9f275272	uintptr
	_25633211	uintptr
	_af94c45d	uintptr
	_651fbd7b	uintptr
	_57be8f16	uintptr
	_b0069c24	uintptr
	_7b24d19f	uintptr
	_e7b67ec1	uintptr
	_52fd7c0e	uintptr
	_ae2db789	uintptr
	_e7058c1a	uintptr
	_44e1139a	uintptr
	_cf5736e6	uintptr
	_c540f9c3	uintptr
	_33f78d2f	uintptr
	_5568f05b	uintptr
	_fa2b1a1e	uintptr
	_2eaa17f1	uintptr
}

func (_61de9d6e *_3bb2fce1) _09f34622(_8f61fa3d windows.GUID, _60a5cb92 unsafe.Pointer) error {
	_587fa871("Entering into methodinfo.QueryInterface()...")
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

func (_61de9d6e *_3bb2fce1) _aca9fb9e() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._aca9fb9e,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

func (_61de9d6e *_3bb2fce1) _2819d4ac() uintptr {
	_8d822e59, _, _ := syscall.Syscall(
		_61de9d6e._89dd8d85._2819d4ac,
		1,
		uintptr(unsafe.Pointer(_61de9d6e)),
		0,
		0)
	return _8d822e59
}

// Invoke_3 Invokes the method or constructor reflected by this MethodInfo instance.
//
//	virtual HRESULT __stdcall Invoke_3 (
//	/*[in]*/ VARIANT obj,
//	/*[in]*/ SAFEARRAY * parameters,
//	/*[out,retval]*/ VARIANT * pRetVal ) = 0;
//
// https://docs.microsoft.com/en-us/dotnet/api/system.reflection.methodbase.invoke?view=net-5.0
func (_61de9d6e *_3bb2fce1) _33f78d2f(_cccf792c _60f45b99, _261a7772 *_52965c54) (_a30f3ccf error) {
	_587fa871("Entering into methodinfo.Invoke_3()...")
	var _f613cbe1 *_60f45b99
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall6(
		_61de9d6e._89dd8d85._33f78d2f,
		4,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_cccf792c)),
		uintptr(unsafe.Pointer(_261a7772)),
		uintptr(unsafe.Pointer(_f613cbe1)),
		0,
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the MethodInfo::Invoke_3 method returned an error:\r\n%s", _a30f3ccf)
		return
	}

	// If the HRESULT is a TargetInvocationException, attempt to get the inner error
	// This currentl doesn't work
	if uint32(_a0e7c2cc) == _eb3194ae {
		var _e7504228 *_da01f371
		// See if MethodInfo supports the ISupportErrorInfo interface
		_a30f3ccf = _61de9d6e._09f34622(_1359f7cb, unsafe.Pointer(&_e7504228))
		if _a30f3ccf != nil {
			_a30f3ccf = fmt.Errorf("the MethodInfo::QueryInterface method returned an error when looking for the ISupportErrorInfo interface:\r\n%s", _a30f3ccf)
			return
		}

		// See if the ICorRuntimeHost interface supports the IErrorInfo interface
		// Not sure if there is an Interface ID for MethodInfo
		_a30f3ccf = _e7504228._eef19406(_d40bda53)
		if _a30f3ccf != nil {
			_a30f3ccf = fmt.Errorf("there was an error with the ISupportErrorInfo::InterfaceSupportsErrorInfo method:\r\n%s", _a30f3ccf)
			return
		}

		// Get the IErrorInfo object
		_bcc085c2, _7e0c7ead := _6443fdce()
		if _7e0c7ead != nil {
			_a30f3ccf = fmt.Errorf("there was an error getting the IErrorInfo object:\r\n%s", _7e0c7ead)
			return _a30f3ccf
		}

		// Read the IErrorInfo description
		_46428b2c, _f85af1fc := _bcc085c2._46b0dcbd()
		if _f85af1fc != nil {
			_a30f3ccf = fmt.Errorf("the IErrorInfo::GetDescription method returned an error:\r\n%s", _f85af1fc)
			return _a30f3ccf
		}
		if _46428b2c == nil {
			_a30f3ccf = fmt.Errorf("the Assembly::Invoke_3 method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
			return
		}
		_a30f3ccf = fmt.Errorf("the Assembly::Invoke_3 method returned a non-zero HRESULT: 0x%x with an IErrorInfo description of: %s", _a0e7c2cc, *_46428b2c)
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the Assembly::Invoke_3 method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}

	if _f613cbe1 != nil {
		_a30f3ccf = fmt.Errorf("the Assembly::Invoke_3 method returned a non-zero pRetVal: %+v", _f613cbe1)
		return
	}
	_a30f3ccf = nil
	return
}

// GetString returns a string that represents the current object
// a string version of the method's signature
// public virtual string ToString ();
// https://docs.microsoft.com/en-us/dotnet/api/system.object.tostring?view=net-5.0#System_Object_ToString
func (_61de9d6e *_3bb2fce1) GetString() (_142d18df string, _a30f3ccf error) {
	_587fa871("Entering into methodinfo.GetString()...")
	var _99fa303f *string
	_a0e7c2cc, _, _a30f3ccf := syscall.Syscall(
		_61de9d6e._89dd8d85._2a990e1b,
		2,
		uintptr(unsafe.Pointer(_61de9d6e)),
		uintptr(unsafe.Pointer(&_99fa303f)),
		0,
	)
	if _a30f3ccf != syscall.Errno(0) {
		_a30f3ccf = fmt.Errorf("the MethodInfo::ToString method returned an error:\r\n%s", _a30f3ccf)
		return
	}
	if _a0e7c2cc != _56702a9c {
		_a30f3ccf = fmt.Errorf("the Assembly::ToString method returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
		return
	}
	_a30f3ccf = nil
	_142d18df = _6b7f1f9f(unsafe.Pointer(_99fa303f))
	return
}

//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

// SafeArray represents a safe array
// defined in OAIdl.h
//
//	typedef struct tagSAFEARRAY {
//	  USHORT         cDims;
//	  USHORT         fFeatures;
//	  ULONG          cbElements;
//	  ULONG          cLocks;
//	  PVOID          pvData;
//	  SAFEARRAYBOUND rgsabound[1];
//	} SAFEARRAY;
//
// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-safearray
// https://docs.microsoft.com/en-us/archive/msdn-magazine/2017/march/introducing-the-safearray-data-structure
type _52965c54 struct {
	// cDims is the number of dimensions
	_b0a30233	uint16
	// fFeatures is the feature flags
	_05fa6f91	uint16
	// cbElements is the size of an array element
	_e6130b23	uint32
	// cLocks is the number of times the array has been locked without a corresponding unlock
	_b5d34d7c	uint32
	// pvData is the data
	_0956d922	uintptr
	// rgsabout is one bound for each dimension
	_ce250679	[1]_46591b97
}

// SafeArrayBound represents the bounds of one dimension of the array
//
//	typedef struct tagSAFEARRAYBOUND {
//	  ULONG cElements;
//	  LONG  lLbound;
//	} SAFEARRAYBOUND, *LPSAFEARRAYBOUND;
//
// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-safearraybound
type _46591b97 struct {
	// cElements is the number of elements in the dimension
	_9a46542b	uint32
	// lLbound is the lowerbound of the dimension
	_e4325703	int32
}

// CreateSafeArray is a wrapper function that takes in a Go byte array and creates a SafeArray containing unsigned bytes
// by making two syscalls and copying raw memory into the correct spot.
func _f5037928(_6d843506 []byte) (*_52965c54, error) {
	_587fa871("Entering into safearray.CreateSafeArray()...")

	_7f115702 := _46591b97{
		_9a46542b:	uint32(len(_6d843506)),
		_e4325703:	int32(0),
	}

	_8c4015b3, _a30f3ccf := _c7def8d3(_11946586, 1, &_7f115702)
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}
	// now we need to use RtlCopyMemory to copy our bytes to the SafeArray
	_ab69821c := syscall.MustLoadDLL("ntdll.dll")
	_6c5f95f3 := _ab69821c.MustFindProc("RtlCopyMemory")

	// TODO Replace RtlCopyMemory with SafeArrayPutElement or SafeArrayAccessData

	// void RtlCopyMemory(
	//   void*       Destination,
	//   const void* Source,
	//   size_t      Length
	// );
	// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/wdm/nf-wdm-rtlcopymemory
	_, _, _a30f3ccf = _6c5f95f3.Call(
		_8c4015b3._0956d922,
		uintptr(unsafe.Pointer(&_6d843506[0])),
		uintptr(len(_6d843506)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		return nil, _a30f3ccf
	}

	return _8c4015b3, nil
}

// SafeArrayCreate creates a new array descriptor, allocates and initializes the data for the array, and returns a pointer to the new array descriptor.
// SAFEARRAY * SafeArrayCreate(
//
//	VARTYPE        vt,
//	UINT           cDims,
//	SAFEARRAYBOUND *rgsabound
//
// );
// Varient types: https://docs.microsoft.com/en-us/windows/win32/api/wtypes/ne-wtypes-varenum
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraycreate
func _c7def8d3(_ab2c8ac0 uint16, _b0a30233 uint32, _ce250679 *_46591b97) (_8c4015b3 *_52965c54, _a30f3ccf error) {
	_587fa871("Entering into safearray.SafeArrayCreate()...")

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_6c7de62c := _1de575ed.MustFindProc("SafeArrayCreate")

	_8d822e59, _, _a30f3ccf := _6c7de62c.Call(
		uintptr(_ab2c8ac0),
		uintptr(_b0a30233),
		uintptr(unsafe.Pointer(_ce250679)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		return
	}
	_a30f3ccf = nil

	if _8d822e59 == 0 {
		_a30f3ccf = fmt.Errorf("the OleAut32!SafeArrayCreate function return 0x%x and the SafeArray was not created", _8d822e59)
		return
	}

	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	_8c4015b3 = (*_52965c54)(unsafe.Pointer(_8d822e59))
	return
}

// SysAllocString converts a Go string to a BTSR string, that is a unicode string prefixed with its length.
// Allocates a new string and copies the passed string into it.
// It returns a pointer to the string's content.
//
//	BSTR SysAllocString(
//	  const OLECHAR *psz
//	);
//
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-sysallocstring
func _b8bd4b22(_142d18df string) (unsafe.Pointer, error) {
	_587fa871("Entering into safearray.SysAllocString()...")

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_fd2381b2 := _1de575ed.MustFindProc("SysAllocString")

	_d974a256 := _53fa0904(_142d18df)
	_8d822e59, _, _a30f3ccf := _fd2381b2.Call(
		uintptr(unsafe.Pointer(&_d974a256[0])),
	)

	if _a30f3ccf != syscall.Errno(0) {
		return nil, _a30f3ccf
	}
	// TODO Return a pointer to a BSTR instead of an unsafe.Pointer
	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	return unsafe.Pointer(_8d822e59), nil
}

// SafeArrayPutElement pushes an element to the safe array at a given index
//
//	 HRESULT SafeArrayPutElement(
//		  SAFEARRAY *psa,
//		  LONG      *rgIndices,
//		  void      *pv
//	 );
//
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayputelement
func _07e2a1ee(_8622ebec *_52965c54, _65ba55f9 int32, _897cdc64 unsafe.Pointer) error {
	_587fa871("Entering into safearray.SafeArrayPutElement()...")

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_2fe29fb5 := _1de575ed.MustFindProc("SafeArrayPutElement")

	_a0e7c2cc, _, _a30f3ccf := _2fe29fb5.Call(
		uintptr(unsafe.Pointer(_8622ebec)),
		uintptr(unsafe.Pointer(&_65ba55f9)),
		uintptr(_897cdc64),
	)
	if _a30f3ccf != syscall.Errno(0) {
		return _a30f3ccf
	}
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the OleAut32!SafeArrayPutElement call returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return nil
}

// SafeArrayLock increments the lock count of an array, and places a pointer to the array data in pvData of the array descriptor
// HRESULT SafeArrayLock(
//
//	SAFEARRAY *psa
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraylock
func _f520e5d2(_8622ebec *_52965c54) error {
	_587fa871("Entering into safearray.SafeArrayLock()...")

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_b8a1c3a8 := _1de575ed.MustFindProc("SafeArrayCreate")

	_a0e7c2cc, _, _a30f3ccf := _b8a1c3a8.Call(uintptr(unsafe.Pointer(_8622ebec)))

	if _a30f3ccf != syscall.Errno(0) {
		return _a30f3ccf
	}

	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the OleAut32!SafeArrayCreate function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}

	return nil
}

// SafeArrayGetVartype gets the VARTYPE stored in the specified safe array
// HRESULT SafeArrayGetVartype(
//
//	SAFEARRAY *psa,
//	VARTYPE   *pvt
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetvartype
func _68669962(_8622ebec *_52965c54) (uint16, error) {
	_587fa871("Entering into safearray.SafeArrayGetVartype()...")

	var _ab2c8ac0 uint16

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_a17c22ee := _1de575ed.MustFindProc("SafeArrayGetVartype")

	_a0e7c2cc, _, _a30f3ccf := _a17c22ee.Call(
		uintptr(unsafe.Pointer(_8622ebec)),
		uintptr(unsafe.Pointer(&_ab2c8ac0)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		return 0, _a30f3ccf
	}
	if _a0e7c2cc != _56702a9c {
		return 0, fmt.Errorf("the OleAut32!SafeArrayGetVartype function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return _ab2c8ac0, nil
}

// SafeArrayAccessData increments the lock count of an array, and retrieves a pointer to the array data
// HRESULT SafeArrayAccessData(
//
//	SAFEARRAY  *psa,
//	void HUGEP **ppvData
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearrayaccessdata
func _cd58c35d(_8622ebec *_52965c54) (*uintptr, error) {
	_587fa871("Entering into safearray.SafeArrayAccessData()...")

	var _7e87a378 *uintptr

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_5bdc13a0 := _1de575ed.MustFindProc("SafeArrayAccessData")

	_a0e7c2cc, _, _a30f3ccf := _5bdc13a0.Call(
		uintptr(unsafe.Pointer(_8622ebec)),
		uintptr(unsafe.Pointer(&_7e87a378)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		return nil, _a30f3ccf
	}
	if _a0e7c2cc != _56702a9c {
		return nil, fmt.Errorf("the oleaut32!SafeArrayAccessData function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return _7e87a378, nil
}

// SafeArrayGetLBound gets the lower bound for any dimension of the specified safe array
// HRESULT SafeArrayGetLBound(
//
//	SAFEARRAY *psa,
//	UINT      nDim,
//	LONG      *plLbound
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetlbound
func _85689652(_8622ebec *_52965c54, _39c39ae4 uint32) (uint32, error) {
	_587fa871("Entering into safearray.SafeArrayGetLBound()...")
	var _df4f155d uint32
	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_48718959 := _1de575ed.MustFindProc("SafeArrayGetLBound")

	_a0e7c2cc, _, _a30f3ccf := _48718959.Call(
		uintptr(unsafe.Pointer(_8622ebec)),
		uintptr(_39c39ae4),
		uintptr(unsafe.Pointer(&_df4f155d)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		return 0, _a30f3ccf
	}
	if _a0e7c2cc != _56702a9c {
		return 0, fmt.Errorf("the oleaut32!SafeArrayGetLBound function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return _df4f155d, nil
}

// SafeArrayGetUBound gets the upper bound for any dimension of the specified safe array
// HRESULT SafeArrayGetUBound(
//
//	SAFEARRAY *psa,
//	UINT      nDim,
//	LONG      *plUbound
//
// );
// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-safearraygetubound
func _c8b7c069(_8622ebec *_52965c54, _39c39ae4 uint32) (uint32, error) {
	_587fa871("Entering into safearray.SafeArrayGetUBound()...")

	var _924e5f22 uint32

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_d6da7249 := _1de575ed.MustFindProc("SafeArrayGetUBound")

	_a0e7c2cc, _, _a30f3ccf := _d6da7249.Call(
		uintptr(unsafe.Pointer(_8622ebec)),
		uintptr(_39c39ae4),
		uintptr(unsafe.Pointer(&_924e5f22)),
	)

	if _a30f3ccf != syscall.Errno(0) {
		return 0, _a30f3ccf
	}
	if _a0e7c2cc != _56702a9c {
		return 0, fmt.Errorf("the oleaut32!SafeArrayGetUBound function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return _924e5f22, nil
}

// SafeArrayDestroy Destroys an existing array descriptor and all of the data in the array.
// If objects are stored in the array, Release is called on each object in the array.
// HRESULT SafeArrayDestroy(
//
//	SAFEARRAY *psa
//
// );
func _9fa5b547(_8622ebec *_52965c54) error {
	_587fa871("Entering into safearray.SafeArrayDestroy()...")

	_1de575ed := syscall.MustLoadDLL("OleAut32.dll")
	_1e4eafef := _1de575ed.MustFindProc("SafeArrayDestroy")

	_a0e7c2cc, _, _a30f3ccf := _1e4eafef.Call(
		uintptr(unsafe.Pointer(_8622ebec)),
		0,
		0,
	)

	if _a30f3ccf != syscall.Errno(0) {
		return fmt.Errorf("the oleaut32!SafeArrayDestroy function call returned an error:\n%s", _a30f3ccf)
	}
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("the oleaut32!SafeArrayDestroy function returned a non-zero HRESULT: 0x%x", _a0e7c2cc)
	}
	return nil
}

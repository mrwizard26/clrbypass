//go:build windows
// +build windows

package main

const (
	_3185a8d8	uint16	= 0x0000
	// VT_NULL A propagating null value was specified. (This should not be confused with the null pointer.)
	// The null value is used for tri-state logic, as with SQL.
	_fbe2964b	uint16	= 0x0001
	// VT_UI1 is a Variant Type of Unsigned Integer of 1-byte
	_11946586	uint16	= 0x0011
	// VT_UT4 is a Varriant Type of Unsigned Integer of 4-byte
	_c847b0fd	uint16	= 0x0013
	// VT_BSTR is a Variant Type of BSTR, an OLE automation type for transfering length-prefixed strings
	// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-oaut/9c5a5ce4-ff5b-45ce-b915-ada381b34ac1
	_67c57298	uint16	= 0x0008
	// VT_VARIANT is a Variant Type of VARIANT, a container for a union that can hold many types of data
	_ea4d99d1	uint16	= 0x000c
	// VT_ARRAY is a Variant Type of a SAFEARRAY
	// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-oaut/2e87a537-9305-41c6-a88b-b79809b3703a
	_4018bef9	uint16	= 0x2000
)

// from https://github.com/go-ole/go-ole/blob/master/variant_amd64.go
// https://docs.microsoft.com/en-us/windows/win32/winauto/variant-structure
// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-variant
// https://docs.microsoft.com/en-us/previous-versions/windows/embedded/ms891678(v=msdn.10)?redirectedfrom=MSDN
// VARIANT Type Constants https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-oaut/3fe7db9f-5803-4dc4-9d14-5425d3f5461f
type _60f45b99 struct {
	_6fec2b20	uint16	// VARTYPE
	_249e64d4	uint16
	_d11a0a13	uint16
	_aca3ba49	uint16
	_b52f9f3f	uintptr
	_		[8]byte
}

//go:build windows
// +build windows

package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var _2e49904f = false

// checkOK evaluates a HRESULT code for a caller and determines if there was an error
func _c27ead27(_a0e7c2cc uintptr, _648633e6 string) error {
	if _a0e7c2cc != _56702a9c {
		return fmt.Errorf("%s returned 0x%08x", _648633e6, _a0e7c2cc)
	} else {
		return nil
	}
}

func _53fa0904(_bc66dc16 string) []byte {
	_86158d7d := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	var _41b9ddb8 bytes.Buffer
	_2b2ec939 := transform.NewWriter(&_41b9ddb8, _86158d7d)
	_2b2ec939.Write([]byte(_bc66dc16))
	return _41b9ddb8.Bytes()
}

func _176f2599(_d974a256 string) bool {
	return !strings.Contains(_d974a256, "Void Main()")
}

// ReadUnicodeStr takes a pointer to a unicode string in memory and returns a string value
func _6b7f1f9f(_8cfca4d7 unsafe.Pointer) string {
	_587fa871("Entering into utils.ReadUnicodeStr()...")
	var _c48bcd78 uint16
	_b74a5d70 := make([]uint16, 0)
	for _0dd0a21d := 0; ; _0dd0a21d++ {
		_c48bcd78 = *(*uint16)(unsafe.Pointer(_8cfca4d7))
		if _c48bcd78 == 0x0000 {
			break
		}
		_b74a5d70 = append(_b74a5d70, _c48bcd78)
		_8cfca4d7 = unsafe.Pointer(uintptr(_8cfca4d7) + 2)
	}
	return string(utf16.Decode(_b74a5d70))
}

// debugPrint is used to print messages only when debug has been enabled
func _587fa871(_59a58cf1 string) {
	if _2e49904f {
		fmt.Println("[DEBUG] " + _59a58cf1)
	}
}

// PrepareParameters creates a safe array of strings (arguments) nested inside a Variant object, which is itself
// appended to the final safe array
func _b130eb4d(_11666973 []string) (*_52965c54, error) {
	_894e0d17 := _46591b97{
		_9a46542b:	uint32(len(_11666973)),
		_e4325703:	0,
	}
	_9df5a5a0, _a30f3ccf := _c7def8d3(_67c57298, 1, &_894e0d17)	// VT_BSTR
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}
	for _0dd0a21d, _cad88847 := range _11666973 {
		_0c68908b, _a30f3ccf := _b8bd4b22(_cad88847)
		if _a30f3ccf != nil {
			return nil, _a30f3ccf
		}
		_07e2a1ee(_9df5a5a0, int32(_0dd0a21d), _0c68908b)
	}

	_88cb0ec1 := _60f45b99{
		_6fec2b20:	_67c57298 | _4018bef9,	// VT_BSTR | VT_ARRAY
		_b52f9f3f:	uintptr(unsafe.Pointer(_9df5a5a0)),
	}

	_fe39fe0b := _46591b97{
		_9a46542b:	uint32(1),
		_e4325703:	0,
	}
	_63b36428, _a30f3ccf := _c7def8d3(_ea4d99d1, 1, &_fe39fe0b)	// VT_VARIANT
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}
	_a30f3ccf = _07e2a1ee(_63b36428, int32(0), unsafe.Pointer(&_88cb0ec1))
	if _a30f3ccf != nil {
		return nil, _a30f3ccf
	}
	return _63b36428, nil
}

func _58ef3fa7(_cad88847 uintptr) {
	_ede583ea := windows.NewLazyDLL("kernel32.dll")
	_a757bd55 := _ede583ea.NewProc("GlobalFree")
	_a757bd55.Call(_cad88847)
}

func _1e3d9ece(_cad88847 uintptr) {
	_ede583ea := windows.NewLazyDLL("kernel32.dll")
	_a757bd55 := _ede583ea.NewProc("GlobalFree")
	_a757bd55.Call(_cad88847)
}

func _84d30ed8(_5392e397 uint32) uint32 {
	if _5392e397 <= 0 {
		return _5392e397
	}
	return (_5392e397 & 0x0000FFFF) | (7 << 16) | 0x80000000
}

func _61d36a31(_0abbc90a *byte, _595becea uint32) *_d9a435f2 {
	_ede583ea := windows.NewLazyDLL("shlwapi.dll")
	_f4688ff9 := _ede583ea.NewProc("SHCreateMemStream")
	_18a7b7b5, _, _a30f3ccf := _f4688ff9.Call(uintptr(unsafe.Pointer(_0abbc90a)), uintptr(_595becea))
	if _18a7b7b5 == 0 {
		fmt.Println("Error in SHCreateMemStream ", _18a7b7b5)
		fmt.Println(_a30f3ccf)
	}
	return (*_d9a435f2)(unsafe.Pointer(_18a7b7b5))
}

func _f721ad6b(_ec4730c1 uintptr) uintptr {
	_ede583ea := windows.NewLazyDLL("kernel32.dll")
	_a757bd55 := _ede583ea.NewProc("HeapCreate")
	_b55f5491, _, _ := _a757bd55.Call(_ec4730c1, 0, 0)
	return _b55f5491
}

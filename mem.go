package main

import (
	"unsafe"
)

type _165be515 = unsafe.Pointer
type _92fc56bd = int

func _9d1d58a0(_0dcc3140, _d8c54c0e unsafe.Pointer, len _92fc56bd) unsafe.Pointer {

	_3c429551 := len >> 3
	var _0dd0a21d _92fc56bd = 0
	for _0dd0a21d = 0; _0dd0a21d < _3c429551; _0dd0a21d++ {
		var _d066d27b *uint64 = (*uint64)(_165be515(uintptr(_0dcc3140) + uintptr(8*_0dd0a21d)))
		var _c016274d *uint64 = (*uint64)(_165be515(uintptr(_d8c54c0e) + uintptr(8*_0dd0a21d)))
		*_d066d27b = *_c016274d
	}
	_be3b4d44 := len & 7
	for _0dd0a21d = 0; _0dd0a21d < _be3b4d44; _0dd0a21d++ {
		var _d066d27b *uint8 = (*uint8)(_165be515(uintptr(_0dcc3140) + uintptr(8*_3c429551+_0dd0a21d)))
		var _c016274d *uint8 = (*uint8)(_165be515(uintptr(_d8c54c0e) + uintptr(8*_3c429551+_0dd0a21d)))

		*_d066d27b = *_c016274d
	}
	return _0dcc3140
}

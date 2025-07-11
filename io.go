//go:build windows
// +build windows

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

// origSTDOUT is a Windows Handle to the program's original STDOUT
var _a9b38311 = windows.Stdout

// origSTDERR is a Windows Handle to the program's original STDERR
var _afb267d1 = windows.Stderr

// rSTDOUT is an io.Reader for STDOUT
var _37ead9b4 *os.File

// wSTDOUT is an io.Writer for STDOUT
var _5d394cec *os.File

// rSTDERR is an io.Reader for STDERR
var _74b40d30 *os.File

// wSTDERR is an io.Writer for STDERR
var _e3f0d8c6 *os.File

// Stdout is a buffer to collect anything written to STDOUT
// The CLR will return the COR_E_TARGETINVOCATION error on subsequent Invoke_3 calls if the
// redirected STDOUT writer is EVER closed while the parent process is running (e.g., a C2 Agent)
// The redirected STDOUT reader will never recieve EOF and therefore reads will block and that is
// why a buffer is used to stored anything that has been written to STDOUT while subsequent calls block
var _baacd201 bytes.Buffer

// Stderr is a buffer to collect anything written to STDERR
var _fb936b21 bytes.Buffer

// errors is used to capture an errors from a goroutine
var _d25193f1 = make(chan error)

// mutex ensures exclusive access to read/write on STDOUT/STDERR by one routine at a time
var _5ad84997 = &sync.Mutex{}

// RedirectStdoutStderr redirects the program's STDOUT/STDERR to an *os.File that can be read from this Go program
// The CLR executes assemblies outside of Go and therefore STDOUT/STDERR can't be captured using normal functions
// Intended to be used with a Command & Control framework so STDOUT/STDERR can be captured and returned
func _ef0e36d8() (_a30f3ccf error) {
	// Create a new reader and writer for STDOUT
	_37ead9b4, _5d394cec, _a30f3ccf = os.Pipe()
	if _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error calling the os.Pipe() function to create a new STDOUT:\n%s", _a30f3ccf)
		return
	}

	// Create a new reader and writer for STDERR
	_74b40d30, _e3f0d8c6, _a30f3ccf = os.Pipe()
	if _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error calling the os.Pipe() function to create a new STDERR:\n%s", _a30f3ccf)
		return
	}

	_6acf5dbf := windows.NewLazySystemDLL("kernel32.dll")
	_b535b6da := _6acf5dbf.NewProc("GetConsoleWindow")

	// Ensure the process has a console because if it doesn't there will be no output to capture
	_0ee1a68d, _, _ := _b535b6da.Call()
	if _0ee1a68d == 0 {
		// https://learn.microsoft.com/en-us/windows/console/allocconsole
		_f19bed8f := _6acf5dbf.NewProc("AllocConsole")
		// BOOL WINAPI AllocConsole(void);
		_8d822e59, _, _a30f3ccf := _f19bed8f.Call()
		// A process can be associated with only one console, so the AllocConsole function fails if the calling process
		// already has a console. So long as any console exists we are good to go and therefore don't care about errors
		if _8d822e59 == 0 {
			return fmt.Errorf("there was an error calling kernel32!AllocConsole with return code %d: %s", _8d822e59, _a30f3ccf)
		}

		// Get a handle to the newly created/allocated console
		_0ee1a68d, _, _ = _b535b6da.Call()

		_8a526f19 := windows.NewLazySystemDLL("user32.dll")
		_faf4f368 := _8a526f19.NewProc("ShowWindow")
		// Hide the console window
		_8d822e59, _, _a30f3ccf = _faf4f368.Call(_0ee1a68d, windows.SW_HIDE)
		if _a30f3ccf != syscall.Errno(0) {
			return fmt.Errorf("there was an error calling user32!ShowWindow with return %+v: %s", _8d822e59, _a30f3ccf)
		}
	}

	// Set STDOUT/STDERR to the new files from os.Pipe()
	// https://docs.microsoft.com/en-us/windows/console/setstdhandle
	if _a30f3ccf = windows.SetStdHandle(windows.STD_OUTPUT_HANDLE, windows.Handle(_5d394cec.Fd())); _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error calling the windows.SetStdHandle function for STDOUT:\n%s", _a30f3ccf)
		return
	}

	if _a30f3ccf = windows.SetStdHandle(windows.STD_ERROR_HANDLE, windows.Handle(_e3f0d8c6.Fd())); _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error calling the windows.SetStdHandle function for STDERR:\n%s", _a30f3ccf)
		return
	}

	// Start STDOUT/STDERR buffer and collection
	go _052467e6()
	go _8481d299()

	return
}

// RestoreStdoutStderr returns the program's original STDOUT/STDERR handles before they were redirected an *os.File
// Previously instantiated CLRs will continue to use the REDIRECTED STDOUT/STDERR handles and will not resume
// using the restored handles
func _746dccff() error {
	if _a30f3ccf := windows.SetStdHandle(windows.STD_OUTPUT_HANDLE, _a9b38311); _a30f3ccf != nil {
		return fmt.Errorf("there was an error calling the windows.SetStdHandle function to restore the original STDOUT handle:\n%s", _a30f3ccf)
	}
	if _a30f3ccf := windows.SetStdHandle(windows.STD_ERROR_HANDLE, _afb267d1); _a30f3ccf != nil {
		return fmt.Errorf("there was an error calling the windows.SetStdHandle function to restore the original STDERR handle:\n%s", _a30f3ccf)
	}
	return nil
}

// ReadStdoutStderr reads from the REDIRECTED STDOUT/STDERR
// Only use when RedirectStdoutStderr was previously called
func _8726ee7b() (_6e4a0eae string, _d4783fa6 string, _a30f3ccf error) {
	_587fa871("Entering into io.ReadStdoutStderr()...")

	// Sleep for one Microsecond to wait for STDOUT/STDERR goroutines to finish reading
	// Race condition between reading the buffers and reading STDOUT/STDERR to the buffers
	// Can't close STDOUT/STDERR writers once the CLR invokes on assembly and EOF is not
	// returned because parent program is perpetually running
	time.Sleep(1 * time.Microsecond)

	// Check the error channel to see if any of the goroutines generated an error
	if len(_d25193f1) > 0 {
		var _03b2d953 string
		for _0787a2a8 := range _d25193f1 {
			_03b2d953 += _0787a2a8.Error()
		}
		_a30f3ccf = fmt.Errorf(_03b2d953)
		return
	}

	// Read STDOUT Buffer
	if _baacd201.Len() > 0 {
		_6e4a0eae = _baacd201.String()
		_baacd201.Reset()
	}

	// Read STDERR Buffer
	if _fb936b21.Len() > 0 {
		_d4783fa6 = _fb936b21.String()
		_fb936b21.Reset()
	}
	return
}

// CloseStdoutStderr closes the Reader/Writer for the previously redirected STDOUT/STDERR
// that was changed to an *os.File
func _e9a81e22() (_a30f3ccf error) {
	_a30f3ccf = _37ead9b4.Close()
	if _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error closing the STDOUT Reader:\n%s", _a30f3ccf)
		return
	}

	_a30f3ccf = _5d394cec.Close()
	if _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error closing the STDOUT Writer:\n%s", _a30f3ccf)
		return
	}

	_a30f3ccf = _74b40d30.Close()
	if _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error closing the STDERR Reader:\n%s", _a30f3ccf)
		return
	}

	_a30f3ccf = _e3f0d8c6.Close()
	if _a30f3ccf != nil {
		_a30f3ccf = fmt.Errorf("there was an error closing the STDERR Writer:\n%s", _a30f3ccf)
		return
	}
	return nil
}

// BufferStdout is designed to be used as a go routine to monitor for data written to the REDIRECTED STDOUT
// and collect it into a buffer so that it can be collected and sent back to a server
func _052467e6() {
	_587fa871("Entering into io.BufferStdout()...")
	_a9118195 := bufio.NewReader(_37ead9b4)
	for {
		// Standard STDOUT buffer size is 4k
		_41b9ddb8 := make([]byte, 4096)
		_83759fa3, _a30f3ccf := _a9118195.Read(_41b9ddb8)
		if _a30f3ccf != nil {
			_d25193f1 <- fmt.Errorf("there was an error reading from STDOUT in io.BufferStdout:\n%s", _a30f3ccf)
		}
		if _83759fa3 > 0 {
			// Remove null bytes and add contents to the buffer
			_baacd201.Write(bytes.TrimRight(_41b9ddb8, "\x00"))
		}
	}
}

// BufferStderr is designed to be used as a go routine to monitor for data written to the REDIRECTED STDERR
// and collect it into a buffer so that it can be collected and sent back to a server
func _8481d299() {
	_587fa871("Entering into io.BufferStderr()...")
	_a90c73e8 := bufio.NewReader(_74b40d30)
	for {
		// Standard STDOUT buffer size is 4k
		_41b9ddb8 := make([]byte, 4096)
		_83759fa3, _a30f3ccf := _a90c73e8.Read(_41b9ddb8)
		if _a30f3ccf != nil {
			_d25193f1 <- fmt.Errorf("there was an error reading from STDOUT in io.BufferStdout:\n%s", _a30f3ccf)
		}
		if _83759fa3 > 0 {
			// Remove null bytes and add contents to the buffer
			_fb936b21.Write(bytes.TrimRight(_41b9ddb8, "\x00"))
		}
	}
}

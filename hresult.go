package main

// https://docs.microsoft.com/en-us/dotnet/framework/interop/how-to-map-hresults-and-exceptions
// https://docs.microsoft.com/en-us/windows/win32/seccrypto/common-hresult-values

const (
	_56702a9c	= 0x00
	_4ef33ea8	= 0x01
	// COR_E_TARGETINVOCATION is TargetInvocationException
	// https://docs.microsoft.com/en-us/dotnet/api/system.reflection.targetinvocationexception?view=net-5.0
	_eb3194ae	uint32	= 0x80131604
	// COR_E_SAFEARRAYRANKMISMATCH is SafeArrayRankMismatchException
	_72a60f9f	uint32	= 0x80131538
	// COR_E_BADIMAGEFORMAT is BadImageFormatException
	_f9060f41	uint32	= 0x8007000b
	// DISP_E_BADPARAMCOUNT is invalid number of parameters
	_06cd7ebf	uint32	= 0x8002000e
	// E_POINTER Pointer that is not valid
	_92386235	uint32	= 0x80004003
	// E_NOINTERFACE No such interface supported
	_482e4fcc	uint32	= 0x80004002
	// Not Implemented
	_7903a6c8	uint32	= 0x80004001
)

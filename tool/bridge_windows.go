package main

// #include <windows.h>
import "C"
import "unsafe"

func loadLibraryFoo(path string) unsafe.Pointer {
	return C.LoadLibraryA(C.CString(path))
}

func findSymbolFoo(library unsafe.Pointer, symbol string) unsafe.Pointer {
	return C.GetProcAddress(library, C.CString(symbol))
}

func closeLibraryFoo(library unsafe.Pointer) {
	C.FreeLibrary(library)
}

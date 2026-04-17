package main

// #include <windows.h>
import "C"
import "unsafe"

func loadLibrary(path string) unsafe.Pointer {
	return C.LoadLibraryA(C.CString(path))
}

func findSymbol(library unsafe.Pointer, symbol string) unsafe.Pointer {
	return C.GetProcAddress(library, C.CString(symbol))
}

func closeLibrary(library unsafe.Pointer) {
	C.FreeLibrary(library)
}

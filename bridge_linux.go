package main

// #include <dlfcn.h>
import "C"
import "unsafe"

func loadLibrary(path string) unsafe.Pointer {
	return C.dlopen(C.CString(path), C.RTLD_LAZY)
}

func findSymbol(library unsafe.Pointer, symbol string) unsafe.Pointer {
	return C.dlsym(library, C.CString(symbol))
}

func closeLibrary(library unsafe.Pointer) {
	C.dlclose(library)
}

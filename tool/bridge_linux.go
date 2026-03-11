package main

// #include <dlfcn.h>
import "C"
import "unsafe"

func loadLibraryFoo(path string) unsafe.Pointer {
	return C.dlopen(C.CString(path), C.RTLD_LAZY)
}

func findSymbolFoo(library unsafe.Pointer, symbol string) unsafe.Pointer {
	return C.dlsym(library, C.CString(symbol))
}

func closeLibraryFoo(library unsafe.Pointer) {
	C.dlclose(library)
}

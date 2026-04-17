package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s %s\n", os.Args[1], os.Args[2])
	var library = loadLibrary(os.Args[1])
	if library == nil {
		fmt.Printf("library not found: %s\n", os.Args[1])
		return
	}
	defer closeLibrary(library)
	fmt.Printf("%v\n", library)
	var symbol = findSymbol(library, os.Args[2])
	fmt.Printf("%v\n", symbol)
}

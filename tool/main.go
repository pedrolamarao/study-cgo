package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "tool",
		Short: "A tool to study cgo",
	}

	var testCmd = &cobra.Command{
		Use:   "test [arg]",
		Short: "A command to test functionality",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s %s\n", args[0], args[1])
			var library = loadLibraryFoo(args[0])
			if library == nil {
				fmt.Printf("library not found: %s\n", args[0])
				return
			}
			defer closeLibraryFoo(library)
			fmt.Printf("%v\n", library)
			var symbol = findSymbolFoo(library, args[1])
			fmt.Printf("%v\n", symbol)
		},
	}

	rootCmd.AddCommand(testCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

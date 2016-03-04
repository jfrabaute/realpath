package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "realpath: need at least one filename")
		os.Exit(1)
	}
	for _, f := range os.Args[1:] {
		rp, err := filepath.Abs(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", f, err)
			continue
		}
		rp, err = filepath.EvalSymlinks(rp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", f, err)
			continue
		}
		fmt.Println(rp)
	}
}

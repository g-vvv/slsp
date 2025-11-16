package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/g-vvv/slsp"
)

func main() {
	searchStr := flag.String("s", "", "The string to be patched (required).")
	replaceStr := flag.String("r", "", "The patch string (required).")
	limit := flag.Int("l", -1, "Optional: Limit patching to the first L bytes. (default: full file length)")
	occurrences := flag.Int("n", -1, "Optional: Max number of patches. (default: all occurrences)")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Error: Exactly one filename argument is required.")
		fmt.Println("Usage: sslr -s <search> -r <replace> [-l <bytes>] [-n <count>] <filename>")
		fmt.Println("Example: sslr -s 'old' -r 'new' -l 1024 -n 5 config.txt")
		flag.PrintDefaults()
		os.Exit(1)
	}
	filename := args[0]

	if *searchStr == "" || *replaceStr == "" {
		log.Println("Error: -s (search) and -r (replace/patch) flags are required.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Enforce the Same-Length Constraint
	if len(*searchStr) != len(*replaceStr) {
		log.Fatalln("Error: -s (search) and -r (replace/patch) strings must be the same length.")
	}
	if len(*searchStr) == 0 {
		log.Fatalln("Error: Strings cannot be empty.")
	}

	slsp.Patch(filename, *searchStr, *replaceStr, *limit, *occurrences)
}

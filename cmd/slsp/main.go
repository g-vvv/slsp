package main

import (
	"flag"
	"fmt"
	"os"

	slsp "github.com/g-vvv/slsp"
)

func main() {
	searchStr := flag.String("s", "", "The string to be patched (required).")
	replaceStr := flag.String("r", "", "The patch string (required).")
	limit := flag.Int("l", -1, "Optional: Limit patching to the first L bytes. (default: full file length)")
	occurrences := flag.Int("n", -1, "Optional: Max number of patches. (default: all occurrences)")
	help := flag.Bool("h", false, "Print this help message.")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Error: Exactly one filename argument is required.")
		*help = true
	} else if *searchStr == "" || *replaceStr == "" {
		fmt.Println("Error: -s (search) and -r (replace/patch) flags are required.")
		*help = true
	}
	if *help {
		fmt.Println("Usage: slsp -s <search> -r <replace> [-l <bytes>] [-n <count>] <filename>")
		fmt.Println("Example: slsp -s 'old' -r 'new' -l 1024 -n 5 config.txt")
		flag.PrintDefaults()
		os.Exit(1)
	}
	filename := args[0]

	result, err := slsp.Patch(filename, *searchStr, *replaceStr, *limit, *occurrences)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

package sslr

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Patch(filename string, searchStr string, replaceStr string, limit int, occurrences int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file %s: %v\n", filename, err)
	}
	fileLength := len(content)

	effectiveLimit := fileLength
	if limit != -1 { // -1 is our default, so if it's *not* -1, the user set it.
		if limit < 0 {
			log.Fatalf("Error: -l (limit) cannot be negative.\n")
		}
		if limit > fileLength {
			log.Fatalf("Error: -l (%d) cannot be greater than file length (%d).\n", limit, fileLength)
		}
		effectiveLimit = limit // Use the user-provided limit
	}

	n := occurrences

	oldBytes := []byte(searchStr)
	newBytes := []byte(replaceStr)

	contentToSearch := content[0:effectiveLimit] // The part we will modify
	contentToKeep := content[effectiveLimit:]    // The part we will not touch

	if !bytes.Contains(contentToSearch, oldBytes) {
		fmt.Printf("No occurrences of '%s' found in the first %d bytes of %s. File is unchanged.\n", searchStr, effectiveLimit, filename)
		os.Exit(0)
	}

	// Perform the replacement, passing the 'n' value directly.
	modifiedPart := bytes.Replace(contentToSearch, oldBytes, newBytes, n)

	// Combine the modified first part with the untouched second part.
	modifiedContent := append(modifiedPart, contentToKeep...)

	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Error getting file stats %s: %v\n", filename, err)
	}

	err = os.WriteFile(filename, modifiedContent, fileInfo.Mode())
	if err != nil {
		log.Fatalf("Error writing file %s: %v\n", filename, err)
	}

	nStr := "all"
	if n != -1 {
		nStr = fmt.Sprintf("%d", n)
	}

	fmt.Printf("Successfully patched up to %s instances of '%s' with '%s' in the first %d bytes of %s.\n",
		nStr, searchStr, replaceStr, effectiveLimit, filename)
}

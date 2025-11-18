package slsp

import (
	"bytes"
	"fmt"
	"os"
)

func Patch(
	filename string,
	searchStr string,
	replaceStr string,
	limit int,
	occurrences int,
) (
	result string,
	err error,
) {
	// Enforce the Same-Length Constraint
	if len(searchStr) != len(replaceStr) {
		err = fmt.Errorf("-s (search) and -r (replace/patch) strings must be the same length")
		return
	}
	if len(searchStr) == 0 {
		err = fmt.Errorf("strings cannot be empty")
		return
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("while reading file %s: %w", filename, err)
		return
	}
	fileLength := len(content)

	effectiveLimit := fileLength
	if limit != -1 { // -1 is our default, so if it's *not* -1, the user set it.
		if limit < 0 {
			err = fmt.Errorf("-l (limit) cannot be negative")
			return
		}
		if limit > fileLength {
			err = fmt.Errorf("-l (%d) cannot be greater than file length (%d)", limit, fileLength)
			return
		}
		effectiveLimit = limit // Use the user-provided limit
	}

	n := occurrences

	oldBytes := []byte(searchStr)
	newBytes := []byte(replaceStr)

	contentToSearch := content[0:effectiveLimit] // The part we will modify
	contentToKeep := content[effectiveLimit:]    // The part we will not touch

	if !bytes.Contains(contentToSearch, oldBytes) {
		result = fmt.Sprintf(
			"No occurrences of %q found in the first %d bytes of %s. File is unchanged.",
			searchStr, effectiveLimit, filename)
		return
	}

	// Perform the replacement, passing the 'n' value directly.
	modifiedPart := bytes.Replace(contentToSearch, oldBytes, newBytes, n)

	// Combine the modified first part with the untouched second part.
	modifiedContent := append(modifiedPart, contentToKeep...)

	fileInfo, err := os.Stat(filename)
	if err != nil {
		result = fmt.Sprintf("while getting file stats %s: %v", filename, err)
		return
	}

	err = os.WriteFile(filename, modifiedContent, fileInfo.Mode())
	if err != nil {
		err = fmt.Errorf("while writing file %s: %w", filename, err)
		return
	}

	nStr := "all"
	if n != -1 {
		nStr = fmt.Sprintf("%d", n)
	}

	result = fmt.Sprintf(
		"Successfully patched up to %s instances of '%s' with '%s' in the first %d bytes of %s.",
		nStr, searchStr, replaceStr, effectiveLimit, filename)

	return
}

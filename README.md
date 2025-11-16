## Installation
Requires Go 1.23+

```
$ go install github.com/g-vvv/slsp/cmd/slsp@latest
```
## Usage

```
Usage: slsp -s <search> -r <replace> [-l <bytes>] [-n <count>] <filename>
Example: slsp -s 'old' -r 'new' -l 1024 -n 5 config.txt
  -l int
        Optional: Limit patching to the first L bytes. (default: full file length) (default -1)
  -n int
        Optional: Max number of patches. (default: all occurrences) (default -1)
  -r string
        The patch string (required).
  -s string
        The string to be patched (required).
```

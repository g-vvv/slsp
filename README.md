# 🔧 slsp (Offset-Safe String Patcher)

`slsp` is a lightweight, dependency-free command-line utility for quick and precise string replacement in files.

Its primary feature is that it's **offset-safe**: it enforces that the replacement string is the **exact same byte length** as the search string. This preserves all file offsets, making it 100% safe for patching binaries, executables, and other offset-sensitive files where changing the file size would cause corruption.

-----

## ✨ Key Features

  * **🔒 Offset Preservation:** `slsp` guarantees file integrity by requiring the search and replace strings to be the exact same length.
  * 🧬 **Binary Safe:** Works directly on bytes, making it safe for executables and text files alike.
  * ⚙️ **Scope Control:** Use the `-l` flag to limit patching to only the first $N$ bytes of a file (perfect for header manipulation).
  * 🔢 **Occurrence Limiting:** Use the `-n` flag to control the *maximum* number of replacements performed.

-----

## 🚀 Installation

Requires **Go 1.23+**

```bash
go install github.com/g-vvv/slsp/cmd/slsp@latest
```

-----

## ⌨️ Usage

### Syntax

```bash
slsp -s <search> -r <replace> [-l <bytes>] [-n <count>] <filename>
```

> **Important:** The byte length of the string for `-s` *must* equal the byte length of the string for `-r`. The tool will exit with an error if they do not match.

### Flags

| Flag | Description | Default |
| :--- | :--- | :--- |
| **`-s`** | **(Required)** The string to search for. | N/A |
| **`-r`** | **(Required)** The string to replace with. | N/A |
| `-l` | Optional: Limit patching to the first $L$ bytes. | `-1` (Full file) |
| `-n` | Optional: Max number of patches to perform. | `-1` (All occurrences) |

-----

## 💡 Examples

### 1\. Basic Replacement (Length-Matched)

Replace all instances of `ver_01` with `ver_02` in `config.ini`.

```bash
# "ver_01" and "ver_02" are both 6 bytes
slsp -s "ver_01" -r "ver_02" config.ini
```

### 2\. Patching a Binary Header

Replace the *first* occurrence of `BAD_MAGIC` with `GOOD_MAGIC` within the *first 1024 bytes* of an executable.

```bash
# "BAD_MAGIC" and "GOOD_MAGIC" are both 9 bytes
slsp -s "BAD_MAGIC" -r "GOOD_MAGIC" -l 1024 -n 1 program.exe
```

### 3\. Limiting Total Replacements

Update only the *first 5 instances* of `DEBUG=1` to `DEBUG=0` in a large log file.

```bash
slsp -s "DEBUG=1" -r "DEBUG=0" -n 5 massive.log
```

### 4\. Invalid (Failed) Example

The following command will fail because the string lengths do not match:

```bash
# "Error" (5 bytes) vs "Warn" (4 bytes)
slsp -s "Error" -r "Warn" app.log

# Error: -s (search) and -r (replace/patch) strings must be the same length.
```

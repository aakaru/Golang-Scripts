# gowc-tool

A lightweight Go implementation of the Unix/Linux `wc` (word count) utility that counts lines, words, and bytes from standard input.

## Overview

`gowc-tool` is a simple, efficient text analysis tool written in Go that provides basic text statistics similar to the classic Unix `wc` command. It reads from standard input and outputs counts for:

- Number of lines
- Number of words
- Number of bytes

## Installation

### Prerequisites

- Go (version 1.11 or higher recommended)

### Building from Source

1. Clone the repository or download the source code:
   ```
   git clone https://github.com/yourusername/gowc-tool.git
   cd gowc-tool
   ```

2. Build the executable:
   ```
   go build mywc.go
   ```

This will create an executable named `mywc` in your current directory.

## Usage

### Basic Usage

The tool reads from standard input and outputs the count of lines, words, and bytes:

```
./mywc
```

After running this command, type or paste your text and press `Ctrl+D` (Unix/Linux/macOS) or `Ctrl+Z` followed by `Enter` (Windows) to signal the end of input.

### Pipe Input

You can pipe text from another command:

```
echo "This is a test." | ./mywc
```

Output:
```
1 lines 4 words 15 bytes
```

### Process a File

Process the contents of a file:

```
cat myfile.txt | ./mywc
```

Or using input redirection:

```
./mywc < myfile.txt
```


## Implementation Details

The tool uses Go's `bufio.Scanner` to efficiently read input line by line, counting:
- Lines: Incremented for each line read
- Words: Determined by splitting each line into fields using whitespace as a delimiter
- Bytes: Calculated based on line length + 1 (for the newline character)

## License

[MIT LICENSE]

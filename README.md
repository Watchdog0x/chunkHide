# chunkHide
![Static Badge](https://img.shields.io/badge/Version-1.0.0-brightgreen?style=for-the-badge&labelColor=%23161B22&color=rgb(93%2C%2063%2C%20211))

`chunkHide` is a command-line utility for modifying PNG image files by manipulating chunks. 
It allows you to add, modify, or read text chunks in a PNG image.

## Table of Contents
- [Options](#options)
- [Installation](#installation)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Options
- **-i [image.png]:** Path to the PNG image file.
- **-t [tEXt]:** Specify chunk types:
    - tEXt: Text chunk
    - zTXt: Compressed text chunk
    - PLTE: Palette chunk
- **-o [output.png]:** Where to output the modified image. Default: output.png
- **-v [image.png]:** Validate the image chunks. Example: chunkHide -v image.png
- **-r:** Read the chunk data. Requires the -t option. Example: chunkHide -r -t tEXt
- **-keyword:** Keyword for the chunk. Required with -t tEXt or -t zTXt.
- **-text:** Text for the chunk. Required with -t tEXt, -t zTXt, or -t PLTE.

## Installation

To use `chunkHide` in your Go projects or as a command-line tool, follow these installation steps:

### Downloading the Latest Release

Visit the [Latest Releases](https://github.com/Watchdog0x/chunkHide/releases) page on the GitHub repository to download the binary executable for your operating system. Pre-built binaries for common platforms, such as Linux, macOS, and Windows, are available.

### Manual Installation

Alternatively, install `chunkHide` manually by cloning the GitHub repository and building the binary using Go:

```bash
git clone https://github.com/Watchdog0x/chunkHide.git
cd chunkHide
go build -o chunkHide cmd/main.go
```

### Installing as a Go Module

If you want to integrate chunkHide as a dependency in your Go project, you can add it as a module:

```bash
go get github.com/Watchdog0x/chunkHide
```

Import it in your Go code:

```go
import "github.com/Watchdog0x/chunkHide"
```

Then, you can use chunkHide as a module in your project.


## Examples

```bash
# Add a text chunk to the image
chunkHide -t tEXt -keyword mykey -text Hello -i input.png -o output.png

# Add a compressed text chunk to the image
chunkHide -t zTXt -keyword mykey -text Hello -i input.png -o output.png

# Read text data from the image
chunkHide -r -t tEXt -i input.png

# Validate image chunks
chunkHide -v input.png
```


## Contributing
Contributions are welcome! Feel free to open issues, submit pull requests, or provide suggestions. Please follow the Contributing Guidelines.

## License
This project is licensed under the MIT License



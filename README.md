# linux_coreutils

A portable, Go-based reimplementation of common Linux/Unix command-line utilities.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.24-blue.svg)](https://golang.org)
[![Release](https://img.shields.io/github/v/release/user/linux_coreutils)](https://github.com/user/linux_coreutils/releases)

## Overview

`linux_coreutils` provides self-contained, single-binary implementations of 44 frequently used Linux/Unix commands. The project focuses on the POSIX/GNU Coreutils subset used in daily development, operations, and scripting, with first-class support for Linux, Windows, and macOS via Go's native cross-compilation.

## Features

- **44 commands** covering file operations, text processing, system information, networking, compression, and utilities.
- **Cross-platform**: builds natively for Linux, Windows, and macOS (`amd64` / `arm64`).
- **Zero or minimal external dependencies**: simple commands use only the Go standard library.
- **Single-binary per command**: each command compiles to an independent executable for easy testing and selective deployment.
- **Built-in help**: every command supports `-h` / `--help`.
- **Tested**: unit tests for every command plus shell-based integration tests.

## Command List

### File & Directory Operations

| Command | Description |
|---------|-------------|
| `ls`    | List directory contents (`-l`, `-a`, `-h`, `-R`) |
| `cat`   | Concatenate and print files (`-n`, `-b`) |
| `cp`    | Copy files/directories (`-r`, `-v`, `-i`) |
| `mv`    | Move/rename files (`-i`, `-v`) |
| `rm`    | Remove files/directories (`-r`, `-f`, `-i`) |
| `mkdir` | Create directories (`-p`, `-m`) |
| `rmdir` | Remove empty directories (`-p`) |
| `touch` | Create empty files or update timestamps |
| `find`  | Find files in a directory tree (`-name`, `-type`, `-size`) |
| `which` | Locate executables in `PATH` |
| `pwd`   | Print working directory |
| `realpath` | Print canonical absolute path |
| `basename` / `dirname` | Extract path components |

### Text Processing

| Command | Description |
|---------|-------------|
| `grep`  | Search with regex (`-i`, `-v`, `-n`, `-r`) |
| `cut`   | Cut by bytes/characters/fields |
| `sort`  | Sort lines (`-n`, `-d`, `-u`) |
| `uniq`  | Filter adjacent duplicate lines (`-c`) |
| `wc`    | Count lines/words/bytes |
| `head`  | Output first lines/bytes |
| `tail`  | Output last lines/bytes |
| `tr`    | Translate/delete characters |
| `tee`   | Redirect output to files (`-a`) |
| `xargs` | Build and execute commands from stdin |
| `echo`  | Print strings (`-n`) |
| `printf`| Format and print data |
| `paste` | Merge lines of files |
| `join`  | Join lines on a common field |

### System Information

| Command | Description |
|---------|-------------|
| `uname` | Print system information |
| `hostname` | Print host name |
| `whoami` | Print current user |

### Networking

| Command | Description |
|---------|-------------|
| `wget`  | HTTP/HTTPS file download |
| `telnet`| TCP connection test |
| `nc`    | Netcat-style TCP/UDP tools |

### Compression & Archiving

| Command | Description |
|---------|-------------|
| `tar`   | Create/extract/list tar archives |
| `gzip` / `gunzip` | gzip compression/decompression |
| `zip` / `unzip` | ZIP archive handling |
| `bzip2` | bzip2 compression/decompression |
| `xz`    | xz compression/decompression |

### Utilities

| Command | Description |
|---------|-------------|
| `date`  | Display/format date and time |
| `cal`   | Display a calendar |
| `sleep` | Pause execution |
| `timeout` | Run a command with a time limit |
| `watch` | Repeatedly run a command |
| `clear` / `reset` | Clear/reset terminal |

## Installation

### Prebuilt Binaries

Download prebuilt binaries for your platform from the [Releases](https://github.com/user/linux_coreutils/releases) page.

### Build from Source

Requirements:
- Go 1.24 or later

```bash
# Clone the repository
git clone https://github.com/user/linux_coreutils.git
cd linux_coreutils

# Build all commands
go build ./cmd/...

# Run tests
go test ./...

# Install to /usr/local/bin
make install
```

### Cross-Compilation

Use Go's built-in cross-compilation to target other platforms:

```bash
# Linux amd64
GOOS=linux GOARCH=amd64 go build -o dist/linux_amd64/ ./cmd/...

# Windows amd64
GOOS=windows GOARCH=amd64 go build -o dist/windows_amd64/ ./cmd/...

# macOS arm64
GOOS=darwin GOARCH=arm64 go build -o dist/darwin_arm64/ ./cmd/...
```

## Usage Examples

```bash
# List files
./ls -la

# Count words
./wc -w file.txt

# Find Go files
./find . -name "*.go"

# Compress a file
./gzip file.txt

# Download a file
./wget https://example.com/file.txt
```

Each command supports `-h` or `--help` for detailed usage:

```bash
./ls --help
```

## Project Structure

```text
linux_coreutils/
├── cmd/              # One directory per command
├── pkg/              # Shared packages
│   ├── common/       # Errors, path, I/O helpers
│   └── cliutil/      # CLI helpers
├── internal/         # Internal test helpers
├── tests/            # Shell integration tests
├── docs/             # Manual pages
├── scripts/          # CI/build scripts
├── .github/workflows/# GitHub Actions
├── go.mod
├── Makefile
└── README.md
```

## Testing

Run the full test suite:

```bash
go test ./...
```

Run with the race detector:

```bash
go test -race ./...
```

Run integration tests:

```bash
make test
```

## Continuous Integration & Releases

The project uses GitHub Actions for automated releases:

- **Release workflow**: Manually triggered via `workflow_dispatch`.
- Builds binaries for `linux/amd64`, `linux/arm64`, `windows/amd64`, `darwin/amd64`, and `darwin/arm64`.
- Publishes the binaries to a GitHub Release.

To create a release:

1. Go to **Actions** → **Release**.
2. Click **Run workflow**.
3. Enter a tag (e.g., `v1.0.0`) and optional settings.
4. The workflow will build and publish the release automatically.

## Compatibility

| Platform | Requirements |
|----------|--------------|
| Linux    | Kernel 3.10+, Go 1.24+ |
| Windows  | Windows 10+ |
| macOS    | macOS 12+ (Intel/Apple Silicon) |

## Notes

- Output format aims to match GNU Coreutils but is not guaranteed to be identical.
- Exit codes follow POSIX conventions: `0` for success, non-zero for failure.
- Error messages are printed to `stderr` in English.

## License

[MIT](LICENSE)

## Acknowledgments

Inspired by GNU Coreutils and the Unix philosophy of small, composable tools.

# System Info CLI 🖥️

**Phase 1: Foundations** | **Duration:** 2-3 days | **Status:** ⬜ Not Started

## Description

A cross-platform command-line tool to display comprehensive system information including CPU, memory, disk usage, running processes, and network interfaces.

## Features

- ✅ Display CPU information and usage
- ✅ Show memory (RAM) statistics
- ✅ Disk usage for all mounted volumes
- ✅ List running processes
- ✅ Network interface information
- ✅ Cross-platform support (Windows/Linux/macOS)
- ✅ Formatted output (tables or JSON)

## Learning Outcomes

- [x] Working with system calls
- [x] Third-party packages (`gopsutil`)
- [x] Cross-platform development
- [x] Formatted console output
- [x] Goroutines for concurrent data fetching
- [x] CLI argument parsing

## Tech Stack

- **System Info**: `github.com/shirou/gopsutil/v3`
- **Table Formatting**: `github.com/olekukonko/tablewriter`
- **CLI**: `github.com/spf13/cobra` (optional)

## Installation

```bash
# Clone the project
cd 03-system-info-cli

# Install dependencies
go mod download

# Build
go build -o sysinfo main.go

# Run
./sysinfo
```

## Usage

### Basic Usage

```bash
# Display all system info
./sysinfo

# Display specific info
./sysinfo --cpu
./sysinfo --memory
./sysinfo --disk
./sysinfo --network
./sysinfo --processes

# Output as JSON
./sysinfo --json

# Continuous monitoring (refresh every 2 seconds)
./sysinfo --watch --interval 2
```

## Output Example

```
╔════════════════════════════════════════════════════════╗
║           SYSTEM INFORMATION                           ║
╚════════════════════════════════════════════════════════╝

CPU Information:
┌─────────────────┬──────────────────────────────────────┐
│ Physical Cores  │ 8                                    │
│ Logical Cores   │ 16                                   │
│ Model Name      │ Intel Core i7-10700K                 │
│ CPU Usage       │ 23.5%                                │
└─────────────────┴──────────────────────────────────────┘

Memory Information:
┌─────────────────┬──────────────────────────────────────┐
│ Total           │ 16.0 GB                              │
│ Used            │ 8.2 GB (51.2%)                       │
│ Available       │ 7.8 GB                               │
└─────────────────┴──────────────────────────────────────┘

Disk Information:
┌──────────────┬────────────┬──────────┬──────────┬───────┐
│ DEVICE       │ MOUNT      │ TOTAL    │ USED     │ USAGE │
├──────────────┼────────────┼──────────┼──────────┼───────┤
│ /dev/sda1    │ /          │ 500 GB   │ 320 GB   │ 64%   │
│ /dev/sdb1    │ /media/sdb │ 1.0 TB   │ 450 GB   │ 45%   │
└──────────────┴────────────┴──────────┴──────────┴───────┘

Network Interfaces:
┌───────────┬───────────────────────┬──────────────────┐
│ NAME      │ IP ADDRESS            │ STATUS           │
├───────────┼───────────────────────┼──────────────────┤
│ eth0      │ 192.168.1.100         │ UP               │
│ wlan0     │ 192.168.1.101         │ UP               │
│ lo        │ 127.0.0.1             │ UP               │
└───────────┴───────────────────────┴──────────────────┘

Top 5 Processes by Memory:
┌────────┬─────────────────────────┬──────────┬──────────┐
│ PID    │ NAME                    │ CPU %    │ MEMORY   │
├────────┼─────────────────────────┼──────────┼──────────┤
│ 1234   │ chrome                  │ 12.3%    │ 2.1 GB   │
│ 5678   │ code                    │ 8.5%     │ 1.8 GB   │
│ 9012   │ docker                  │ 5.2%     │ 1.2 GB   │
└────────┴─────────────────────────┴──────────┴──────────┘
```

## Project Structure

```
03-system-info-cli/
├── main.go                      # Main application
├── go.mod                       # Dependencies
├── go.sum                       # Checksums
├── README.md                    # This file
└── .gitignore
```

## Cross-Platform Build

```bash
# Build for Windows
GOOS=windows GOARCH=amd64 go build -o sysinfo.exe main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o sysinfo-linux main.go

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o sysinfo-mac main.go

# Build for all platforms (using script)
./build-all.sh
```

## Development

### Run directly

```bash
go run main.go
```

### Test

```bash
go test -v ./...
```

## Advanced Features (Future)

- [ ] Export to file (JSON, CSV, HTML)
- [ ] Temperature monitoring
- [ ] Battery status (for laptops)
- [ ] GPU information
- [ ] Historical data tracking
- [ ] Alerts on threshold breach
- [ ] Web dashboard
- [ ] Remote monitoring

## Resources

- [gopsutil Documentation](https://github.com/shirou/gopsutil)
- [tablewriter Documentation](https://github.com/olekukonko/tablewriter)
- [Cobra Documentation](https://github.com/spf13/cobra)

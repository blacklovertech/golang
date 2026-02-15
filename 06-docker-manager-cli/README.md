# Docker Container Manager CLI 🐳

**Phase 2: Intermediate** | **Duration:** 4-5 days | **Status:** ⬜ Not Started

## Description

CLI tool to manage Docker containers with real-time monitoring.

## Features

- List all containers with status
- Start, stop, restart containers
- View container logs (real-time)
- Show container resource stats (CPU, memory)
- Remove containers
- Pull images

## Learning Outcomes

- Docker SDK for Go
- Working with Docker API
- Streaming logs
- Context and cancellation
- Real-time stats monitoring
- CLI design patterns

## Tech Stack

- `github.com/docker/docker` (Docker SDK)
- `cobra` for CLI (optional)

## Quick Start

```bash
cd 06-docker-manager-cli
go mod init github.com/blacklovertech/06-docker-manager-cli
go get github.com/docker/docker github.com/spf13/cobra
go build -o dockman.exe
```

## CLI Commands

```
dockman list                    - List containers
dockman start <id>              - Start container
dockman stop <id>               - Stop container
dockman logs <id>               - View logs
dockman stats <id>              - Show stats
dockman rm <id>                 - Remove container
```

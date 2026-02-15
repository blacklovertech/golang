# Windows Service Monitor Daemon 🔧

**Phase 3: Advanced** | **Duration:** 7-10 days | **Status:** ⬜ Not Started

## Description

System daemon to monitor and auto-restart Windows services with web dashboard.

## Features

- Background daemon running as Windows service
- Monitor configured Windows services
- Auto-restart failed services
- Configurable health check intervals
- REST API for configuration
- Email/Slack notifications on failures
- Service status dashboard (React)
- Logs of all service actions
- Start/stop monitoring via API

## Learning Outcomes

- Windows service creation in Go
- Windows API calls (`golang.org/x/sys/windows`)
- Service management and monitoring
- Background task scheduling
- Email/Slack integration
- System daemon patterns
- Alert mechanisms

## Tech Stack

- `golang.org/x/sys/windows` for Windows APIs
- `kardianos/service` for service management
- Gin for REST API
- PostgreSQL for logs
- React dashboard

## Quick Start

```bash
cd 08-windows-service-monitor
go mod init github.com/blacklovertech/08-windows-service-monitor
go get golang.org/x/sys/windows github.com/kardianos/service
go build -o servicemon.exe cmd/daemon/main.go
servicemon.exe install
servicemon.exe start
```

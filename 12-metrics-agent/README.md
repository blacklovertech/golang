# Metrics Collection Agent 📈

**Phase 4: Production Ready** | **Duration:** 5-7 days | **Status:** ⬜ Not Started

## Description

Lightweight monitoring agent for system metrics with Prometheus exporter support.

## Features

- Collect CPU, memory, disk, network metrics
- Export to Prometheus format
- Run as system service (systemd/Windows service)
- Plugin system for custom metrics
- Configurable collection intervals
- Low resource footprint
- Multiple export formats (Prometheus, JSON, InfluxDB)

## Learning Outcomes

- System metrics collection
- Prometheus exporter format
- Time-series data
- Running as system service
- Plugin architecture
- Efficient data collection

## Tech Stack

- `gopsutil` for metrics
- `prometheus/client_golang`
- `kardianos/service`

## Quick Start

```bash
cd 12-metrics-agent
go mod init github.com/blacklovertech/12-metrics-agent
go get github.com/shirou/gopsutil/v3 github.com/prometheus/client_golang
go build -o metricsagent.exe cmd/agent/main.go
```

## Configuration

Edit `config/agent.yaml`:

```yaml
collection_interval: 60s
exports:
  - prometheus
  - json
metrics:
  - cpu
  - memory
  - disk
  - network
```

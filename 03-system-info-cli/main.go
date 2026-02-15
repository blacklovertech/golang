package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║           SYSTEM INFORMATION                           ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝\n")

	// Display all system information
	displayCPUInfo()
	displayMemoryInfo()
	displayDiskInfo()
	displayNetworkInfo()
	displayTopProcesses(5)
}

func displayCPUInfo() {
	fmt.Println("CPU Information:")
	
	// Get CPU info
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v\n", err)
		return
	}

	// Get CPU usage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Error getting CPU usage: %v\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetBorder(true)

	if len(cpuInfo) > 0 {
		table.Append([]string{"Model Name", cpuInfo[0].ModelName})
		table.Append([]string{"Physical Cores", fmt.Sprintf("%d", runtime.NumCPU())})
		table.Append([]string{"Logical Cores", fmt.Sprintf("%d", runtime.GOMAXPROCS(0))})
	}
	
	if len(cpuPercent) > 0 {
		table.Append([]string{"CPU Usage", fmt.Sprintf("%.2f%%", cpuPercent[0])})
	}

	table.Render()
	fmt.Println()
}

func displayMemoryInfo() {
	fmt.Println("Memory Information:")

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Value"})
	table.SetBorder(true)

	table.Append([]string{"Total", formatBytes(vmStat.Total)})
	table.Append([]string{"Used", fmt.Sprintf("%s (%.2f%%)", formatBytes(vmStat.Used), vmStat.UsedPercent)})
	table.Append([]string{"Available", formatBytes(vmStat.Available)})
	table.Append([]string{"Free", formatBytes(vmStat.Free)})

	table.Render()
	fmt.Println()
}

func displayDiskInfo() {
	fmt.Println("Disk Information:")

	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Printf("Error getting disk info: %v\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Device", "Mount Point", "Total", "Used", "Free", "Usage %"})
	table.SetBorder(true)

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}

		table.Append([]string{
			partition.Device,
			partition.Mountpoint,
			formatBytes(usage.Total),
			formatBytes(usage.Used),
			formatBytes(usage.Free),
			fmt.Sprintf("%.2f%%", usage.UsedPercent),
		})
	}

	table.Render()
	fmt.Println()
}

func displayNetworkInfo() {
	fmt.Println("Network Interfaces:")

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error getting network info: %v\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "IP Addresses", "Status"})
	table.SetBorder(true)

	for _, iface := range interfaces {
		status := "DOWN"
		if len(iface.Flags) > 0 {
			status = "UP"
		}

		ips := ""
		for i, addr := range iface.Addrs {
			if i > 0 {
				ips += ", "
			}
			ips += addr.Addr
		}

		if ips == "" {
			ips = "N/A"
		}

		table.Append([]string{iface.Name, ips, status})
	}

	table.Render()
	fmt.Println()
}

func displayTopProcesses(count int) {
	fmt.Printf("Top %d Processes by Memory:\n", count)

	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("Error getting processes: %v\n", err)
		return
	}

	type ProcessInfo struct {
		PID     int32
		Name    string
		CPU     float64
		Memory  float32
	}

	var processInfos []ProcessInfo

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		memPercent, err := p.MemoryPercent()
		if err != nil {
			continue
		}

		processInfos = append(processInfos, ProcessInfo{
			PID:    p.Pid,
			Name:   name,
			CPU:    cpuPercent,
			Memory: memPercent,
		})

		if len(processInfos) >= count*10 {
			break
		}
	}

	// Sort by memory (simple bubble sort for demonstration)
	for i := 0; i < len(processInfos)-1; i++ {
		for j := 0; j < len(processInfos)-i-1; j++ {
			if processInfos[j].Memory < processInfos[j+1].Memory {
				processInfos[j], processInfos[j+1] = processInfos[j+1], processInfos[j]
			}
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PID", "Name", "CPU %", "Memory %"})
	table.SetBorder(true)

	for i := 0; i < count && i < len(processInfos); i++ {
		table.Append([]string{
			fmt.Sprintf("%d", processInfos[i].PID),
			processInfos[i].Name,
			fmt.Sprintf("%.2f%%", processInfos[i].CPU),
			fmt.Sprintf("%.2f%%", processInfos[i].Memory),
		})
	}

	table.Render()
	fmt.Println()
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

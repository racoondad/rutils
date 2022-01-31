package sysinfo

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	b  = 1
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
)

type Machine struct {
	SysOs SysOs `json:"sysOs"`
	Cpu   Cpu   `json:"cpu"`
	Ram   Ram   `json:"ram"`
	Disk  Disk  `json:"disk"`
}

type SysOs struct {
	// 系统
	GOOS string `json:"goos"`
	// 核心数
	CPUCores int `json:"cpuCores"`
	// 平台
	OSPlatform string `json:"osPlatform"`
	Compiler   string `json:"compiler"`
	// go版本号
	GoVersion string `json:"goVersion"`
	// 协程数量
	GoroutineTotal int `json:"goroutineTotal"`
}

type Cpu struct {
	Usage []float64 `json:"usage"`
	Cores int       `json:"cores"`
}

type Ram struct {
	UsedMB      int `json:"usedMB"`
	TotalMB     int `json:"totalMB"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMB"`
	UsedGB      int `json:"usedGB"`
	TotalMB     int `json:"totalMB"`
	TotalGB     int `json:"totalGB"`
	UsedPercent int `json:"usedPercent"`
}

//@function: GetOSInfo
//@description: OS信息
//@return: o Os, err error
func GetOSInfo() (o SysOs) {
	o.GOOS = runtime.GOOS
	o.CPUCores = runtime.NumCPU()
	o.OSPlatform = runtime.GOARCH
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.GoroutineTotal = runtime.NumGoroutine()
	return o
}

//@function: GetCPUInfo
//@description: CPU信息
//@return: c Cpu, err error
func GetCPUInfo() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Usage = cpus
	}
	return c, nil
}

//@function: GetRAMInfo
//@description: ARM信息
//@return: r Rrm, err error
func GetRAMInfo() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / mb
		r.TotalMB = int(u.Total) / mb
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}

//@function: GetDiskInfo
//@description: 硬盘信息
//@return: d Disk, err error
func GetDiskInfo() (d Disk, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / mb
		d.UsedGB = int(u.Used) / gb
		d.TotalMB = int(u.Total) / mb
		d.TotalGB = int(u.Total) / gb
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, nil
}

package resp

// RespCPU CPU信息
type RespCPU struct {
	CpuCounts      string `json:"cpu_counts"`       // CPU物理核心数
	CpuUsedpercent string `json:"cpu_used_percent"` // CPU使用率
}

// RespMemory 内存信息
type RespMemory struct {
	MemTotal       string `json:"mem_total"`        // 全部内存，单位GB
	MemUsed        string `json:"mem_used"`         // 已使用内存，单位GB
	MemFree        string `json:"mem_free"`         // 空闲内存，单位GB
	MemUsedPercent string `json:"mem_used_percent"` // 内存使用率
}

// RespDisk 磁盘信息
type RespDisk struct {
	DiskTotal string `json:"disk_total"` // 全部硬盘容量，单位GB
	DiskUsed  string `json:"disk_used"`  // 已使用硬盘，单位GB
	DiskFree  string `json:"disk_free"`  // 空闲硬盘，单位GB
}

// RespSystem 系统信息
type RespSystem struct {
	CPU    RespCPU    `json:"cpu"`    // CPU
	Memory RespMemory `json:"memory"` // 内存
	Disk   RespDisk   `json:"disk"`   // 磁盘
}

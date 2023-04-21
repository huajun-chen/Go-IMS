package resstruct

// CPUReturn CPU信息
type CPUReturn struct {
	CpuCounts      string `json:"cpu_counts"`       // CPU物理核心数
	CpuUsedpercent string `json:"cpu_used_percent"` // CPU使用率
}

// MemoryReturn 内存信息
type MemoryReturn struct {
	MemTotal       string `json:"mem_total"`        // 全部内存，单位GB
	MemUsed        string `json:"mem_used"`         // 已使用内存，单位GB
	MemFree        string `json:"mem_free"`         // 空闲内存，单位GB
	MemUsedPercent string `json:"mem_used_percent"` // 内存使用率
}

// DiskReturn 磁盘信息
type DiskReturn struct {
	DiskTotal string `json:"disk_total"` // 全部硬盘容量，单位GB
	DiskUsed  string `json:"disk_used"`  // 已使用硬盘，单位GB
	DiskFree  string `json:"disk_free"`  // 空闲硬盘，单位GB
}

// SystemReturn 系统信息
type SystemReturn struct {
	CPU    CPUReturn    `json:"cpu"`    // CPU
	Memory MemoryReturn `json:"memory"` // 内存
	Disk   DiskReturn   `json:"disk"`   // 磁盘
}

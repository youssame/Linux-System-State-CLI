// Package memory: The system memory state
package memory

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/youssame/linux-system-state-cli/pkg"
	"sync"
)

func (memory MemoryInfo) DisplayTitle() string {
	return `
Memory Information:
-------------------`
}
func (memory MemoryInfo) DisplayInfo() string {
	return fmt.Sprintf(`Total Memory    : %d GB
Used Memory     : %d GB
Free Memory     : %d GB
Swap Usage      : %d%%`, memory.total, memory.used, memory.free, memory.swap)
}

// GetMemoryInfos Get the system state of memory
func GetMemoryInfos(wg *sync.WaitGroup) MemoryInfo {
	defer wg.Done()

	var ret MemoryInfo
	swap, err := mem.SwapMemory()
	pkg.SysErrChecker(err, "SwapMemory")
	vm, err := mem.VirtualMemory()
	pkg.SysErrChecker(err, "VirtualMemory")
	ret.swap = swap.Total / 1073741824
	ret.free = vm.Free / 1073741824
	ret.used = vm.Used / 1073741824
	ret.total = vm.Total / 1073741824
	return ret
}

// MemoryInfo The structure system state's memory
type MemoryInfo struct {
	total uint64
	used  uint64
	free  uint64
	swap  uint64
}

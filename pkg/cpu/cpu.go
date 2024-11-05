// Package cpu: The system cpu state
package cpu

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/youssame/linux-system-state-cli/pkg"
	"golang.org/x/sys/unix"
	"sync"
	"time"
)

func (cpu CPUInfo) DisplayTitle() string {
	return `
CPU Information:
----------------`
}

func (cpu CPUInfo) DisplayInfo() string {
	return fmt.Sprintf(`Model           : %s
Cores           : %d
Frequency       : %f GHz
Usage           : %f %%`, cpu.model, cpu.cores, cpu.frequency, cpu.usage)
}

// GetCPUInfos Get CPU system information
func GetCPUInfos(wg *sync.WaitGroup) CPUInfo {
	defer wg.Done()

	var ret CPUInfo
	model, err := unix.Sysctl("machdep.cpu.brand_string")
	pkg.SysErrChecker(err, "model")
	ret.model = model
	cores, err := unix.SysctlUint32("machdep.cpu.core_count")
	pkg.SysErrChecker(err, "cores")
	ret.cores = int32(cores)
	fmt.Println(unix.Getwd())
	cpuFrequency, err := unix.SysctlUint64("hw.cpufrequency")
	pkg.SysErrChecker(err, "cpuFrequency")
	if cpuFrequency == uint64(0) {
		ret.frequency = float64(cpuFrequency) / 1000000.0
	}
	usagePer, err := cpu.Percent(time.Second, false)
	pkg.SysErrChecker(err, "usage")
	ret.usage = usagePer[0]
	return ret
}

// CPUInfo The structure of the CPU state
type CPUInfo struct {
	model     string
	cores     int32
	frequency float64
	usage     float64
}

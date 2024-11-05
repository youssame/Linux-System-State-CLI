// Package storage: The system storage state
package storage

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/youssame/linux-system-state-cli/pkg"
	"sync"
)

// GetStorageInfos Get the system storage state
func GetStorageInfos(wg *sync.WaitGroup) StorageInfo {
	defer wg.Done()

	var ret StorageInfo
	usage, err := disk.Usage("/")
	pkg.SysErrChecker(err, "DiskUsage")
	ret.free = usage.Free / 1073741824
	ret.used = usage.Used / 1073741824
	ret.total = usage.Total / 1073741824
	return ret
}

func (storage StorageInfo) DisplayTitle() string {
	return `
Storage Information:
--------------------`
}
func (storage StorageInfo) DisplayInfo() string {
	return fmt.Sprintf(`Total Storage   : %d GB
Used Storage    : %d GB
Free Storage    : %d GB`, storage.total, storage.used, storage.free)
}

// StorageInfo The storage state
type StorageInfo struct {
	total uint64
	used  uint64
	free  uint64
}

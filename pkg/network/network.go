// Package network: The system network state
package network

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/youssame/linux-system-state-cli/pkg"
	"sync"
)

func (network NetworkInfo) DisplayTitle() string {
	return `
Network Information:
--------------------`
}
func (network NetworkInfo) DisplayInfo() string {
	return fmt.Sprintf(`IP Address      : %s
Data Sent       : %d GB
Data Received   : %d GB`, network.ip, network.sent_data, network.received_data)
}

// GetNetworkInfos Get Network state information
func GetNetworkInfos(wg *sync.WaitGroup) NetworkInfo {
	defer wg.Done()

	var ret NetworkInfo
	inter, err := net.Interfaces()
	pkg.SysErrChecker(err, "NetInterfaces")
	counter, err := net.IOCounters(false)
	pkg.SysErrChecker(err, "NetIOCounters")
	ret.ip = inter[0].Addrs[0].Addr
	ret.sent_data = counter[0].BytesSent / 1073741824
	ret.received_data = counter[0].BytesRecv / 1073741824
	return ret
}

// NetworkInfo the structure of network info
type NetworkInfo struct {
	ip            string
	sent_data     uint64
	received_data uint64
}

package main

import (
	"flag"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/youssame/linux-system-state-cli/pkg/cpu"
	"github.com/youssame/linux-system-state-cli/pkg/memory"
	"github.com/youssame/linux-system-state-cli/pkg/network"
	"github.com/youssame/linux-system-state-cli/pkg/printer"
	"github.com/youssame/linux-system-state-cli/pkg/storage"
	"sync"
	"time"
)

func printsystemState(cpuArg, memoryArg, storageArg, networkArg, logArg bool) string {
	var res = new(string)
	*res = `===============================
System Information Overview
===============================`
	var wg sync.WaitGroup
	if cpuArg {
		wg.Add(1)
		go printer.DisplayFullInfo(cpu.GetCPUInfos(&wg), &*res, logArg)
	}
	if memoryArg {
		wg.Add(1)
		go printer.DisplayFullInfo(memory.GetMemoryInfos(&wg), &*res, logArg)
	}
	if storageArg {
		wg.Add(1)
		go printer.DisplayFullInfo(storage.GetStorageInfos(&wg), &*res, logArg)
	}
	if networkArg {
		wg.Add(1)
		go printer.DisplayFullInfo(network.GetNetworkInfos(&wg), &*res, logArg)
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
	return *res + `
===============================
End of System Information
===============================`

}
func main() {
	var cpuArg, memoryArg, networkArg, storageArg, log, monitor bool
	flag.BoolVar(&cpuArg, "cpu", false, "CPU Information")
	flag.BoolVar(&memoryArg, "memoryArg", false, "Memory Information")
	flag.BoolVar(&networkArg, "network", false, "Network Information")
	flag.BoolVar(&storageArg, "storage", false, "Storage Information")
	flag.BoolVar(&log, "log", false, "Log the output in a file")
	flag.BoolVar(&monitor, "monitor", false, "Refresh the system state every few seconds")
	flag.Parse()

	if !cpuArg && !memoryArg && !networkArg && !storageArg {
		cpuArg = true
		memoryArg = true
		networkArg = true
		storageArg = true
	}

	if monitor {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				tm.MoveCursor(1, 1)
				fmt.Println(printsystemState(cpuArg, memoryArg, storageArg, networkArg, log))
				tm.Flush()
			}
		}
	} else {
		fmt.Println(printsystemState(cpuArg, memoryArg, storageArg, networkArg, log))
	}
}

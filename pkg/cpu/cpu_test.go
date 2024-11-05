package cpu

import (
	"testing"
)

func TestDisplayTitle(t *testing.T) {
	cpuInfoMock := CPUInfo{}
	res := cpuInfoMock.DisplayTitle()
	exp := `
CPU Information:
----------------`
	if res != exp {
		t.Errorf("DisplayTitle failed, expected %s, got %s", exp, res)
	}
}

func TestDisplayInfo(t *testing.T) {
	cpuInfoMock := getMockedCPUInfos()
	res := cpuInfoMock.DisplayInfo()
	exp := `Model           : model
Cores           : 1
Frequency       : 3.000000 GHz
Usage           : 2.000000 %`
	if res != exp {
		t.Errorf("DisplayInfo failed, expected %s, got %s", exp, res)
	}
}

func getMockedCPUInfos() CPUInfo {
	return CPUInfo{
		model:     "model",
		cores:     1,
		frequency: 3,
		usage:     2}
}

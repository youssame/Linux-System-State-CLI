package memory

import "testing"

func TestDisplayTitle(t *testing.T) {
	res := getMockedMemoryInfos().DisplayTitle()
	exp := `
Memory Information:
-------------------`
	if res != exp {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", exp, res)
	}
}

func TestGetMemoryInfos(t *testing.T) {
	res := getMockedMemoryInfos().DisplayInfo()
	exp := `Total Memory    : 0 GB
Used Memory     : 1 GB
Free Memory     : 2 GB
Swap Usage      : 3%`
	if res != exp {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", exp, res)
	}
}

func getMockedMemoryInfos() MemoryInfo {
	return MemoryInfo{
		total: 0,
		used:  1,
		free:  2,
		swap:  3,
	}
}

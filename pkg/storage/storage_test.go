package storage

import "testing"

func TestStorageDisplayTitle(t *testing.T) {
	res := getMockedStorageInfo().DisplayTitle()
	exp := `
Storage Information:
--------------------`
	if res != exp {
		t.Errorf("got %s, want %s", res, exp)
	}
}

func TestStorageDisplayInfo(t *testing.T) {
	res := getMockedStorageInfo().DisplayInfo()
	exp := `Total Storage   : 1 GB
Used Storage    : 2 GB
Free Storage    : 3 GB`
	if res != exp {
		t.Errorf("got %s, want %s", res, exp)
	}
}

func getMockedStorageInfo() StorageInfo {
	return StorageInfo{
		total: 1,
		used:  2,
		free:  3,
	}
}

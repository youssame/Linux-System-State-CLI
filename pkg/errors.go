package pkg

import (
	"fmt"
	"os"
)

func sysErr(info string) error {
	return fmt.Errorf("error: can't retreive %s info", info)
}

// SysErrChecker Display the system error while retrieving the system state info
func SysErrChecker(error error, info string) {
	if os.Getenv("MODE") != "DEV" {
		return
	}
	if error != nil {
		fmt.Println(sysErr(info))
		fmt.Println(error.Error())
	}
}

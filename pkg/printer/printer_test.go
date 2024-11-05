package printer

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	res := getMockedState().Print()
	exp := getMockedState().title + "\n" + getMockedState().description
	if res != exp {
		t.Errorf("Print() returned %s, expected %s", res, exp)
	}
}
func TestPrinterDisplayFullInfoWithoutLog(t *testing.T) {
	var str = ""
	DisplayFullInfo(getMockedState(), &str, false)
	exp := `
title
info`

	if str != exp {
		t.Errorf("Print() returned %s, expected %s", str, exp)
	}
}
func TestPrinterDisplayFullInfoWithLog(t *testing.T) {
	defer deleteLogFile()
	var str = ""
	DisplayFullInfo(getMockedState(), &str, true)
	exp := `
title
info`

	if str != exp {
		t.Errorf("Print() returned %s, expected %s", str, exp)
	}
	status, content := getLogFileContent()
	if !status {
		t.Errorf("Log printer didn't work well, check the log please.")
	}
	exp = `titleinfo
`
	if content != exp {
		t.Errorf("Log file don't have the correct content\n returned %s, expected %s", content, exp)
	}
}

// helpers
func (s State) DisplayInfo() string {
	return "info"
}
func (s State) DisplayTitle() string {
	return "title"
}
func getMockedState() State {
	state := State{
		title:       "State's title",
		description: "State's desciption"}
	return state
}
func checkLogFileExists() bool {
	path := "./info.log"
	_, err := os.Stat(path)
	return err == nil
}
func getLogFileContent() (bool, string) {
	path := "./info.log"
	if !checkLogFileExists() {
		fmt.Println("Error finding the log file")
		return false, ""
	}
	cont, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error opening the log file")
		return false, ""
	}
	str := string(cont)
	return true, str
}
func deleteLogFile() {
	path := "./info.log"
	if !checkLogFileExists() {
		fmt.Println("Error finding log file")
	}
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error deleting log file")
	}
}

// Package printer: Printing the system state in the console
package printer

import (
	"github.com/youssame/linux-system-state-cli/pkg"
	"log"
	"os"
)

// StatePrinter State printer interface
type StatePrinter interface {
	DisplayTitle() string
	DisplayInfo() string
}

// State printer
type State struct {
	title       string
	description string
}

func (s State) Print() string {
	return s.title + "\n" + s.description
}

// PrintLogState Print and log the state
func PrintLogState(state string) {
	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	pkg.SysErrChecker(err, "OpenFile")
	_, err = file.WriteString(state + "\n")
	pkg.SysErrChecker(err, "WriteInLogFile")
	defer func(file *os.File) {
		err := file.Close()
		pkg.SysErrChecker(err, "CloseLogFile")
	}(file)
	log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// DisplayFullInfo Display a state
func DisplayFullInfo(printer StatePrinter, res *string, logg bool) {
	s := State{
		title:       printer.DisplayTitle(),
		description: printer.DisplayInfo()}
	state := s.Print()
	*res += "\n" + state
	if logg {
		PrintLogState(s.title + s.description)
	}
}

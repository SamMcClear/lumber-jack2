package cpu

import (
	"fmt"
	"runtime"
)

func GetUsage() string {
	cpuCount := runtime.NumCPU()
	//  CPU usage logic here
	return fmt.Sprintf("Number of CPU threads: %d", cpuCount)
}

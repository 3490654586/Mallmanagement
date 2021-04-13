package utlls

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

func Cpu() {
 data,_ :=cpu.Times(false)
 fmt.Println(data)
}

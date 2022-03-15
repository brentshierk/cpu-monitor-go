package widget

import (
	"cpu-monitor-go/util"
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)

type StatsMemory struct {
	Total float64
	Available float64
	Used float64
	UsedPercent float64


}

func (m *StatsMemory) GetStats()  StatsMemory{
	return StatsMemory{
		Total:       m.GetTotal(),
		Available:   m.GetAvailable(),
		Used:        m.GetUsed(),
		UsedPercent: m.GetUsedPercent(),
	}
}

func (m *StatsMemory) GetTotal() float64{
	ms, err := mem.VirtualMemory()
	 s, _ := util.ConvertBytes(ms.Total)
	if err != nil {
		fmt.Println("error unable to get memory ")
	}
	return s
}
func (m *StatsMemory) GetAvailable() float64{
	ms, err := mem.VirtualMemory()
	s, _ := util.ConvertBytes(ms.Available)
	if err != nil {
		fmt.Println("error unable to get memory ")
	}
	return s
}
func (m *StatsMemory) GetUsed() float64{
	ms, err := mem.VirtualMemory()
	s, _ := util.ConvertBytes(ms.Available)
	if err != nil {
		fmt.Println("error unable to get memory ")
	}
	return s
}
func (m *StatsMemory) GetUsedPercent() float64{
	ms, err := mem.VirtualMemory()
	s, _ := util.ConvertBytes(ms.Available)
	if err != nil {
		fmt.Println("error unable to get memory ")
	}
	return s
}

package main

import (
	utils "cpu-monitor-go/util"
	"fmt"
	ui "github.com/gizak/termui/v3"
	//"github.com/gizak/termui/v3/widgets"
	"log"
	"cpu-monitor-go/widgets"
)

func main() {
	utils.TempToFarenhiet(3)

	fmt.Println(utils.TempToFarenhiet(3))
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	s:= widget.GetS
	fmt.Println(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
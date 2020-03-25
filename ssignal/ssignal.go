package ssignal

import (
	"os"
	"os/signal"
)

type HandleSignal func(s os.Signal) bool

func waitSignal(handle HandleSignal, sig ...os.Signal)  {
	c := make(chan os.Signal)
	signal.Notify(c, sig...)

	//the other implementation is to use two defers which are possible
	defer func() {
		signal.Stop(c)
		close(c)
	}()

	for s := range c {
		if handle == nil || !handle(s) {
			break
		}
	}
}

func WaitCtrlC(handle HandleSignal) {
	waitSignal(handle)
}



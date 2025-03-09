package main

import "fmt"

func main() {
	receiver := new(Receiver)

	ReportA := &ConcreteReportA{receiver: receiver}
	ReportB := &ConcreteReportB{receiver: receiver}

	invoker := new(Invoker)
	// Report A then Report B
	invoker.Schedule(ReportA)
	invoker.Schedule(ReportB)
	invoker.Run()
	fmt.Println("Next")
	// Report B then Report A
	invoker.Schedule(ReportB)
	invoker.Schedule(ReportA)
	invoker.Run()
}

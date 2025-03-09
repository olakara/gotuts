package main

// The command interface
type Report interface {
	Execute()
}

// The Concrete Command
type ConcreteReportA struct {
	receiver *Receiver
}

func (c *ConcreteReportA) Execute() {
	c.receiver.Action("Report A")
}

// The Concrete Command
type ConcreteReportB struct {
	receiver *Receiver
}

func (c *ConcreteReportB) Execute() {
	c.receiver.Action("Report B")
}

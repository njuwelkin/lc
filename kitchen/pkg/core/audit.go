package core

import (
	"fmt"
	"io"
	"os"
)

type audit struct {
	out io.Writer
}

func newAudit(out io.Writer) *audit {
	if out == nil {
		out = os.Stdout
	}
	return &audit{out: out}
}

func (a *audit) PrintEvent(order *Order, event Event) {
	a.printLine("\n\n==========================================================================")
	switch event {
	case Accepted:
		a.printLine(fmt.Sprintf("Accept order '%s'", order.ID))
	case Cooked:
		a.printLine(fmt.Sprintf("Order '%s' cooked", order.ID))
	case Moved:
		a.printLine(fmt.Sprintf("Order '%s' moved to single temp shelf", order.ID))
	case Picked:
		a.printLine(fmt.Sprintf("Order '%s' picked", order.ID))
	case Discarded:
		a.printLine(fmt.Sprintf("Order '%s' discarded", order.ID))
	case Delivered:
		a.printLine(fmt.Sprintf("Order '%s' delivered", order.ID))
	}
}

func (a *audit) PrintShelfContent(content [CountTempType + 1][]*Order) {
	for t := Hot; t <= CountTempType; t++ {
		a.printSingleShelf(t, content[t])
	}
}

func (a *audit) PrintStatistic(delieved, discarded int) {
	a.printLine(fmt.Sprintf("%d orders delivered, %d orders discarded.", delieved, discarded))
}

func (a *audit) printLine(s string) {
	io.WriteString(a.out, s+"\n")
}

var shelfName = [CountTempType + 1]string{
	"Hot",
	"Cold",
	"Fronzen",
	"Overflow",
}

func (a *audit) printSingleShelf(t OrderTemp, content []*Order) {
	name := shelfName[t]
	a.printLine(fmt.Sprintf("------------------------ %s --------------------------", name))
	for i, order := range content {
		a.printLine(fmt.Sprintf("%d: %s", i+1, order.ID))
	}
}

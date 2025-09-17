package main

import "fmt"

type BrainFuck struct {
	cells   []int
	crtCell int

	pc      int
	pcStack []int
}

func NewBrainFuck() *BrainFuck {
	return &BrainFuck{}
}

func (b *BrainFuck) Run(code string) {
	b.pc = 0
	b.crtCell = 0
	b.cells = []int{0}
	b.pcStack = []int{}

	endLoop := false
	for !endLoop {
		c := b.getIns(code)
		//fmt.Println(code[:b.pc])
		switch c {
		case '>':
			b.crtCell++
			if b.crtCell == len(b.cells) {
				b.cells = append(b.cells, 0)
			}
		case '<':
			b.crtCell--
			if b.crtCell < 0 {
				endLoop = true
			}
		case '+':
			b.cells[b.crtCell]++
		case '-':
			b.cells[b.crtCell]--
		case '[':
			b.pcStack = append(b.pcStack, b.pc)
		case ']':
			if len(b.pcStack) == 0 {
				fmt.Printf("mismatched brace at idx %d\n", b.pc)
				endLoop = true
			}
			if b.cells[b.crtCell] != 0 {
				b.pc = b.pcStack[len(b.pcStack)-1]
			} else {
				b.pcStack = b.pcStack[:len(b.pcStack)-1]
			}
		case '.':
			fmt.Printf("%c", b.cells[b.crtCell])
		default:
			endLoop = true
		}
		//fmt.Println(b.cells)
	}
	fmt.Println()
}

func (b *BrainFuck) getIns(code string) byte {
	for b.pc < len(code) && (code[b.pc] == ' ' || code[b.pc] == '\n' || code[b.pc] == '\r') {
		b.pc++
	}
	if b.pc == len(code) {
		return 255
	}
	ret := code[b.pc]
	b.pc++
	return ret
}

func main() {
	fmt.Println("vim-go")
	b := NewBrainFuck()
	//b.Run("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
	//b.Run("++++++++[>++++++++<-]>++++++++.>++++++++[>++++++++++++<-]>+++++.+++++++..+++.>++++++++[>+++++<-]>++++.------------.<<<<+++++++++++++++.>>.+++.------.--------.>>+.")
	//b.Run(">++++++++[-<+++++++++>]<.>>+>-[+]++>++>+++[>[->+++<<+++>]<<]>-----.>->+++..+++.>-.<<+[>[+>+]>>]<--------------.>>.+++.------.--------.>+.>+.")
	b.Run("--<-<<+[+[<+>--->->->-<<<]>]<<--.<++++++.<<-..<<.<+.>>.>>.<<<.+++.>>.>>-.<<<+.")
	b.Run("+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+.")
}

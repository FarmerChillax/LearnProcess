package main

// mac电脑对象, 实现了computer对象
type mac struct {
	printer printer
}

func (m *mac) setPrinter(printer printer) {
	m.printer = printer
}

func (m *mac) print() {
	m.printer.printFile()
}

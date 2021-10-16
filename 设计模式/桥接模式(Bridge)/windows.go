package main

type windows struct {
	printer printer
}

func (w *windows) print() {
	w.printer.printFile()
}

func (w *windows) setPrinter(printer printer) {
	w.printer = printer
}

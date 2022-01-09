package demo

import (
	"flag"
	"log"
)

func easyDemo() {
	var name string
	flag.StringVar(&name, "name", "Go工具集", "帮助信息")
	flag.StringVar(&name, "n", "Go工具集", "帮助信息")
	flag.Parse()

	log.Printf("name: %s\n", name)
}

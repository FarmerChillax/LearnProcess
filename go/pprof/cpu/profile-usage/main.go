package main

import (
	"pprof-learn/functions"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	functions.Concat(100)
}

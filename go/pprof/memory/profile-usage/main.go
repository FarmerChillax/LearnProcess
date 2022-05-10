package main

import (
	"pprof-learn/functions"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	functions.Concat(100)
}

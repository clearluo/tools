package main

import (
	"tool/common/basic/runBefore"
	"tool/dbd"
)

func init() {
	runBefore.InitRun()
}

func main() {
	dbd.Dbd()
	//finance.Finance()
}

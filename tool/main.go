package main

import (
	"tool/common/basic/runBefore"
	"tool/finance"
)

func init() {
	runBefore.InitRun()
}

func main() {
	//dbd.Dbd()
	finance.Finance()
}

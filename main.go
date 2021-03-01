package main

import (
	"Builder/compile"
	"Builder/sub"
	"Builder/workspace"
	"os"

)

func main() {
	workspace.MakeMainDir()
	arg := os.Args[2]
	compile.Go(arg)
	sub.MakeSubDir()
}


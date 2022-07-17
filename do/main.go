package main

import (
	"flag"
	"fmt"
)

func main() {
	cdToTopDir()
	fmt.Printf("topDir: '%s'\n", topDir())

	var (
		flgRun    bool
		flgDeploy bool
	)
	{
		flag.BoolVar(&flgRun, "run", false, "runs the program")
		flag.BoolVar(&flgDeploy, "deploy", false, "deploys to production")
		flag.Parse()
	}

	if flgRun {
		doRun()
		return
	}

	if flgDeploy {
		doDeploy()
		return
	}

	flag.Usage()
}

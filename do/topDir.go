package main

import (
	"fmt"
	"os"

	"github.com/keyluck/go_scripts/u"
)

func cdToTopDir() {
	err := os.Chdir("..")
	u.Must(err)
}

func topDir() string {
	dir, err := os.Getwd()
	u.Must(err)
	return dir
}

func doRun() {
	fmt.Println("run triggered")
	return
}

func doDeploy() {
	fmt.Println("production deployment triggered")
	return
}

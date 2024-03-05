package main

import "os"

func doPanic(msg string) {
	panic(msg)
}

func bis_bckp_main() {
	doPanic(os.Args[0])
}

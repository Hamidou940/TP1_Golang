package main

import (
	"flag"
	"fmt"
)

const VERSION = "1.0"

func main() {
	flagPresence:= flag.Bool("version", false, "affiche la version du programme")

	flag.Parse()

	if *flagPresence {
		fmt.Println(VERSION)
	}else{
		fmt.Println("Pas de flag version")
	}
}


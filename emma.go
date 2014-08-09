package main

import (
	"flag"
	"fmt"
	"time"
)

var nbemma int
var name string

func init() {
	flag.IntVar(&nbemma, "i", 0, "nb nom")
	flag.StringVar(&name, "n", "", "mon nom")
}

func main() {
	flag.Parse()

	for i := 0; i < nbemma; i++ {
		fmt.Printf("%d %s\n", i+1, name)
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Bisous les Loulous !\n")

}

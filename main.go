package main

import (
	"fmt"

	"github.com/MrAlias/erase/mushrooms"
)

func main() {
	fmt.Println("When mushroom hunting, look out for these poisonous mushrooms:")
	for _, m := range mushrooms.Poisonous {
		if m.CommonName != "" {
			fmt.Printf(" - %s (%s)\n", m.BinomialName, m.CommonName)
		} else {
			fmt.Printf(" - %s\n", m.BinomialName)
		}
	}
}

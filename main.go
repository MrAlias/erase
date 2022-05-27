package main

import (
	"fmt"

	"github.com/MrAlias/erase/mushroom"
)

func main() {
	fmt.Println("When mushroom hunting, look out for these poisonous mushrooms:")
	for _, m := range mushroom.Poisonous {
		fmt.Printf(" - %s (%s)\n", m.CommonName, m.BinomialName)
	}
}

package main

import (
	"fmt"

	"github.com/MrAlias/erase/mushrooms"
)

func printMushrooms(shrooms []mushrooms.Mushroom) {
	for _, m := range shrooms {
		if m.CommonName != "" {
			fmt.Printf(" - %s (%s)\n", m.BinomialName, m.CommonName)
		} else {
			fmt.Printf(" - %s\n", m.BinomialName)
		}
	}
}

func main() {
	fmt.Println("When mushroom hunting, be on the look out for these tasty mushrooms 🍄:")
	printMushrooms(mushrooms.Tasty)

	fmt.Println("\n⚠️ But be weary of these poisonous mushrooms ☠️!")
	printMushrooms(mushrooms.Poisonous)
}

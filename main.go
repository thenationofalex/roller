package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/TwiN/go-color"
)

func generateNumber(topRange int, modifier int) [2]int {
	num := (rand.Intn(topRange) + 1)
	total := num + modifier

	return [2]int{total, num}
}

func main() {
	var choice string
	var modifier string

	fmt.Println(color.With(color.Red, "🔥 Dice Roller 🔥"))
	flag.StringVar(&choice, "d", "20", "Dice, default is 20")
	flag.StringVar(&modifier, "m", "0", "Modifier, default is 0")
	flag.Parse()

	choiceInt, err := strconv.Atoi(choice)

	if err != nil {
		fmt.Println("invalid die")
		panic(err)
	}

	modifierInt, err := strconv.Atoi(modifier)

	if err != nil {
		fmt.Println("invalid modifier")
		panic(err)
	}

	found := false
	validChoices := [6]int{4, 6, 8, 10, 12, 20}

	for _, v := range validChoices {
		if v == choiceInt {
			found = true
			break
		}
	}

	if found {
		num := generateNumber(choiceInt, modifierInt)
		output := fmt.Sprintf("🎲 rolled %d (%d + %d modifier)", num[0], num[1], modifierInt)
		fmt.Println(color.With(color.Yellow, output))
	} else {
		fmt.Println(color.With(color.Red, "🔥 Invalid choice, please pick d4, d6, d8, d10, d12 or d20 🔥"))
	}

}

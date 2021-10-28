package randomroll

import "math/rand"

func randomDiceRoll(numberOfFaces int) int {
	return rand.Int()%numberOfFaces + 1
}

// D6 rolls a random six sided die.
func D6() int {
	return randomDiceRoll(6)
}

// D4 rolls a random 4 sided die.
func D4() int {
	return randomDiceRoll(4)
}

// D12 rolls a random 12 sided die.
func D12() int {
	return randomDiceRoll(12)
}

// D8 rolls a random 8 sided die.
func D8() int {
	return randomDiceRoll(8)
}

// D10 rolls a random 10 sided die.
func D10() int {
	return randomDiceRoll(10)
}

// D20 rolls a random 20 sided die.
func D20() int {
	return randomDiceRoll(20)
}

// RollAStat rolls a statistic for a player character.
func RollAStat(numberOfDice int) int {
	roll := make([]int, numberOfDice)

	for i := 0; i < numberOfDice; i++ {
		roll[i] = D6()
	}

	sortRoll(roll)
	return sumRoll(roll)
}

func sumRoll(roll []int) int {
	return roll[0] + roll[1] + roll[2]
}

func sortRoll(roll []int) {
	didSwap := true
	for didSwap {
		didSwap = false
		for i := 0; i < len(roll)-1; i++ {
			if roll[i+1] > roll[i] {
				roll[i], roll[i+1] = roll[i+1], roll[i]
				didSwap = true
			}
		}
	}
}

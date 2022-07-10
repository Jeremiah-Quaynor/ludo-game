// Will need to define the rules of the game for a player.
// If player rolls a 6 then count should start and then a player can move.
// If a player rools the dice and gets another 6 then player should be able
// to decide whether to move or bring in a new chip.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var steps int  = 0
var isHome bool = false

func dice() (rolledDice int) {
	rolledDice = rand.Intn(6)
	return rolledDice
}
func roll() {
	var value int; 
	value = dice()
	steps += value;
	fmt.Println("Your rolled a ", value)
	fmt.Println("Move ", value, "steps")
	home()
}
func home() {
	var celebrate rune = '🎉'
	if (steps > 57) {
		isHome = true
		fmt.Println("YOU WON ", string(celebrate))
	}else {
		fmt.Println("Roll again")
	}
}

func main() {
	for {
		fmt.Println("Hit 'r' to roll the dice")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if (isHome == true) {
			break
		}
		if (userInput == "r") {
			roll()
		}else {
			fmt.Println("End Game!!!!!")
			fmt.Println("Your total number of steps moved is ", steps)
			break
		}
	}
}


package main

import (
	"fmt"

	"bitbucket.org/VitJRBOG/hangman/judgeman"
	"bitbucket.org/VitJRBOG/hangman/opponent"
	"bitbucket.org/VitJRBOG/hangman/painter"
	"bitbucket.org/VitJRBOG/hangman/tools"
)

func main() {
	mistakesNumber, word, wordField, hangmanImage := preparingToGame()
	for {
		drawPlayground(hangmanImage, wordField)
		letter := usersStep()
		computersStep(hangmanImage, wordField, word, letter, mistakesNumber)
		continueGame := summarizing(hangmanImage, wordField)
		if !(continueGame) {
			break
		}
	}
}

func preparingToGame() (int, string, string, string) {
	word := opponent.GetWord()
	wordField := painter.DrawWordField(len(word))
	gallowsImage := painter.DrawGallows()

	return 0, word, wordField, gallowsImage
}

func usersStep() string {
	for {
		fmt.Println("-- Enter a one letter and press \"Enter\" --")
		userInput := tools.ReceiveUserInput()
		if checkUserAnswer(userInput) {
			return userInput
		}
	}
}

func checkUserAnswer(userInput string) bool {
	switch true {
	case len(userInput) == 0:
		fmt.Println("You must enter one letter. Please try again...")
		return false
	case len(userInput) > 1:
		fmt.Println("You must enter only one letter. Please try again...")
		return false
	default:
		return true
	}
}

func computersStep(hangmanImage, wordField, word, letter string, mistakesNumber int) (string, string) {
	indexes, ok := checkLetter(word, letter)
	if ok {
		for _, index := range indexes {
			wordField = painter.AddLetter(index, letter, wordField)
		}
	} else {
		mistakesNumber++
		hangmanImage = painter.AddHangmanElement(mistakesNumber)
	}
	return hangmanImage, wordField
}

func checkLetter(word, letter string) ([]int, bool) {
	indexes := judgeman.CheckLetter(letter, word)
	if len(indexes) > 0 {
		return indexes, true
	}
	return indexes, false
}

func drawPlayground(hangmanImage, wordField string) {
	fmt.Println(hangmanImage)
	fmt.Println(wordField)
}

func summarizing(hangmanImage, wordField string) bool {
	switch true {
	case judgeman.CheckWord(wordField):
		fmt.Println("Congratulations! You are winner!")
		return false
	case judgeman.CheckHangmanImage(hangmanImage):
		fmt.Println("Game over! Lucky next time...")
		return false
	default:
		return true
	}
}

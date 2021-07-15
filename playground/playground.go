package playground

import (
	"fmt"

	"github.com/VitJRBOG/hangman/judgeman"
	"github.com/VitJRBOG/hangman/painter"
	"github.com/VitJRBOG/hangman/tools"
)

func PreparingPlayground(word string) (string, string) {
	wordField := painter.DrawWordField(len([]rune(word)))
	gallowsImage := painter.DrawGallows()
	fmt.Printf("%s\n%s\n\n", gallowsImage, wordField)

	return wordField, gallowsImage
}

func GetLetterFromUser() string {
	for {
		fmt.Println("-- Enter a one letter and press \"Enter\" --")
		userInput := tools.ReceiveUserInput()
		if checkUserInput(userInput) {
			return userInput
		}
	}
}

func checkUserInput(userInput string) bool {
	switch true {
	case len([]rune(userInput)) == 0:
		fmt.Println("You must enter one letter. Please try again...")
		return false
	case len([]rune(userInput)) > 1:
		fmt.Println("You must enter only one letter. Please try again...")
		return false
	default:
		return true
	}
}

func LookingLetterInWord(letter, word string) []int {
	indexes := judgeman.CheckLetter(letter, word)
	return indexes
}

func Painting(mistakesNumber int, indexes []int, letter, wordField, hangmanImage string) (int, string, string) {
	if len(indexes) > 0 {
		for _, index := range indexes {
			wordField = painter.AddLetter(index, letter, wordField)
		}
	} else {
		mistakesNumber++
		hangmanImage = painter.AddHangmanElement(mistakesNumber)
	}
	return mistakesNumber, wordField, hangmanImage
}

func Summarizing(hangmanImage, wordField string) bool {
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
